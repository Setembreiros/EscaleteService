package user_created_e2e_test

import (
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/e2e_test_common"
	e2e_test_action "escalateservice/test/e2e_test_common/action"
	e2e_test_assert "escalateservice/test/e2e_test_common/assert"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestUserCreated(t *testing.T) {
	log.Info().Msg("Starting UserCreated E2E test")
	e2e_test_common.SetUpE2E(t)
	defer e2e_test_common.TearDownE2E()
	userWasRegisteredEvent := &event.UserWasRegisteredEvent{
		Username: "testuser123",
	}
	expectedUser := &model.User{
		Username: userWasRegisteredEvent.Username,
	}

	e2e_test_action.PublishEvent(t, event.UserWasRegisteredEventName, userWasRegisteredEvent)

	e2e_test_assert.AssertUserExists(t, userWasRegisteredEvent.Username, expectedUser)
	log.Info().Msg("UserCreated E2E test Finished Successfully")
}
