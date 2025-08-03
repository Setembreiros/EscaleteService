package like_post_test

import (
	"errors"
	"escalateservice/internal/handler/like_post"
	mock_like_post "escalateservice/internal/handler/like_post/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_like_post.MockRepository
var service *like_post.LikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_like_post.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = like_post.NewLikePostService(serviceRepository)
}

func TestAddLikePostWithService(t *testing.T) {
	setUpService(t)
	data := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().AddLikePost(data)

	service.AddLikePost(data)

	assert.Contains(t, loggerOutput.String(), "LikePost was added, username username1 -> post post1")
}

func TestAddLikePostWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.LikePost{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().AddLikePost(data).Return(errors.New("some error"))

	service.AddLikePost(data)

	assert.Contains(t, loggerOutput.String(), "Error adding likePost, username username1 -> post post1")
}
