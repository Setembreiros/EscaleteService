package unlike_post_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/unlike_post"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *unlike_post.UserUnlikedPostEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := unlike_post.NewUnlikePostRepository(db)
	service := unlike_post.NewUnlikePostService(repository)
	handler = unlike_post.NewUserUnlikedPostEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateUnlikePost_WhenDatabaseReturnsSuccess(t *testing.T) {
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
	likePost := &model.LikePost{
		Username: user.Username,
		PostId:   post.PostId,
	}
	integration_test_arrange.AddLikePost(t, likePost)
	unlikePost := &event.UserUnlikedPostEvent{
		Username: user.Username,
		PostId:   post.PostId,
	}
	data, _ := test_common.SerializeData(unlikePost)

	handler.Handle(data)

	integration_test_assert.AssertLikePostDoesNotExist(t, db, likePost)
}
