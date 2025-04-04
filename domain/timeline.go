package domain

import "context"

type TimelineRepository interface {
	Get(ctx context.Context, userId string) (*Timeline, error)
}

type Timeline struct {
	Posts []Post
}
