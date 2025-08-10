package unfollow_user_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/unfollow_user"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *unfollow_user.UserAUnfollowedUserBEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := unfollow_user.NewUnfollowRepository(db)
	service := unfollow_user.NewUnfollowService(repository)
	handler = unfollow_user.NewUserAUnfollowedUserBEventHandler(service)
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
	follow := &model.Follow{
		Follower: follower.Username,
		Followee: followee.Username,
	}
	integration_test_arrange.AddFollow(t, follow)
	userAUnfollowedUserBEvent := &event.UserAUnfollowedUserBEvent{
		Follower: follower.Username,
		Followee: followee.Username,
	}
	data, _ := test_common.SerializeData(userAUnfollowedUserBEvent)

	handler.Handle(data)

	integration_test_assert.AssertFollowDoesNotExist(t, db, follow)
}
