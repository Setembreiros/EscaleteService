package review_created

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type ReviewCreatedRepository struct {
	dataRepository *database.Database
}

func NewReviewCreatedRepository(dataRepository *database.Database) *ReviewCreatedRepository {
	return &ReviewCreatedRepository{
		dataRepository: dataRepository,
	}
}

func (r *ReviewCreatedRepository) AddReview(data *model.Review) error {
	return r.dataRepository.Client.AddReview(data)
}
