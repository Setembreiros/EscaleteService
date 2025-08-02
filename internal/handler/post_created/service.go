package post_created

import (
	model "escalateservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	AddPost(data *model.Post) error
}

type PostCreatedService struct {
	repository Repository
}

func NewPostCreatedService(repository Repository) *PostCreatedService {
	return &PostCreatedService{
		repository: repository,
	}
}

func (s *PostCreatedService) AddPost(data *model.Post) {
	err := s.repository.AddPost(data)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error adding post %s", data.PostId)
		return
	}

	log.Info().Msgf("Post %s was added", data.PostId)
}
