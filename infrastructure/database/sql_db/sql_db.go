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

func (sd *SqlDatabase) Close() {
	err := sd.Client.Close()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to close the database connection")
	} else {
		log.Info().Msg("Database connection closed successfully")
	}
}

func (sd *SqlDatabase) CallProcedure(name string) error {
	query := fmt.Sprintf("CALL escalateservice.%s()", name)
	_, err := sd.Client.Exec(query)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to call procedure %s", name)
	}
	return err
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

func (sd *SqlDatabase) BatchAddUsers(users []*model.User) error {
	if len(users) == 0 {
		log.Warn().Msg("No users to batch add")
		return nil
	}
	query := `
		INSERT INTO escalateservice.users (
			username
		) VALUES %s
	`
	values := make([]string, len(users))
	for i, user := range users {
		values[i] = fmt.Sprintf("('%s')", user.Username)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))
	_, err := sd.Client.Exec(query)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to batch add users")
		return err
	}
	log.Info().Msgf("Batch added %d users successfully", len(users))
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
			return nil, nil
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

func (sd *SqlDatabase) BatchAddPosts(posts []*model.Post) error {
	if len(posts) == 0 {
		log.Warn().Msg("No posts to batch add")
		return nil
	}
	query := `
		INSERT INTO escalateservice.posts (
			post_id,
			username
		) VALUES %s
	`
	values := make([]string, len(posts))
	for i, post := range posts {
		values[i] = fmt.Sprintf("('%s', '%s')", post.PostId, post.Username)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))
	_, err := sd.Client.Exec(query)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to batch add posts")
		return err
	}

	log.Info().Msgf("Batch added %d posts successfully", len(posts))
	return nil
}

func (sd *SqlDatabase) GetPost(postId string) (*model.Post, error) {
	query := `
		SELECT 
			post_id,
        	username,
			score
		FROM escalateservice.posts
		WHERE post_id = $1
	`

	var post model.Post
	err := sd.Client.QueryRow(query, postId).Scan(
		&post.PostId,
		&post.Username,
		&post.Score,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error().Stack().Err(err).Msgf("Get post by postId %s failed", postId)
		return nil, err
	}

	return &post, nil
}

func (sd *SqlDatabase) AddReview(review *model.Review) error {
	query := `
		INSERT INTO escalateservice.reviews (
			review_id,
			post_id,
        	reviewer,
			rating
    	) VALUES ($1, $2, $3, $4)
	`
	err := sd.insertData(query, review.ReviewId, review.PostId, review.Reviewer, review.Rating)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create review, reviewId: %d", review.ReviewId)
		return err
	}

	log.Info().Msgf("Review created successfully, reviewId: %d", review.ReviewId)
	return nil
}

func (sd *SqlDatabase) BatchAddReviews(reviews []*model.Review) error {
	if len(reviews) == 0 {
		log.Warn().Msg("No reviews to batch add")
		return nil
	}

	query := `
		INSERT INTO escalateservice.reviews (
			review_id,
			post_id,
			reviewer,
			rating
		) VALUES %s
	`
	values := make([]string, len(reviews))
	for i, review := range reviews {
		if review == nil {
			log.Error().Stack().Msg("Review is null")
			return fmt.Errorf("review is null")
		}
		values[i] = fmt.Sprintf("(%d, '%s', '%s', %d)", review.ReviewId, review.PostId, review.Reviewer, review.Rating)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))

	_, err := sd.Client.Exec(query)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to batch add reviews")
		return err
	}

	log.Info().Msgf("Batch added %d reviews successfully", len(reviews))
	return nil
}

func (sd *SqlDatabase) GetReview(reviewId uint64) (*model.Review, error) {
	query := `
		SELECT 
			review_id,
			post_id,
        	reviewer,
			rating
		FROM escalateservice.reviews
		WHERE review_id = $1
	`

	var review model.Review
	err := sd.Client.QueryRow(query, reviewId).Scan(
		&review.ReviewId,
		&review.PostId,
		&review.Reviewer,
		&review.Rating,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No review found with the given ID
		}
		log.Error().Stack().Err(err).Msgf("Get post by reviewId %d failed", reviewId)
		return nil, err
	}

	return &review, nil
}

func (sd *SqlDatabase) AddLikePost(likePost *model.LikePost) error {
	query := `
		INSERT INTO escalateservice.likePosts (
			username,
			post_id
    	) VALUES ($1, $2)
	`
	err := sd.insertData(query, likePost.Username, likePost.PostId)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create likePost, username: %s -> post %s", likePost.Username, likePost.PostId)
		return err
	}

	log.Info().Msgf("LikePost created successfully, username: %s -> post %s", likePost.Username, likePost.PostId)
	return nil
}

