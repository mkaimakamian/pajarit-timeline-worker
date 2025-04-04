package fanout

type PostCreatedEvent struct {
	PostId    string `json:"post_id"`
	AuthorId  string `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
