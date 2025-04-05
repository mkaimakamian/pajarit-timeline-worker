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

	var mu sync.Mutex

	for _, follow := range followUp {
		followerId := follow.FollowerId

		go func() {

			// Es un poco contradictorio el uso de go routines con la única función
			// lockeada, pero es sencillamente porque SQLite no soporta concurrencia
			// y la idea era mostrar que el fanout se puede paralelizar (con una base adecuada)
			mu.Lock()
			err := e.timelineRepository.Save(ctx, post, followerId)
			mu.Unlock()
			log.Println(err)

			// TODO - falta implementar una política de reintento ante errores.
			// Se podría generar un evento del tipo timetable.update.failed y
			// enviar cada update fallido al broker de eventos para un futuro reintento.
		}()
	}

}
