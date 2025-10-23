package follow_user_e2e_test

import (
	model "escalateservice/internal/model/domain"
	"escalateservice/internal/model/event"
	"escalateservice/test/e2e_test_common"
	e2e_test_action "escalateservice/test/e2e_test_common/action"
	e2e_test_arrange "escalateservice/test/e2e_test_common/arrange"
	e2e_test_assert "escalateservice/test/e2e_test_common/assert"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestFollow(t *testing.T) {
	log.Info().Msg("Starting Follow E2E test")
	e2e_test_common.SetUpE2E(t)
	defer e2e_test_common.TearDownE2E()
	follower := &model.User{
		Username: "testuser123",
	}
	e2e_test_arrange.AddUser(t, follower)
	followee := &model.User{
		Username: "testuser456",
	}
	e2e_test_arrange.AddUser(t, followee)
	userAFollowedUserBEvent := &event.UserAFollowedUserBEvent{
		Follower: follower.Username,
		Followee: followee.Username,
	}
	expectedFollow := &model.Follow{
		Follower: userAFollowedUserBEvent.Follower,
		Followee: userAFollowedUserBEvent.Followee,
	}

	e2e_test_action.PublishEvent(t, event.UserAFollowedUserBEventName, userAFollowedUserBEvent)

	e2e_test_assert.AssertFollowExists(t, expectedFollow)
	log.Info().Msg("Follow E2E test Finished Successfully")
}
