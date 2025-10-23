package unsuperlike_post_test

import (
	"errors"
	"escalateservice/internal/handler/unsuperlike_post"
	mock_unsuperlike_post "escalateservice/internal/handler/unsuperlike_post/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_unsuperlike_post.MockRepository
var service *unsuperlike_post.UnsuperlikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_unsuperlike_post.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = unsuperlike_post.NewUnsuperlikePostService(serviceRepository)
}

func TestRemoveSuperlikePostWithService(t *testing.T) {
	setUpService(t)
	data := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().RemoveSuperlikePost(data)

	service.RemoveSuperlikePost(data)

	assert.Contains(t, loggerOutput.String(), "SuperlikePost was removed, username username1 -> post post1")
}

func TestRemoveSuperlikePostWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.SuperlikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().RemoveSuperlikePost(data).Return(errors.New("some error"))

	service.RemoveSuperlikePost(data)

	assert.Contains(t, loggerOutput.String(), "Error removing superlikePost, username username1 -> post post1")
}
