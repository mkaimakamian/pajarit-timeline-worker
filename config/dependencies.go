package config

import (
	"database/sql"
	"fmt"
	"pajarit-timeline-worker/domain"
	"pajarit-timeline-worker/infrastructure"

	_ "modernc.org/sqlite"
)

type Dependencies struct {
	TimelineRepository domain.TimelineRepository
	FollowUpRepository domain.FollowUpRepository
}

func BuildDependencies(cfg *Configuration) (*Dependencies, error) {

	dbClient, err := dbClient(cfg)
	if err != nil {
		return nil, err
	}

	// postRepository := infrastructure.NewSqlitePostRepository(dbClient)
	folloUpRepository := infrastructure.NewSqliteFollowUpRepository(dbClient)

	deps := &Dependencies{
		TimelineRepository: nil,
		FollowUpRepository: folloUpRepository,
	}

	return deps, nil
}

func dbClient(cfg *Configuration) (*sql.DB, error) {
	client, err := sql.Open("sqlite", cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("db can't be opened: %v", err)
	}

	// Valores arbitrarios para el challenge
	client.SetMaxOpenConns(cfg.DBMaxConnection)
	client.SetMaxIdleConns(cfg.DBMaxIdleConnection)

	if err = client.Ping(); err != nil {
		return nil, fmt.Errorf("db is not responding: %v", err)
	}

	return client, nil
}
