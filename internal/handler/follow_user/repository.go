package follow_user

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type FollowRepository struct {
	dataRepository *database.Database
}

func NewFollowRepository(dataRepository *database.Database) *FollowRepository {
	return &FollowRepository{
		dataRepository: dataRepository,
	}
}

func (r *FollowRepository) AddFollow(data *model.Follow) error {
	return r.dataRepository.Client.AddFollow(data)
}
