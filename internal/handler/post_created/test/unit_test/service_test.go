package post_created_test

import (
	"errors"
	"escalateservice/internal/handler/post_created"
	mock_post_created "escalateservice/internal/handler/post_created/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_post_created.MockRepository
var service *post_created.PostCreatedService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_post_created.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = post_created.NewPostCreatedService(serviceRepository)
}

func TestAddPostProfileWithService(t *testing.T) {
	setUpService(t)
	data := &model.Post{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().AddPost(data)

	service.AddPost(data)

	assert.Contains(t, loggerOutput.String(), "Post post1 was added")
}

func TestAddPostWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.Post{
		PostId:   "post1",
		Username: "username1",
	}
	serviceRepository.EXPECT().AddPost(data).Return(errors.New("some error"))

	service.AddPost(data)

	assert.Contains(t, loggerOutput.String(), "Error adding post post1")
}
