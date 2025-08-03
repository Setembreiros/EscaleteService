package user_created_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/user_created"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *user_created.UserWasRegisteredEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := user_created.NewUserCreatedRepository(db)
	service := user_created.NewUserCreatedService(repository)
	handler = user_created.NewUserWasRegisteredEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateUser_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	user := &event.UserWasRegisteredEvent{
		Username: "usernameA",
	}
	data, _ := test_common.SerializeData(user)
	expectedUser := &model.User{
		Username: user.Username,
	}

	handler.Handle(data)

	integration_test_assert.AssertUserExists(t, db, user.Username, expectedUser)
}
