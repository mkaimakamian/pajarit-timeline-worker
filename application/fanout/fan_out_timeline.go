package fanout

import (
	"context"
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
	post, err := mapEventToPost(event)
	if err != nil {
		log.Println(err)
		return err
	}

	followUp, err := e.followUpRepository.Get(ctx, post.AuthorId)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, follow := range followUp {
		err := e.timelineRepository.Save(ctx, post, follow.FollowerId)
		log.Println(err)

		// TODO - guarda que puede fallar y hay que ver cómo hacemos con eso.
	}

	// 2.1 garantizar el orden

	return nil
}
