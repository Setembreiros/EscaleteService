package unsuperlike_post_test

import (
	"encoding/json"
	"escalateservice/internal/handler/unsuperlike_post"
	mock_unsuperlike_post "escalateservice/internal/handler/unsuperlike_post/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_unsuperlike_post.MockService
var handler *unsuperlike_post.UserUnsuperlikedPostEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_unsuperlike_post.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = unsuperlike_post.NewUserUnsuperlikedPostEventHandler(serviceHandler)
}

func TestHandleUserUnsuperlikedPostEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.UserUnsuperlikedPostEvent{
		PostId:   "post1",
		Username: "username1",
	}
	event, _ := test_common.SerializeData(data)
	expectedSuperlikePost := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceHandler.EXPECT().RemoveSuperlikePost(expectedSuperlikePost)

	handler.Handle(event)
}

func TestHandleUserUnsuperlikedPostEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
