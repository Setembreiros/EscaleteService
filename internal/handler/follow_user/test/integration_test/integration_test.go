package follow_user_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/follow_user"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *follow_user.UserAFollowedUserBEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := follow_user.NewFollowRepository(db)
	service := follow_user.NewFollowService(repository)
	handler = follow_user.NewUserAFollowedUserBEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateFollow_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	follower := &model.User{
		Username: "username1",
	}
	integration_test_arrange.AddUser(t, follower)
	followee := &model.User{
		Username: "username2",
	}
	integration_test_arrange.AddUser(t, followee)
	follow := &event.UserAFollowedUserBEvent{
		Follower: follower.Username,
		Followee: followee.Username,
	}
	data, _ := test_common.SerializeData(follow)
	expectedFollow := &model.Follow{
		Follower: follower.Username,
		Followee: followee.Username,
	}

	handler.Handle(data)

	integration_test_assert.AssertFollowExists(t, db, expectedFollow)
}
