package post_created_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/post_created"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository post_created.PostCreatedRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *post_created.NewPostCreatedRepository(database.NewDatabase(client))
}

func TestAddPostInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.Post{
		PostId:   "post1",
		Username: "username1",
	}
	client.EXPECT().AddPost(data)

	repository.AddPost(data)
}
