package like_post_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/like_post"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository like_post.LikePostRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *like_post.NewLikePostRepository(database.NewDatabase(client))
}

func TestAddLikePostInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	client.EXPECT().AddLikePost(data)

	repository.AddLikePost(data)
}
