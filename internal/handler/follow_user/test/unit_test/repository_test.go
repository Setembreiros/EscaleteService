package follow_user_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/follow_user"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository follow_user.FollowRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *follow_user.NewFollowRepository(database.NewDatabase(client))
}

func TestAddFollowInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.Follow{
		Follower: "username1",
		Followee: "username2",
	}
	client.EXPECT().AddFollow(data)

	repository.AddFollow(data)
}
