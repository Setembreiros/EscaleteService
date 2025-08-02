package user_created_test

import (
	"encoding/json"
	"escalateservice/internal/handler/user_created"
	mock_user_created "escalateservice/internal/handler/user_created/test/mock"
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/test_common"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceHandler *mock_user_created.MockService
var handler *user_created.UserWasRegisteredEventHandler

func setUpHandler(t *testing.T) {
	setUp(t)
	serviceHandler = mock_user_created.NewMockService(ctrl)
	log.Logger = log.Output(&loggerOutput)
	handler = user_created.NewUserWasRegisteredEventHandler(serviceHandler)
}

func TestHandleUserWasRegisteredEvent(t *testing.T) {
	setUpHandler(t)
	data := &event.UserWasRegisteredEvent{
		Username: "username1",
	}
	event, _ := test_common.SerializeData(data)
	expectedUserprofile := &model.User{
		Username: "username1",
	}
	serviceHandler.EXPECT().AddUser(expectedUserprofile)

	handler.Handle(event)
}

func TestHandleUserWasRegisteredEvent_InvalidDataError(t *testing.T) {
	setUpHandler(t)
	invalidData := "invalid data"
	event, _ := json.Marshal(invalidData)

	handler.Handle(event)

	assert.Contains(t, loggerOutput.String(), "Invalid event data")
}
