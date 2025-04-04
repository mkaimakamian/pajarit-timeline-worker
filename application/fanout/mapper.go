package fanout

import "pajarit-timeline-worker/domain"

func mapEventToPost(event PostCreatedEvent) (*domain.Post, error) {
	return domain.NewPost(event.Id, event.AuthorId, event.Content, event.CreatedAt)
}
