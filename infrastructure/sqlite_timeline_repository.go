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
	// completa (LPUSH o JSON.ARRINSERT en Redis, list_append en Dynamo, etc) en la implementación actual
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

	index, exists := r.findInsertIndex(post, posts)
	if exists {
		return nil
	}

	timeline := append(posts, nil)
	copy(timeline[index+1:], timeline[index:])
	timeline[index] = post

	timelineToDB, err := json.Marshal(timeline)
	if err != nil {
		return err
	}

	_, err = r.dbClient.Exec("UPDATE timelines SET posts = ? WHERE user_id = ?", string(timelineToDB), followerId)
	return err
}

// Looks for the position that newPost should be inserted at, if it doesn't exist.
// If the element already exists, returns its position and true as it exists.
// If the element doesn't exist, returns the position to be inserted and false.
func (r *SqliteTimelineRepository) findInsertIndex(newPost *domain.Post, posts []*domain.Post) (int, bool) {

	// En un sistema real podrían utilizarse las estructuras de datos nativas de bases
	// key-value (ZSET en Redis, por ejemplo) que permiten insertar elementos
	// solo si no existen, asegurando unicidad de forma eficiente.
	// Al estar simulando con SQLite una base key-value, la idempotencia y el orden
	// requieren más esfuerzo.

	for i, post := range posts {
		if post.CreatedAt == newPost.CreatedAt && post.Id == newPost.Id {
			return i, true
		}

		if post.CreatedAt.Before(newPost.CreatedAt) {
			return i, false
		}
	}

	return len(posts), false
}
