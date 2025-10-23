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
	user1 := createUser(t, "username1")
	user2 := createUser(t, "username2")
	post := &model.Post{
		PostId:   "post1",
		Username: user1.Username,
	}
	integration_test_arrange.AddPost(t, post)
	data1, expectedReview1 := createReview(user1.Username, 5)
	data2, expectedReview2 := createReview(user2.Username, 1)

	handler.Handle(data1)
	handler.Handle(data2)

	integration_test_assert.AssertReviewExists(t, db, expectedReview1.ReviewId, expectedReview1)
	integration_test_assert.AssertReviewExists(t, db, expectedReview2.ReviewId, expectedReview2)
	integration_test_assert.AssertPostReactionScore(t, db, post.PostId, model.GetScore(fmt.Sprintf("review%dstar", expectedReview1.Rating))+model.GetScore(fmt.Sprintf("review%dstar", expectedReview2.Rating)))
}

func createUser(t *testing.T, username string) *model.User {
	user := &model.User{
		Username: username,
	}
	integration_test_arrange.AddUser(t, user)

	return user
}

var lastReviewID uint64 = 0

func createReview(reviewer string, rating int) ([]byte, *model.Review) {
	lastReviewID = lastReviewID + 1
	review := &event.ReviewWasCreatedEvent{
		ReviewId: lastReviewID,
		PostId:   "post1",
		Username: reviewer,
		Rating:   rating,
	}
	data, _ := test_common.SerializeData(review)
	expectedReview := &model.Review{
		ReviewId: review.ReviewId,
		PostId:   review.PostId,
		Reviewer: review.Username,
		Rating:   review.Rating,
	}

	return data, expectedReview
}
