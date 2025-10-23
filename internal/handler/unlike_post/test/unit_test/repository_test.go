package unlike_post_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/unlike_post"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository unlike_post.UnlikePostRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *unlike_post.NewUnlikePostRepository(database.NewDatabase(client))
}

func TestRemoveLikePostInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	client.EXPECT().RemoveLikePost(data)

	repository.RemoveLikePost(data)
}
