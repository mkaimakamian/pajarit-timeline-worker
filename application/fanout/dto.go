package fanout

import "time"

type PostCreatedEvent struct {
	Id        string    `json:"id"`
	AuthorId  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
