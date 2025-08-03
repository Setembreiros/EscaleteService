package superlike_post_test

import (
	"encoding/json"
	"escalateservice/internal/handler/superlike_post"
	mock_superlike_post "escalateservice/internal/handler/superlike_post/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_superlike_post.MockService
var handler *superlike_post.UserSuperlikedPostEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_superlike_post.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = superlike_post.NewUserSuperlikedPostEventHandler(serviceHandler)
}

func TestHandleUserSuperlikedPostEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.UserSuperlikedPostEvent{
		PostId:   "post1",
		Username: "username1",
	}
	event, _ := test_common.SerializeData(data)
	expectedSuperlikePost := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceHandler.EXPECT().AddSuperlikePost(expectedSuperlikePost)

	handler.Handle(event)
}

func TestHandleUserSuperlikedPostEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
