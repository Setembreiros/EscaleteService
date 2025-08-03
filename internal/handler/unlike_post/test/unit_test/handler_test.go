package unlike_post_test

import (
	"encoding/json"
	"escalateservice/internal/handler/unlike_post"
	mock_unlike_post "escalateservice/internal/handler/unlike_post/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_unlike_post.MockService
var handler *unlike_post.UserUnlikedPostEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_unlike_post.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = unlike_post.NewUserUnlikedPostEventHandler(serviceHandler)
}

func TestHandleUserUnlikedPostEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.UserUnlikedPostEvent{
		PostId:   "post1",
		Username: "username1",
	}
	event, _ := test_common.SerializeData(data)
	expectedLikePost := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceHandler.EXPECT().RemoveLikePost(expectedLikePost)

	handler.Handle(event)
}

func TestHandleUserUnlikedPostEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
