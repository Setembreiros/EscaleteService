package review_created_test

import (
	"encoding/json"
	"escalateservice/internal/handler/review_created"
	mock_review_created "escalateservice/internal/handler/review_created/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_review_created.MockService
var handler *review_created.ReviewWasCreatedEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_review_created.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = review_created.NewReviewWasCreatedEventHandler(serviceHandler)
}

func TestHandleReviewWasCreatedEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.ReviewWasCreatedEvent{
		ReviewId: 1,
		PostId:   "post1",
		Username: "username1",
		Rating:   5,
	}
	event, _ := test_common.SerializeData(data)
	expectedReviewprofile := &model.Review{
		ReviewId: 1,
		PostId:   "post1",
		Reviewer: "username1",
		Rating:   5,
	}
	serviceHandler.EXPECT().AddReview(expectedReviewprofile)

	handler.Handle(event)
}

func TestHandleReviewWasCreatedEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
