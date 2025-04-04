package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"pajarit-timeline-worker/domain"
	"time"

	"github.com/google/uuid"
)

type SqlitePostRepository struct {
	dbClient *sql.DB
}

func NewSqlitePostRepository(dbClient *sql.DB) *SqlitePostRepository {
	return &SqlitePostRepository{dbClient: dbClient}
}

func (r *SqlitePostRepository) Save(ctx context.Context, post *domain.Post) (*domain.Post, error) {

	toInsert := post
	toInsert.Id = uuid.New().String()
	toInsert.CreatedAt = time.Now().UTC()

	_, err := r.dbClient.Exec(
		"INSERT INTO posts (id, author_id, content, created_at) VALUES (?, ?, ?, ?)",
		toInsert.Id, post.AuthorId, post.Content, post.CreatedAt,
	)

	// TODO - tipar el error

	if err != nil {
		return nil, fmt.Errorf("can't insert post %v", err)
	}

	return toInsert, nil
}
