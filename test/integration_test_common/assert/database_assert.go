package integration_test_assert

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertUserExists(t *testing.T, db *database.Database, username string, expectedUser *model.User) {
	user, err := db.Client.GetUser(username)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Username, user.Username)
}

func AssertPostExists(t *testing.T, db *database.Database, postId string, expectedPost *model.Post) {
	post, err := db.Client.GetPost(postId)
	assert.Nil(t, err)
	assert.Equal(t, expectedPost.PostId, post.PostId)
	assert.Equal(t, expectedPost.Username, post.Username)
}

func AssertReviewExists(t *testing.T, db *database.Database, reviewId uint64, expectedReview *model.Review) {
	review, err := db.Client.GetReview(reviewId)
	assert.Nil(t, err)
	assert.Equal(t, expectedReview.ReviewId, review.ReviewId)
	assert.Equal(t, expectedReview.PostId, review.PostId)
	assert.Equal(t, expectedReview.Reviewer, review.Reviewer)
	assert.Equal(t, expectedReview.Rating, review.Rating)
}

func AssertLikePostExists(t *testing.T, db *database.Database, expectedLikePost *model.LikePost) {
	likePost, err := db.Client.GetLikePost(expectedLikePost.Username, expectedLikePost.PostId)
	assert.Nil(t, err)
	assert.Equal(t, expectedLikePost.Username, likePost.Username)
	assert.Equal(t, expectedLikePost.PostId, likePost.PostId)
}

func AssertLikePostDoesNotExist(t *testing.T, db *database.Database, expectedLikePost *model.LikePost) {
	likePost, err := db.Client.GetLikePost(expectedLikePost.Username, expectedLikePost.PostId)
	assert.Nil(t, err)
	assert.Nil(t, likePost)
}

func AssertSuperlikePostExists(t *testing.T, db *database.Database, expectedSuperlike *model.SuperlikePost) {
	likePost, err := db.Client.GetSuperlikePost(expectedSuperlike.Username, expectedSuperlike.PostId)
	assert.Nil(t, err)
	assert.Equal(t, expectedSuperlike.Username, likePost.Username)
	assert.Equal(t, expectedSuperlike.PostId, likePost.PostId)
}

func AssertSuperlikePostDoesNotExist(t *testing.T, db *database.Database, expectedSuperlikePost *model.SuperlikePost) {
	superlikePost, err := db.Client.GetLikePost(expectedSuperlikePost.Username, expectedSuperlikePost.PostId)
	assert.Nil(t, err)
	assert.Nil(t, superlikePost)
}
