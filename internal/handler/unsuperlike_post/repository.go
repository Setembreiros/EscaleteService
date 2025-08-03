package unsuperlike_post

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type UnsuperlikePostRepository struct {
	dataRepository *database.Database
}

func NewUnsuperlikePostRepository(dataRepository *database.Database) *UnsuperlikePostRepository {
	return &UnsuperlikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *UnsuperlikePostRepository) RemoveSuperlikePost(data *model.SuperlikePost) error {
	return r.dataRepository.Client.RemoveSuperlikePost(data)
}
