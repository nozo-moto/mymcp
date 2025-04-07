package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	version = "v0.0.1"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func run() error {
	mcpServer := server.NewMCPServer("nozomoto mcps", version)
	mcpServer.AddTools([]server.ServerTool{
		{
			Tool: mcp.NewTool(
				"uuid",
				mcp.WithDescription("return generated uuid"),
			),
			Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
				uuidStr := uuid.New().String()
				return mcp.NewToolResultText(uuidStr), nil
			},
		},
		{
			Tool: mcp.NewTool(
				"curren_time",
				mcp.WithDescription("return current time"),
			),
			Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
				currentTime := time.Now().Format(time.RFC3339)
				return mcp.NewToolResultText(currentTime), nil
			},
		},
	}...)

	if err := server.ServeStdio(mcpServer); err != nil {
		return err
	}

	return nil
}
