package follow_user_test

import (
	"errors"
	"escalateservice/internal/handler/follow_user"
	mock_follow_user "escalateservice/internal/handler/follow_user/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_follow_user.MockRepository
var service *follow_user.FollowService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_follow_user.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = follow_user.NewFollowService(serviceRepository)
}

func TestAddFollowWithService(t *testing.T) {
	setUpService(t)
	data := &model.Follow{
		Follower: "username1",
		Followee: "username2",
	}
	serviceRepository.EXPECT().AddFollow(data)

	service.AddFollow(data)

	assert.Contains(t, loggerOutput.String(), "Follow was added, username1 -> username2")
}

func TestAddFollowWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.Follow{
		Follower: "username1",
		Followee: "username2",
	}
	serviceRepository.EXPECT().AddFollow(data).Return(errors.New("some error"))

	service.AddFollow(data)

	assert.Contains(t, loggerOutput.String(), "Error adding follow, username1 -> username2")
}
