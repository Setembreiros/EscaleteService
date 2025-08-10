package unfollow_user

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	RemoveFollow(data *model.Follow) error
}

type UnfollowService struct {
	repository Repository
}

func NewUnfollowService(repository Repository) *UnfollowService {
	return &UnfollowService{
		repository: repository,
	}
}

func (s *UnfollowService) RemoveFollow(data *model.Follow) {
	err := s.repository.RemoveFollow(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error removing follow, %s -> %s", data.Follower, data.Followee)
		return
	}

	log.Info().Msgf("Follow was removed, %s -> %s", data.Follower, data.Followee)
}
