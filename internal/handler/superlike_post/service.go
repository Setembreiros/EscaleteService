package superlike_post

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	AddSuperlikePost(data *model.SuperlikePost) error
}

type SuperlikePostService struct {
	repository Repository
}

func NewSuperlikePostService(repository Repository) *SuperlikePostService {
	return &SuperlikePostService{
		repository: repository,
	}
}

func (s *SuperlikePostService) AddSuperlikePost(data *model.SuperlikePost) {
	err := s.repository.AddSuperlikePost(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error adding superlikePost, username %s -> post %s", data.Username, data.PostId)
		return
	}

	log.Info().Msgf("SuperlikePost was added, username %s -> post %s", data.Username, data.PostId)
}
