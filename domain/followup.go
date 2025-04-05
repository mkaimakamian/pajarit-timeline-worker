package domain

import (
	"context"
)

type FollowUpRepository interface {
	Get(ctx context.Context, userId string) ([]*FollowUp, error)
}

type FollowUp struct {
	FollowerId string
	FollowedId string
}
