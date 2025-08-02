package post_created_test

import (
	"encoding/json"
	"escalateservice/internal/handler/post_created"
	mock_post_created "escalateservice/internal/handler/post_created/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_post_created.MockService
var handler *post_created.PostWasCreatedEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_post_created.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = post_created.NewPostWasCreatedEventHandler(serviceHandler)
}

func TestHandlePostWasCreatedEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.PostWasCreatedEvent{
		PostId: "post1",
		Metadata: event.Metadata{
			Username: "username1",
		},
	}
	event, _ := test_common.SerializeData(data)
	expectedPostprofile := &model.Post{
		PostId:   "post1",
		Username: "username1",
	}
	serviceHandler.EXPECT().AddPost(expectedPostprofile)

	handler.Handle(event)
}

func TestHandlePostWasCreatedEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
