package domain

import (
	"fmt"
	"time"
)

const MAX_ALLOWED_LENGTH = 280
const ZERO_LENGTH = 0 // TODO - mover esta constante porque se usa en otras entidades

type Post struct {
	Id        string
	AuthorId  string
	Content   string
	CreatedAt time.Time
}

func NewPost(id, authorId, content string, createdAt time.Time) (*Post, error) {

	// En principio, la validación en esta instancia del flujo es opcional ya que
	// no deberían llegar mensajes incompletos; sin embargo, no deja
	// de ser una buena práctica.

	if len(authorId) == ZERO_LENGTH {
		return nil, fmt.Errorf("author id can't be %d length", ZERO_LENGTH)
	}

	if len(content) == ZERO_LENGTH {
		return nil, fmt.Errorf("post can't be %d length", ZERO_LENGTH)
	}

	if len(content) > MAX_ALLOWED_LENGTH {
		return nil, fmt.Errorf("post can't exceed %d characters", MAX_ALLOWED_LENGTH)
	}

	return &Post{Id: id, AuthorId: authorId, Content: content, CreatedAt: createdAt}, nil
}
