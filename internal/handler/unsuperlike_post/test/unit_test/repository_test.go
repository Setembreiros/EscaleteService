package unsuperlike_post_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/unsuperlike_post"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository unsuperlike_post.UnsuperlikePostRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *unsuperlike_post.NewUnsuperlikePostRepository(database.NewDatabase(client))
}

func TestRemoveSuperlikePostInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	client.EXPECT().RemoveSuperlikePost(data)

	repository.RemoveSuperlikePost(data)
}
