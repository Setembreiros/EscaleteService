package like_post_test

import (
	"encoding/json"
	"escalateservice/internal/handler/like_post"
	mock_like_post "escalateservice/internal/handler/like_post/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_like_post.MockService
var handler *like_post.UserLikedPostEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_like_post.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = like_post.NewUserLikedPostEventHandler(serviceHandler)
}

func TestHandleUserLikedPostEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.UserLikedPostEvent{
		PostId:   "post1",
		Username: "username1",
	}
	event, _ := test_common.SerializeData(data)
	expectedLikePost := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceHandler.EXPECT().AddLikePost(expectedLikePost)

	handler.Handle(event)
}

func TestHandleUserLikedPostEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
