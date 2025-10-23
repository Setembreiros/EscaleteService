package review_created_test

import (
	"errors"
	"escalateservice/internal/handler/review_created"
	mock_review_created "escalateservice/internal/handler/review_created/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_review_created.MockRepository
var service *review_created.ReviewCreatedService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_review_created.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = review_created.NewReviewCreatedService(serviceRepository)
}

func TestAddReviewProfileWithService(t *testing.T) {
	setUpService(t)
	data := &model.Review{
		ReviewId: 1,
		PostId:   "post1",
		Reviewer: "username1",
		Rating:   5,
	}
	serviceRepository.EXPECT().AddReview(data)

	service.AddReview(data)

	assert.Contains(t, loggerOutput.String(), "Review 1 was added")
}

func TestAddReviewWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.Review{
		ReviewId: 1,
		PostId:   "post1",
		Reviewer: "username1",
		Rating:   5,
	}
	serviceRepository.EXPECT().AddReview(data).Return(errors.New("some error"))

	service.AddReview(data)

	assert.Contains(t, loggerOutput.String(), "Error adding review 1")
}
