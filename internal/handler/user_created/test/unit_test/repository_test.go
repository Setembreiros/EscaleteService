package user_created_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/user_created"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository user_created.UserCreatedRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *user_created.NewUserCreatedRepository(database.NewDatabase(client))
}

func TestAddUserInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.User{
		Username: "username1",
	}
	client.EXPECT().AddUser(data)

	repository.AddUser(data)
}
