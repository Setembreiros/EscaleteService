package unlike_post_e2e_test

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

func TestUnlikePost(t *testing.T) {
	log.Info().Msg("Starting UnlikePost E2E test")
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
	post := &model.Post{
		PostId:   "test-post-id",
		Username: user.Username,
	}
	e2e_test_arrange.AddPost(t, post)
	likePost := &model.LikePost{
		Username: user2.Username,
		PostId:   post.PostId,
	}
	e2e_test_arrange.AddLikePost(t, likePost)
	userUnlikedPostEvent := &event.UserUnlikedPostEvent{
		Username: likePost.Username,
		PostId:   likePost.PostId,
	}

	e2e_test_action.PublishEvent(t, event.UserUnlikedPostEventName, userUnlikedPostEvent)

	e2e_test_assert.AssertLikePostDoesNotExist(t, likePost)
	log.Info().Msg("UnlikePost E2E test Finished Successfully")
}
