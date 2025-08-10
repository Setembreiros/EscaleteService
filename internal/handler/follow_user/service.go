package follow_user

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	AddFollow(data *model.Follow) error
}

type FollowService struct {
	repository Repository
}

func NewFollowService(repository Repository) *FollowService {
	return &FollowService{
		repository: repository,
	}
}

func (s *FollowService) AddFollow(data *model.Follow) {
	err := s.repository.AddFollow(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error adding follow, %s -> %s", data.Follower, data.Followee)
		return
	}

	log.Info().Msgf("Follow was added, %s -> %s", data.Follower, data.Followee)
}
