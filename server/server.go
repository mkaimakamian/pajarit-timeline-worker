package server

import (
	"fmt"
	"pajarit-timeline-worker/config"

	"github.com/nats-io/nats.go"
)

func StartServer(cfg *config.Configuration, deps *config.Dependencies) error {

	natsUrl := fmt.Sprintf("%s:%d", cfg.Event.Server, cfg.Event.Port)
	natsConnection, err := nats.Connect(natsUrl)
	if err != nil {
		return err
	}
	defer natsConnection.Close()

	natsConnection.Subscribe("post.created", PostCreatedEventHandler(deps))

	select {}
}
