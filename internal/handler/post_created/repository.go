package post_created

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type PostCreatedRepository struct {
	dataRepository *database.Database
}

func NewPostCreatedRepository(dataRepository *database.Database) *PostCreatedRepository {
	return &PostCreatedRepository{
		dataRepository: dataRepository,
	}
}

func (r *PostCreatedRepository) AddPost(data *model.Post) error {
	return r.dataRepository.Client.AddPost(data)
}
