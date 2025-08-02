package create_post_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/post_created"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"testing"
)

var db *database.Database
var handler *post_created.PostWasCreatedEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := post_created.NewPostCreatedRepository(db)
	service := post_created.NewPostCreatedService(repository)
	handler = post_created.NewPostWasCreatedEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreatePost_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	user := &model.User{
		Username: "username1",
	}
	integration_test_arrange.AddUser(t, user)
	post := &event.PostWasCreatedEvent{
		PostId: "post1",
		Metadata: event.Metadata{
			Username: "username1",
		},
	}
	data, _ := test_common.SerializeData(post)
	expectedPost := &model.Post{
		PostId:   post.PostId,
		Username: post.Metadata.Username,
	}

	handler.Handle(data)

	integration_test_assert.AssertPostExists(t, db, post.PostId, expectedPost)
}
