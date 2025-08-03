package like_post

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type LikePostRepository struct {
	dataRepository *database.Database
}

func NewLikePostRepository(dataRepository *database.Database) *LikePostRepository {
	return &LikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *LikePostRepository) AddLikePost(data *model.LikePost) error {
	return r.dataRepository.Client.AddLikePost(data)
}
