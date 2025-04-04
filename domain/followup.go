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

// func NewFollowUp(followerId, followedId string) (*FollowUp, error) {

// 	// Podemos desacoplar la validación, pero me resultó más
// 	// práctico tratar la entidad como un value object

// 	if len(followerId) == ZERO_LENGTH {
// 		return nil, fmt.Errorf("follower id can't be %d length", ZERO_LENGTH)
// 	}

// 	if len(followedId) == ZERO_LENGTH {
// 		return nil, fmt.Errorf("followed id can't be %d length", ZERO_LENGTH)
// 	}

// 	return &FollowUp{FollowerId: followerId, FollowedId: followedId}, nil
// }
