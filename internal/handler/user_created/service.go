package user_created

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	AddUser(data *model.User) error
}

type UserCreatedService struct {
	repository Repository
}

func NewUserCreatedService(repository Repository) *UserCreatedService {
	return &UserCreatedService{
		repository: repository,
	}
}

func (s *UserCreatedService) AddUser(data *model.User) {
	err := s.repository.AddUser(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error adding user %s", data.Username)
		return
	}

	log.Info().Msgf("User %s was added", data.Username)
}
