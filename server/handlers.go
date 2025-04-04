package server

import (
	"context"
	"encoding/json"
	"log"

	"pajarit-timeline-worker/application/fanout"
	"pajarit-timeline-worker/config"

	"github.com/nats-io/nats.go"
)

func PostCreatedEventHandler(deps *config.Dependencies) nats.MsgHandler {
	usecase := fanout.NewFanOutTimeline(deps.TimelineRepository, deps.FollowUpRepository)

	return func(msg *nats.Msg) {
		var event fanout.PostCreatedEvent
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Printf("failed to unmarshal event: %v", err)
			return
		}

		err = usecase.Exec(context.Background(), event)

		// TODO - ver esto que quedó cualquier cosa
		if err != nil {
			log.Printf("error when executing fan out timeline: %v", err)
		}
	}
}
