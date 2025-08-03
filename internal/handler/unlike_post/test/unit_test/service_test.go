package unlike_post_test

import (
	"errors"
	"escalateservice/internal/handler/unlike_post"
	mock_unlike_post "escalateservice/internal/handler/unlike_post/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_unlike_post.MockRepository
var service *unlike_post.UnlikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_unlike_post.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = unlike_post.NewUnlikePostService(serviceRepository)
}

func TestRemoveLikePostWithService(t *testing.T) {
	setUpService(t)
	data := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().RemoveLikePost(data)

	service.RemoveLikePost(data)

	assert.Contains(t, loggerOutput.String(), "LikePost was removed, username username1 -> post post1")
}

func TestRemoveLikePostWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().RemoveLikePost(data).Return(errors.New("some error"))

	service.RemoveLikePost(data)

	assert.Contains(t, loggerOutput.String(), "Error removing likePost, username username1 -> post post1")
}
