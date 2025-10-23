package unfollow_user_test

import (
	"errors"
	"escalateservice/internal/handler/unfollow_user"
	mock_unfollow_user "escalateservice/internal/handler/unfollow_user/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_unfollow_user.MockRepository
var service *unfollow_user.UnfollowService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_unfollow_user.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = unfollow_user.NewUnfollowService(serviceRepository)
}

func TestRemoveFollowWithService(t *testing.T) {
	setUpService(t)
	data := &model.Follow{
		Follower: "username1",
		Followee: "username2",
	}
	serviceRepository.EXPECT().RemoveFollow(data)

	service.RemoveFollow(data)

	assert.Contains(t, loggerOutput.String(), "Follow was removed, username1 -> username2")
}

func TestRemoveFollowWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.Follow{
		Follower: "username1",
		Followee: "username2",
	}
	serviceRepository.EXPECT().RemoveFollow(data).Return(errors.New("some error"))

	service.RemoveFollow(data)

	assert.Contains(t, loggerOutput.String(), "Error removing follow, username1 -> username2")
}
