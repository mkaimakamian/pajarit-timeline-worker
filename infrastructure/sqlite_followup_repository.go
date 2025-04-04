package infrastructure

import (
	"context"
	"database/sql"
	"pajarit-timeline-worker/domain"
)

type SqliteFollowUpRepository struct {
	dbClient *sql.DB
}

func NewSqliteFollowUpRepository(dbClient *sql.DB) *SqliteFollowUpRepository {
	return &SqliteFollowUpRepository{dbClient: dbClient}
}

func (r *SqliteFollowUpRepository) Get(ctx context.Context, userId string) ([]*domain.FollowUp, error) {

	rows, err := r.dbClient.Query("SELECT follower_id, followed_id FROM followup WHERE followed_id = ?", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var followUps []*domain.FollowUp
	for rows.Next() {
		var follow domain.FollowUp
		err := rows.Scan(&follow.FollowerId, &follow.FollowedId)
		if err != nil {
			return nil, err
		}

		followUps = append(followUps, &follow)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followUps, nil
}
