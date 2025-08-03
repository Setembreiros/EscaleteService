package superlike_post_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/superlike_post"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository superlike_post.SuperlikePostRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *superlike_post.NewSuperlikePostRepository(database.NewDatabase(client))
}

func TestAddSuperlikePostInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	client.EXPECT().AddSuperlikePost(data)

	repository.AddSuperlikePost(data)
}
