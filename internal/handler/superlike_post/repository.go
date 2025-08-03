package superlike_post

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type SuperlikePostRepository struct {
	dataRepository *database.Database
}

func NewSuperlikePostRepository(dataRepository *database.Database) *SuperlikePostRepository {
	return &SuperlikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *SuperlikePostRepository) AddSuperlikePost(data *model.SuperlikePost) error {
	return r.dataRepository.Client.AddSuperlikePost(data)
}
