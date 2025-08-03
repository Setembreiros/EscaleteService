package superlike_post_test

import (
	"errors"
	"escalateservice/internal/handler/superlike_post"
	mock_superlike_post "escalateservice/internal/handler/superlike_post/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_superlike_post.MockRepository
var service *superlike_post.SuperlikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_superlike_post.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = superlike_post.NewSuperlikePostService(serviceRepository)
}

func TestAddSuperlikePostWithService(t *testing.T) {
	setUpService(t)
	data := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().AddSuperlikePost(data)

	service.AddSuperlikePost(data)

	assert.Contains(t, loggerOutput.String(), "SuperlikePost was added, username username1 -> post post1")
}

func TestAddSuperlikePostWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().AddSuperlikePost(data).Return(errors.New("some error"))

	service.AddSuperlikePost(data)

	assert.Contains(t, loggerOutput.String(), "Error adding superlikePost, username username1 -> post post1")
}
