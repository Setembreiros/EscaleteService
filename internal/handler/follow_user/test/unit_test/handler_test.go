package follow_user_test

import (
	"encoding/json"
	"escalateservice/internal/handler/follow_user"
	mock_follow_user "escalateservice/internal/handler/follow_user/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_follow_user.MockService
var handler *follow_user.UserAFollowedUserBEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_follow_user.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = follow_user.NewUserAFollowedUserBEventHandler(serviceHandler)
}

func TestHandleUserAFollowedUserBEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.UserAFollowedUserBEvent{
		Follower: "username1",
		Followee: "username2",
	}
	event, _ := test_common.SerializeData(data)
	expectedLikePost := &model.Follow{
		Follower: "username1",
		Followee: "username2",
	}
	serviceHandler.EXPECT().AddFollow(expectedLikePost)

	handler.Handle(event)
}

func TestHandleUserAFollowedUserBEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
