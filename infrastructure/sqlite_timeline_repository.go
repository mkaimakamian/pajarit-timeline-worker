package infrastructure

import (
	"context"
	"database/sql"
	"encoding/json"
	"pajarit-timeline-worker/domain"
)

type SqliteTimelineRepository struct {
	dbClient *sql.DB
}

func NewSqliteTimelineRepository(dbClient *sql.DB) *SqliteTimelineRepository {
	return &SqliteTimelineRepository{dbClient: dbClient}
}

func (r *SqliteTimelineRepository) Save(ctx context.Context, post *domain.Post, followerId string) error {

	// Para el challenge se está usando una base SQLite, tratando de simular el comportamiento
	// de una base key-value, aunque con limitaciones: se guarda un JSON y no una colección.
	// Mientras que en una base key-value es posible agregar un elemento sin recuperar previamente la colección
	// completa (LPUSH o JSON.ARRINSERT en Redis, list_append con SET en Dynamo, etc) en la implementación actual
	// si o si hay que recuperar todo el JSON para luego actualizarlo.
	var postsFromDB string
	err := r.dbClient.QueryRow("SELECT posts FROM timelines WHERE user_id = ?", followerId).Scan(&postsFromDB)

	if err == sql.ErrNoRows {
		err := r.createTimeline(post, followerId)
		return err
	}

	if err != nil {
		return err
	}

	err = r.updateTimeline(postsFromDB, post, followerId)
	return err
}

func (r *SqliteTimelineRepository) createTimeline(post *domain.Post, followerId string) error {
	posts := []*domain.Post{post}
	postToDB, _ := json.Marshal(posts)
	_, err := r.dbClient.Exec("INSERT INTO timelines (user_id, posts) VALUES (?, ?)", followerId, string(postToDB))
	return err
}

func (r *SqliteTimelineRepository) updateTimeline(postsFromDB string, post *domain.Post, followerId string) error {

	var posts []*domain.Post
	err := json.Unmarshal([]byte(postsFromDB), &posts)
	if err != nil {
		return err
	}

	// Se agrega el elemento como primero en el array para poder garantizar el ordenamiento
	// descendiente, facilitando la paginación al momento de solicitar el timeline
	posts = append([]*domain.Post{post}, posts...)

	updatedPostsToDB, err := json.Marshal(posts)
	if err != nil {
		return err
	}

	_, err = r.dbClient.Exec("UPDATE timelines SET posts = ? WHERE user_id = ?", string(updatedPostsToDB), followerId)
	return err
}
