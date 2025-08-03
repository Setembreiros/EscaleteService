package user_created_e2e_test

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

func TestReviewCreated(t *testing.T) {
	log.Info().Msg("Starting ReviewCreated E2E test")
	e2e_test_common.SetUpE2E(t)
	defer e2e_test_common.TearDownE2E()
	user := &model.User{
		Username: "testuser123",
	}
	e2e_test_arrange.AddUser(t, user)
	reviewer := &model.User{
		Username: "reviewer123",
	}
	e2e_test_arrange.AddUser(t, reviewer)
	expectedPost := &model.Post{
		PostId:   "test-post-id",
		Username: user.Username,
	}
	e2e_test_arrange.AddPost(t, expectedPost)
	reviewWasCreatedEvent := &event.ReviewWasCreatedEvent{
		ReviewId: 1,
		PostId:   expectedPost.PostId,
		Username: reviewer.Username,
		Rating:   5,
	}
	expectedReview := &model.Review{
		ReviewId: reviewWasCreatedEvent.ReviewId,
		PostId:   reviewWasCreatedEvent.PostId,
		Reviewer: reviewWasCreatedEvent.Username,
		Rating:   reviewWasCreatedEvent.Rating,
	}

	e2e_test_action.PublishEvent(t, event.ReviewWasCreatedEventName, reviewWasCreatedEvent)

	e2e_test_assert.AssertReviewExists(t, reviewWasCreatedEvent.ReviewId, expectedReview)
	log.Info().Msg("ReviewCreated E2E test Finished Successfully")
}
