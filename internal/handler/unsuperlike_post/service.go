package unsuperlike_post

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	RemoveSuperlikePost(data *model.SuperlikePost) error
}

type UnsuperlikePostService struct {
	repository Repository
}

func NewUnsuperlikePostService(repository Repository) *UnsuperlikePostService {
	return &UnsuperlikePostService{
		repository: repository,
	}
}

func (s *UnsuperlikePostService) RemoveSuperlikePost(data *model.SuperlikePost) {
	err := s.repository.RemoveSuperlikePost(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error removing superlikePost, username %s -> post %s", data.Username, data.PostId)
		return
	}

	log.Info().Msgf("SuperlikePost was removed, username %s -> post %s", data.Username, data.PostId)
}
