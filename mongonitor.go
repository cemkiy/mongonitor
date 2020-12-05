package mongonitor

import (
	"context"
	"fmt"

	"github.com/biges/mgo/bson"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/event"
)

// NewMongonitor returns new mongonitor.
func NewMongonitor() *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(c context.Context, e *event.CommandStartedEvent) {
			command := make(map[string]interface{})
			_ = bson.Unmarshal(e.Command, command)
			output := ""
			color.Green("%d. request details-> ", e.RequestID)
			for k, v := range command {
				switch k {
				case "lsid", "$clusterTime", "singleBatch", "txnNumber":
					continue
				default:
					output += fmt.Sprintf("%s: %v ", k, v)
				}
			}
			color.Blue(output)
		},
		Failed: func(c context.Context, e *event.CommandFailedEvent) {
			color.Red("%s", e.Failure)
		},
		Succeeded: func(c context.Context, e *event.CommandSucceededEvent) {
			color.Yellow("%d. request completed in %d milliseconds", e.RequestID, e.DurationNanos/1000000)
		},
	}
}
