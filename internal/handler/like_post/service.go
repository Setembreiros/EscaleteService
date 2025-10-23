package like_post

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	AddLikePost(data *model.LikePost) error
}

type LikePostService struct {
	repository Repository
}

func NewLikePostService(repository Repository) *LikePostService {
	return &LikePostService{
		repository: repository,
	}
}

func (s *LikePostService) AddLikePost(data *model.LikePost) {
	err := s.repository.AddLikePost(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error adding likePost, username %s -> post %s", data.Username, data.PostId)
		return
	}

	log.Info().Msgf("LikePost was added, username %s -> post %s", data.Username, data.PostId)
}
