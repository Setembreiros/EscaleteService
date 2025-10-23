package create_review_integration_test

import (
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/review_created"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"escalateservice/test/test_common"
	"fmt"
	"testing"
)

var db *database.Database
var handler *review_created.ReviewWasCreatedEventHandler

func setUp(t *testing.T) {
	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase()
	repository := review_created.NewReviewCreatedRepository(db)
	service := review_created.NewReviewCreatedService(repository)
	handler = review_created.NewReviewWasCreatedEventHandler(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateReview_WhenDatabaseReturnsSuccess(t *testing.T) {
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
	review := &event.ReviewWasCreatedEvent{
		ReviewId: 1,
		PostId:   "post1",
		Username: "username1",
		Rating:   5,
	}
	data, _ := test_common.SerializeData(review)
	expectedReview := &model.Review{
		ReviewId: review.ReviewId,
		PostId:   review.PostId,
		Reviewer: review.Username,
		Rating:   review.Rating,
	}

	handler.Handle(data)

	integration_test_assert.AssertReviewExists(t, db, review.ReviewId, expectedReview)
	integration_test_assert.AssertPostReactionScore(t, db, post.PostId, model.GetScore(fmt.Sprintf("review%dstar", review.Rating)))
}
