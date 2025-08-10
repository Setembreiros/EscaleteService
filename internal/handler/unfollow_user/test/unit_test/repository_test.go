package unfollow_user_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/unfollow_user"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository unfollow_user.UnfollowRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *unfollow_user.NewUnfollowRepository(database.NewDatabase(client))
}

func TestRemoveFollowInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.Follow{
		Follower: "username1",
		Followee: "username2",
	}
	client.EXPECT().RemoveFollow(data)

	repository.RemoveFollow(data)
}
