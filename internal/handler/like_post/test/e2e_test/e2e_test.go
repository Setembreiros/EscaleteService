package like_post_e2e_test

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

func TestLikePost(t *testing.T) {
	log.Info().Msg("Starting LikePost E2E test")
	e2e_test_common.SetUpE2E(t)
	defer e2e_test_common.TearDownE2E()
	user := &model.User{
		Username: "testuser123",
	}
	e2e_test_arrange.AddUser(t, user)
	user2 := &model.User{
		Username: "testuser456",
	}
	e2e_test_arrange.AddUser(t, user2)
	expectedPost := &model.Post{
		PostId:   "test-post-id",
		Username: user.Username,
	}
	e2e_test_arrange.AddPost(t, expectedPost)
	userLikedPostEvent := &event.UserLikedPostEvent{
		PostId:   expectedPost.PostId,
		Username: user2.Username,
	}
	expectedLikePost := &model.LikePost{
		Username: userLikedPostEvent.Username,
		PostId:   userLikedPostEvent.PostId,
	}

	e2e_test_action.PublishEvent(t, event.UserLikedPostEventName, userLikedPostEvent)

	e2e_test_assert.AssertLikePostExists(t, expectedLikePost)
	e2e_test_assert.AssertPostReactionScore(t, expectedPost.PostId, model.GetScore("like"))
	log.Info().Msg("LikePost E2E test Finished Successfully")
}
