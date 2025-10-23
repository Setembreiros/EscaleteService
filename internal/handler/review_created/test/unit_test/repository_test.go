package review_created_test

import (
	database "escalateservice/internal/db"
	mock_database "escalateservice/internal/db/test/mock"
	"escalateservice/internal/handler/review_created"
	model "escalateservice/internal/model/domain"
	"testing"
)

var client *mock_database.MockDatabaseClient
var repository review_created.ReviewCreatedRepository

func setUpRepository(t *testing.T) {
	setUp(t)
	client = mock_database.NewMockDatabaseClient(ctrl)
	repository = *review_created.NewReviewCreatedRepository(database.NewDatabase(client))
}

func TestAddReviewInRepository(t *testing.T) {
	setUpRepository(t)
	data := &model.Review{
		ReviewId: 1,
		PostId:   "post1",
		Reviewer: "username1",
		Rating:   5,
	}
	client.EXPECT().AddReview(data)

	repository.AddReview(data)
}
