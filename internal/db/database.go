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
	Close()
	CallProcedure(name string) error
	AddUser(user *model.User) error
	BatchAddUsers(users []*model.User) error
	GetUser(username string) (*model.User, error)
	AddPost(post *model.Post) error
	BatchAddPosts(posts []*model.Post) error
	GetPost(postId string) (*model.Post, error)
	AddReview(review *model.Review) error
	BatchAddReviews(reviews []*model.Review) error
	GetReview(reviewId uint64) (*model.Review, error)
	AddLikePost(likePost *model.LikePost) error
	BatchAddLikePosts(likePosts []*model.LikePost) error
	GetLikePost(username, postId string) (*model.LikePost, error)
	RemoveLikePost(likePost *model.LikePost) error
	AddSuperlikePost(likePost *model.SuperlikePost) error
	BatchAddSuperlikePosts(superlikePosts []*model.SuperlikePost) error
	GetSuperlikePost(username, postId string) (*model.SuperlikePost, error)
	RemoveSuperlikePost(superlikePost *model.SuperlikePost) error
	AddFollow(follow *model.Follow) error
	BatchAddFollows(follow []*model.Follow) error
	GetFollow(follower, followee string) (*model.Follow, error)
	RemoveFollow(follow *model.Follow) error
}

func NewDatabase(client DatabaseClient) *Database {
	return &Database{
		Client: client,
	}
}
