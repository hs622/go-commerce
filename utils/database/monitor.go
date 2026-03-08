package database

import (
	"context"
	"fmt"

	"github.com/hs622/ecommerce-cart/utils"
	"go.mongodb.org/mongo-driver/v2/event"
)

type CommandMonitor struct{}

func (c *CommandMonitor) Started(ctx context.Context, event *event.CommandStartedEvent) {
	utils.Query(fmt.Sprintf("[%s] %v", event.CommandName, event.Command))
}

func (c *CommandMonitor) Failed(ctx context.Context, evt *event.CommandFailedEvent) {}

func (c *CommandMonitor) Finished(ctx context.Context, evt *event.CommandFinishedEvent) {}

func (c *CommandMonitor) Succeeded(ctx context.Context, evt *event.CommandSucceededEvent) {}
