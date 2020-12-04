package mongonitor

import (
	"context"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/event"
)

// NewMongonitor returns new mongonitor.
func NewMongonitor() *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(c context.Context, e *event.CommandStartedEvent) {
			color.Blue("%s", e.Command)
		},
		Failed: func(c context.Context, e *event.CommandFailedEvent) {
			color.Red("%s", e.Failure)
		},
		Succeeded: func(c context.Context, e *event.CommandSucceededEvent) {
			color.Yellow("%d. request completed in %d milliseconds", e.RequestID, e.DurationNanos/1000000)
		},
	}
}
