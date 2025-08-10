package integration_test_assert

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertUserExists(t *testing.T, db *database.Database, username string, expectedUser *model.User) {
	user, err := db.Client.GetUser(username)
	assert.Nil(t, err, "Error should be nil when checking for existing user")
	assert.Equal(t, expectedUser.Username, user.Username, "Username does not match expected values")
}

func AssertPostExists(t *testing.T, db *database.Database, postId string, expectedPost *model.Post) {
	post, err := db.Client.GetPost(postId)
	assert.Nil(t, err, "Error should be nil when checking for existing post")
	assert.Equal(t, expectedPost.PostId, post.PostId, "PostId does not match expected values")
	assert.Equal(t, expectedPost.Username, post.Username, "Post username does not match expected values")
}

func AssertReviewExists(t *testing.T, db *database.Database, reviewId uint64, expectedReview *model.Review) {
	review, err := db.Client.GetReview(reviewId)
	assert.Nil(t, err, "Error should be nil when checking for existing review")
	assert.Equal(t, expectedReview.ReviewId, review.ReviewId, "ReviewId does not match expected values")
	assert.Equal(t, expectedReview.PostId, review.PostId, "PostId does not match expected values")
	assert.Equal(t, expectedReview.Reviewer, review.Reviewer, "Reviewer does not match expected values")
	assert.Equal(t, expectedReview.Rating, review.Rating, "Review does not match expected values")
}

func AssertLikePostExists(t *testing.T, db *database.Database, expectedLikePost *model.LikePost) {
	likePost, err := db.Client.GetLikePost(expectedLikePost.Username, expectedLikePost.PostId)
	assert.Nil(t, err)
	assert.Equal(t, expectedLikePost.Username, likePost.Username, "Like post username does not match expected values")
	assert.Equal(t, expectedLikePost.PostId, likePost.PostId, "Like post does not match expected values")
}

func AssertLikePostDoesNotExist(t *testing.T, db *database.Database, expectedLikePost *model.LikePost) {
	likePost, err := db.Client.GetLikePost(expectedLikePost.Username, expectedLikePost.PostId)
	assert.Nil(t, err, "Error should be nil when checking for non-existing like post")
	assert.Nil(t, likePost, "Like post should not exist")
}

func AssertSuperlikePostExists(t *testing.T, db *database.Database, expectedSuperlike *model.SuperlikePost) {
	likePost, err := db.Client.GetSuperlikePost(expectedSuperlike.Username, expectedSuperlike.PostId)
	assert.Nil(t, err, "Error should be nil when checking for existing superlike post")
	assert.Equal(t, expectedSuperlike.Username, likePost.Username, "Superlike post username does not match expected values")
	assert.Equal(t, expectedSuperlike.PostId, likePost.PostId, "Superlike post does not match expected values")
}

func AssertSuperlikePostDoesNotExist(t *testing.T, db *database.Database, expectedSuperlikePost *model.SuperlikePost) {
	superlikePost, err := db.Client.GetLikePost(expectedSuperlikePost.Username, expectedSuperlikePost.PostId)
	assert.Nil(t, err)
	assert.Nil(t, superlikePost, "Superlike post should not exist")
}

func AssertFollowExists(t *testing.T, db *database.Database, expectedFollow *model.Follow) {
	follow, err := db.Client.GetFollow(expectedFollow.Follower, expectedFollow.Followee)
	assert.Nil(t, err)
	assert.Equal(t, expectedFollow.Follower, follow.Follower, "Follower does not match expected values")
	assert.Equal(t, expectedFollow.Followee, follow.Followee, "Followee does not match expected values")
}

func AssertFollowDoesNotExist(t *testing.T, db *database.Database, expectedFollow *model.Follow) {
	follow, err := db.Client.GetFollow(expectedFollow.Follower, expectedFollow.Followee)
	assert.Nil(t, err)
	assert.Nil(t, follow, "Follow should not exist")
}

func AssertUserScore(t *testing.T, db *database.Database, username string, expectedScore float64) {
	user, err := db.Client.GetUser(username)
	assert.Nil(t, err, "Error should be nil when checking user score")
	assert.NotNil(t, user, "User should not be nil")
	assert.Equal(t, expectedScore, user.Score, "User score does not match expected value")
}

func AssertPostScore(t *testing.T, db *database.Database, postId string, expectedScore int) {
	post, err := db.Client.GetPost(postId)
	assert.Nil(t, err, "Error should be nil when checking post score")
	assert.NotNil(t, post, "Post should not be nil")
	assert.Equal(t, expectedScore, post.Score, "Post score does not match expected value")
}
