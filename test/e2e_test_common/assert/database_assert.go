package e2e_test_assert

import (
	"escalateservice/cmd/startup"
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
	"escalateservice/test/e2e_test_common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ProvideDatabase(t *testing.T) *database.Database {
	provider := startup.NewProvider(e2e_test_common.Env, e2e_test_common.ConnStr)
	sqlDb, err := provider.ProvideDb()
	if err != nil {
		t.Fatalf("Failed to provide database: %v", err)
	}
	return database.NewDatabase(sqlDb)
}

func AssertUserExists(t *testing.T, username string, expectedUser *model.User) {
	db := ProvideDatabase(t)
	const maxRetries = 20
	const delay = 1000 * time.Millisecond
	for i := 0; i < maxRetries; i++ {
		user, err := db.Client.GetUser(username)
		if err == nil && user == nil {
			time.Sleep(delay)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, expectedUser.Username, user.Username)
			return
		}
	}
	user, err := db.Client.GetUser(username)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Username, user.Username)
}

func AssertPostExists(t *testing.T, postId string, expectedPost *model.Post) {
	db := ProvideDatabase(t)
	const maxRetries = 20
	const delay = 1000 * time.Millisecond
	for i := 0; i < maxRetries; i++ {
		post, err := db.Client.GetPost(postId)
		if err == nil && post == nil {
			time.Sleep(delay)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, expectedPost.PostId, post.PostId)
			assert.Equal(t, expectedPost.Username, post.Username)
			return
		}
	}
	post, err := db.Client.GetPost(postId)
	assert.Nil(t, err)
	assert.Equal(t, expectedPost.PostId, post.PostId)
	assert.Equal(t, expectedPost.Username, post.Username)
}

func AssertReviewExists(t *testing.T, reviewId uint64, expectedReview *model.Review) {
	db := ProvideDatabase(t)
	const maxRetries = 20
	const delay = 1000 * time.Millisecond
	for i := 0; i < maxRetries; i++ {
		review, err := db.Client.GetReview(reviewId)
		if err == nil && review == nil {
			time.Sleep(delay)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, expectedReview.ReviewId, review.ReviewId)
			assert.Equal(t, expectedReview.PostId, review.PostId)
			assert.Equal(t, expectedReview.Reviewer, review.Reviewer)
			assert.Equal(t, expectedReview.Rating, review.Rating)
			return
		}
	}
	review, err := db.Client.GetReview(reviewId)
	assert.Nil(t, err)
	assert.Equal(t, expectedReview.ReviewId, review.ReviewId)
	assert.Equal(t, expectedReview.PostId, review.PostId)
	assert.Equal(t, expectedReview.Reviewer, review.Reviewer)
	assert.Equal(t, expectedReview.Rating, review.Rating)
}