func (sd *SqlDatabase) BatchAddLikePosts(likePosts []*model.LikePost) error {
	if len(likePosts) == 0 {
		log.Warn().Msg("No likePosts to batch add")
		return nil
	}

	query := `
		INSERT INTO escalateservice.likePosts (
			username,
			post_id
		) VALUES %s
	`
	values := make([]string, len(likePosts))
	for i, likePost := range likePosts {
		if likePost == nil {
			log.Error().Stack().Msg("LikePost is null")
			return fmt.Errorf("likePost is null")
		}
		values[i] = fmt.Sprintf("('%s', '%s')", likePost.Username, likePost.PostId)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))

	_, err := sd.Client.Exec(query)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to batch add likePosts")
		return err
	}

	log.Info().Msgf("Batch added %d likePosts successfully", len(likePosts))
	return nil
}

func (sd *SqlDatabase) GetLikePost(username, postId string) (*model.LikePost, error) {
	query := `
		SELECT 
    		username,
    		post_id
		FROM escalateservice.likePosts
		WHERE username = $1 AND post_id = $2
	`

	var likePost model.LikePost
	err := sd.Client.QueryRow(query, username, postId).Scan(
		&likePost.Username,
		&likePost.PostId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No likePost found with the given username and postId
		}
		log.Error().Stack().Err(err).Msgf("Get likePost by username %s and postId %s failed", username, postId)
		return nil, err
	}

	return &likePost, nil
}

func (sd *SqlDatabase) RemoveLikePost(likePost *model.LikePost) error {
	query := `
		DELETE FROM escalateservice.likePosts
		WHERE username = $1 AND post_id = $2
	`

	result, err := sd.Client.Exec(query, likePost.Username, likePost.PostId)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to remove likePost, username: %s -> post %s", likePost.Username, likePost.PostId)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to get rows affected for likePost removal, username: %s -> post %s", likePost.Username, likePost.PostId)
		return err
	}

	if rowsAffected == 0 {
		log.Warn().Msgf("No likePost found to remove, username: %s -> post %s", likePost.Username, likePost.PostId)
		return nil // No error, but no rows were affected
	}

	log.Info().Msgf("LikePost removed successfully, username: %s -> post %s", likePost.Username, likePost.PostId)
	return nil
}

func (sd *SqlDatabase) AddSuperlikePost(superlikePost *model.SuperlikePost) error {
	query := `
		INSERT INTO escalateservice.superlikePosts (
			username,
			post_id
    	) VALUES ($1, $2)
	`
	err := sd.insertData(query, superlikePost.Username, superlikePost.PostId)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create superlikePost, username: %s -> post %s", superlikePost.Username, superlikePost.PostId)
		return err
	}

	log.Info().Msgf("SuperlikePost created successfully, username: %s -> post %s", superlikePost.Username, superlikePost.PostId)
	return nil
}

func (sd *SqlDatabase) BatchAddSuperlikePosts(superlikePosts []*model.SuperlikePost) error {
	if len(superlikePosts) == 0 {
		log.Warn().Msg("No superlikePosts to batch add")
		return nil
	}

	query := `
		INSERT INTO escalateservice.superlikePosts (
			username,
			post_id
		) VALUES %s
	`
	values := make([]string, len(superlikePosts))
	for i, superlikePost := range superlikePosts {
		if superlikePost == nil {
			log.Error().Stack().Msg("SuperlikePost is null")
			return fmt.Errorf("superlikePost is null")
		}
		values[i] = fmt.Sprintf("('%s', '%s')", superlikePost.Username, superlikePost.PostId)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))
	_, err := sd.Client.Exec(query)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to batch add superlikePosts")
		return err
	}

	log.Info().Msgf("Batch added %d superlikePosts successfully", len(superlikePosts))
	return nil
}

func (sd *SqlDatabase) GetSuperlikePost(username, postId string) (*model.SuperlikePost, error) {
	query := `
		SELECT 
    		username,
    		post_id
		FROM escalateservice.superlikePosts
		WHERE username = $1 AND post_id = $2
	`

	var superlikePost model.SuperlikePost
	err := sd.Client.QueryRow(query, username, postId).Scan(
		&superlikePost.Username,
		&superlikePost.PostId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No superlikePost found with the given username and postId
		}
		log.Error().Stack().Err(err).Msgf("Get superlikePost by username %s and postId %s failed", username, postId)
		return nil, err
	}

	return &superlikePost, nil
}

func (sd *SqlDatabase) RemoveSuperlikePost(superlikePost *model.SuperlikePost) error {
	query := `
		DELETE FROM escalateservice.superlikePosts
		WHERE username = $1 AND post_id = $2
	`

	result, err := sd.Client.Exec(query, superlikePost.Username, superlikePost.PostId)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to remove superlikePost, username: %s -> post %s", superlikePost.Username, superlikePost.PostId)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to get rows affected for superlikePost removal, username: %s -> post %s", superlikePost.Username, superlikePost.PostId)
		return err
	}

	if rowsAffected == 0 {
		log.Warn().Msgf("No superlikePost found to remove, username: %s -> post %s", superlikePost.Username, superlikePost.PostId)
		return nil // No error, but no rows were affected
	}

	log.Info().Msgf("SuperlikePost removed successfully, username: %s -> post %s", superlikePost.Username, superlikePost.PostId)
	return nil
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

	_, err = tx.Exec(query, args...)

	return err
}
