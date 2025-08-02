package sql_db

import (
	"database/sql"
	model "escalateservice/internal/model/domain"
	"fmt"
	"strings"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/zerolog/log"
)

type SqlDatabase struct {
	Client *sql.DB
}

func NewDatabase(connStr string) (*SqlDatabase, error) {
	db, err := sql.Open("postgres", strings.Trim(connStr, "\""))
	if err != nil {
		log.Error().Stack().Err(err).Msg("Couldn't open a connection with the database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Database is not reachable")
		return nil, err
	}

	return &SqlDatabase{
		Client: db,
	}, nil
}

func (sd *SqlDatabase) Clean() {
	tx, err := sd.Client.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Clean each table
	for _, table := range tables {
		query := fmt.Sprintf("DELETE FROM escalateservice.%s", table)
		_, err = tx.Exec(query)
		if err != nil {
			log.Error().Stack().Err(err).Msgf("Failed to clean table %s", table)
		}
	}

	log.Info().Msg("Database cleaned successfully")
}

func (sd *SqlDatabase) AddUser(user *model.User) error {
	query := `
		INSERT INTO escalateservice.users (
        	username
    	) VALUES ($1)
	`
	err := sd.insertData(query, user.Username)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create user, username: %s", user.Username)
		return err
	}

	log.Info().Msgf("User created successfully, username: %s", user.Username)
	return nil
}

func (sd *SqlDatabase) GetUser(username string) (*model.User, error) {
	query := `
		SELECT 
			username
		FROM escalateservice.users
		WHERE username = $1
	`

	var user model.User
	err := sd.Client.QueryRow(query, username).Scan(&user.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No review found with the given ID
		}
		log.Error().Stack().Err(err).Msgf("Get user by username %s failed", username)
		return nil, err
	}

	return &user, nil
}

func (sd *SqlDatabase) AddPost(post *model.Post) error {
	query := `
		INSERT INTO escalateservice.posts (
			post_id,
        	username
    	) VALUES ($1, $2)
	`
	err := sd.insertData(query, post.PostId, post.Username)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create post, postId: %s", post.PostId)
		return err
	}

	log.Info().Msgf("Post created successfully, postId: %s", post.PostId)
	return nil
}

func (sd *SqlDatabase) GetPost(postId string) (*model.Post, error) {
	query := `
		SELECT 
			post_id,
        	username
		FROM escalateservice.posts
		WHERE post_id = $1
	`

	var post model.Post
	err := sd.Client.QueryRow(query, postId).Scan(
		&post.PostId,
		&post.Username,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No review found with the given ID
		}
		log.Error().Stack().Err(err).Msgf("Get post by postId %s failed", postId)
		return nil, err
	}

	return &post, nil
}

func (sd *SqlDatabase) insertData(query string, args ...any) error {
	tx, err := sd.Client.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec(
		query,
		args...)

	if err != nil {
		return err
	}

	return nil
}
