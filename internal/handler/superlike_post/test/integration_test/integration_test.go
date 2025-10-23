package superlike_post_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/superlike_post"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *superlike_post.UserSuperlikedPostEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := superlike_post.NewSuperlikePostRepository(db)
	service := superlike_post.NewSuperlikePostService(repository)
	handler = superlike_post.NewUserSuperlikedPostEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateSuperlikePost_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	user := &model.User{
		Username: "username1",
	}
	integration_test_arrange.AddUser(t, user)
	post := &model.Post{
		PostId:   "post1",
		Username: "username1",
	}
	integration_test_arrange.AddPost(t, post)
	userSuperlikedPostEvent := &event.UserSuperlikedPostEvent{
		Username: user.Username,
		PostId:   post.PostId,
	}
	data, _ := test_common.SerializeData(userSuperlikedPostEvent)
	expectedSuperlikePost := &model.SuperlikePost{
		PostId:   userSuperlikedPostEvent.PostId,
		Username: userSuperlikedPostEvent.Username,
	}

	handler.Handle(data)

	integration_test_assert.AssertSuperlikePostExists(t, db, expectedSuperlikePost)
	integration_test_assert.AssertPostReactionScore(t, db, post.PostId, model.GetScore("superlike"))
}
