package fanout

import (
	"context"
	"log"
	"pajarit-timeline-worker/domain"
	"sync"
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

	e.updateFollowersTimeline(ctx, followUp, post)

	return nil
}

func (e *FanOutTimeline) updateFollowersTimeline(ctx context.Context, followUp []*domain.FollowUp, post *domain.Post) {
	// Author's post must be added to its own timeline!
	followUp = append(followUp, &domain.FollowUp{FollowerId: post.AuthorId})

	var wg sync.WaitGroup

	for _, follow := range followUp {
		followerId := follow.FollowerId

		wg.Add(1)
		go func() {
			defer wg.Done()
			err := e.timelineRepository.Save(ctx, post, followerId)
			log.Println(err)

			// TODO - falta implementar una política de reintento ante errores,
			// posiblemente basada en retornar la tarea a la cola de eventos
			// para su posterior tratamiento (según el tipo de error)
		}()
	}
}
