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
