package unlike_post

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type UnlikePostRepository struct {
	dataRepository *database.Database
}

func NewUnlikePostRepository(dataRepository *database.Database) *UnlikePostRepository {
	return &UnlikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *UnlikePostRepository) RemoveLikePost(data *model.LikePost) error {
	return r.dataRepository.Client.RemoveLikePost(data)
}
