package domain

import "context"

type TimelineRepository interface {
	Save(ctx context.Context, post *Post, followerId string) error
}

type Timeline struct {
	Posts []Post
}
