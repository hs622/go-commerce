package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/event"
)

type CommandMonitor struct{}

func (c *CommandMonitor) Started(ctx context.Context, event *event.CommandStartedEvent) {
	fmt.Printf(
		"Command: %s\nDB: %s\nQuery: %v\n\n",
		event.CommandName,
		event.DatabaseName,
		event.Command,
	)
}

func (c *CommandMonitor) Failed(ctx context.Context, evt *event.CommandFailedEvent) {}

func (c *CommandMonitor) Finished(ctx context.Context, evt *event.CommandFinishedEvent) {}

func (c *CommandMonitor) Succeeded(ctx context.Context, evt *event.CommandSucceededEvent) {}
