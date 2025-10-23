package unsuperlike_post_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/unsuperlike_post"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *unsuperlike_post.UserUnsuperlikedPostEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := unsuperlike_post.NewUnsuperlikePostRepository(db)
	service := unsuperlike_post.NewUnsuperlikePostService(repository)
	handler = unsuperlike_post.NewUserUnsuperlikedPostEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateUnsuperlikePost_WhenDatabaseReturnsSuccess(t *testing.T) {
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
	superlikePost := &model.SuperlikePost{
		Username: user.Username,
		PostId:   post.PostId,
	}
	integration_test_arrange.AddSuperlikePost(t, superlikePost)
	unsuperlikePost := &event.UserUnsuperlikedPostEvent{
		Username: user.Username,
		PostId:   post.PostId,
	}
	data, _ := test_common.SerializeData(unsuperlikePost)

	handler.Handle(data)

	integration_test_assert.AssertSuperlikePostDoesNotExist(t, db, superlikePost)
	integration_test_assert.AssertPostReactionScore(t, db, post.PostId, 0)
}
