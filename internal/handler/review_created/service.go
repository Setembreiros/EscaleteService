package review_created

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	AddReview(data *model.Review) error
}

type ReviewCreatedService struct {
	repository Repository
}

func NewReviewCreatedService(repository Repository) *ReviewCreatedService {
	return &ReviewCreatedService{
		repository: repository,
	}
}

func (s *ReviewCreatedService) AddReview(data *model.Review) {
	err := s.repository.AddReview(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error adding review %d", data.ReviewId)
		return
	}

	log.Info().Msgf("Review %d was added", data.ReviewId)
}
