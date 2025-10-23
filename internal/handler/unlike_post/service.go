package unlike_post

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	RemoveLikePost(data *model.LikePost) error
}

type UnlikePostService struct {
	repository Repository
}

func NewUnlikePostService(repository Repository) *UnlikePostService {
	return &UnlikePostService{
		repository: repository,
	}
}

func (s *UnlikePostService) RemoveLikePost(data *model.LikePost) {
	err := s.repository.RemoveLikePost(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error removing likePost, username %s -> post %s", data.Username, data.PostId)
		return
	}

	log.Info().Msgf("LikePost was removed, username %s -> post %s", data.Username, data.PostId)
}
