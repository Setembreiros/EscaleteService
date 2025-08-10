package unfollow_user

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type UnfollowRepository struct {
	dataRepository *database.Database
}

func NewUnfollowRepository(dataRepository *database.Database) *UnfollowRepository {
	return &UnfollowRepository{
		dataRepository: dataRepository,
	}
}

func (r *UnfollowRepository) RemoveFollow(data *model.Follow) error {
	return r.dataRepository.Client.RemoveFollow(data)
}
