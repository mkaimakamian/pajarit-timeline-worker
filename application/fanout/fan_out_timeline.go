package fanout

import (
	"context"
	"fmt"
	"log"
	"pajarit-timeline-worker/domain"
)

type FanOutTimeline struct {
	timelineRepository domain.TimelineRepository
	followUpRepository domain.FollowUpRepository
}

func NewFanOutTimeline(timelineRepository domain.TimelineRepository, followUpRepository domain.FollowUpRepository) FanOutTimeline {
	return FanOutTimeline{
		timelineRepository: timelineRepository,
		followUpRepository: followUpRepository,
	}
}

func (e *FanOutTimeline) Exec(ctx context.Context, event PostCreatedEvent) error {

	followUp, err := e.followUpRepository.Get(ctx, event.AuthorId)
	if err != nil {
		log.Println(err)
	}

	// 2. insertar como primer registro el evento
	for _, follow := range followUp {
		// Go routine para ejecutar en paralelo todas las actualizaciones
		// Guarda con el fallo de una
		fmt.Println(follow)
	}

	// 2.1 garantizar el orden

	return nil
}
