package domain

import (
	"context"
	"fmt"
	"time"
)

const MAX_ALLOWED_LENGTH = 280
const ZERO_LENGTH = 0 // TODO - mover esta constante porque se usa en otras entidades

type PostRepository interface {
	Save(ctx context.Context, post *Post) (*Post, error)
}

type Post struct {
	Id        string
	AuthorId  string
	Content   string
	CreatedAt time.Time
}

func NewPost(authorId, content string) (*Post, error) {

	// Podemos desacoplar la validación, pero me resultó más
	// práctico tratar la entidad como un value object

	if len(authorId) == ZERO_LENGTH {
		return nil, fmt.Errorf("author id can't be %d length", ZERO_LENGTH)
	}

	if len(content) == ZERO_LENGTH {
		return nil, fmt.Errorf("post can't be %d length", ZERO_LENGTH)
	}

	if len(content) > MAX_ALLOWED_LENGTH {
		return nil, fmt.Errorf("post can't exceed %d characters", MAX_ALLOWED_LENGTH)
	}

	return &Post{AuthorId: authorId, Content: content}, nil
}
