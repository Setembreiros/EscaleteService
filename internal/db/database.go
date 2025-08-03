package database

import (
	model "escalateservice/internal/model/domain"

	_ "github.com/lib/pq"
)

//go:generate mockgen -source=database.go -destination=test/mock/database.go

type Database struct {
	Client DatabaseClient
}

type DatabaseClient interface {
	Clean()
	AddUser(user *model.User) error
	GetUser(username string) (*model.User, error)
	AddPost(post *model.Post) error
	GetPost(postId string) (*model.Post, error)
	AddReview(review *model.Review) error
	GetReview(reviewId uint64) (*model.Review, error)
	AddLikePost(likePost *model.LikePost) error
	GetLikePost(username, postId string) (*model.LikePost, error)
	RemoveLikePost(likePost *model.LikePost) error
	AddSuperlikePost(likePost *model.SuperlikePost) error
	GetSuperlikePost(username, postId string) (*model.SuperlikePost, error)
}

func NewDatabase(client DatabaseClient) *Database {
	return &Database{
		Client: client,
	}
}
