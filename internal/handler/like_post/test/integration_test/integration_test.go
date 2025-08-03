package like_post_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/like_post"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *like_post.UserLikedPostEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := like_post.NewLikePostRepository(db)
	service := like_post.NewLikePostService(repository)
	handler = like_post.NewUserLikedPostEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateLikePost_WhenDatabaseReturnsSuccess(t *testing.T) {
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
	likePost := &event.UserLikedPostEvent{
		Username: user.Username,
		PostId:   post.PostId,
	}
	data, _ := test_common.SerializeData(likePost)
	expectedLikePost := &model.LikePost{
		PostId:   likePost.PostId,
		Username: likePost.Username,
	}

	handler.Handle(data)

	integration_test_assert.AssertLikePostExists(t, db, expectedLikePost)
}
