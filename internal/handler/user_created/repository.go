package user_created

import (
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
)

type UserCreatedRepository struct {
	dataRepository *database.Database
}

func NewUserCreatedRepository(dataRepository *database.Database) *UserCreatedRepository {
	return &UserCreatedRepository{
		dataRepository: dataRepository,
	}
}

func (r *UserCreatedRepository) AddUser(data *model.User) error {
	return r.dataRepository.Client.AddUser(data)
}
