package user_created_test

import (
	"errors"
	"escalateservice/internal/handler/user_created"
	mock_user_created "escalateservice/internal/handler/user_created/test/mock"
	model "escalateservice/internal/model/domain"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_user_created.MockRepository
var service *user_created.UserCreatedService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_user_created.NewMockRepository(ctrl)
	log.Logger = log.Output(&loggerOutput)
	service = user_created.NewUserCreatedService(serviceRepository)
}

func TestAddUserProfileWithService(t *testing.T) {
	setUpService(t)
	data := &model.User{
		Username: "username1",
	}
	serviceRepository.EXPECT().AddUser(data)

	service.AddUser(data)

	assert.Contains(t, loggerOutput.String(), "User username1 was added")
}

func TestAddUserWithService_Error(t *testing.T) {
	setUpService(t)
	data := &model.User{
		Username: "username1",
	}
	serviceRepository.EXPECT().AddUser(data).Return(errors.New("some error"))

	service.AddUser(data)

	assert.Contains(t, loggerOutput.String(), "Error adding user username1")
}
