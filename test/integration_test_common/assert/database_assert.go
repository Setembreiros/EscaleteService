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
