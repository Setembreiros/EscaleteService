package integration_test_arrange

import (
	"escalateservice/cmd/startup"
	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
	"testing"
)

func CreateTestDatabase() *database.Database {
	provider := startup.NewProvider("test", "postgres://postgres:artis@localhost:5432/artis?search_path=public&sslmode=disable")
	sqlDb, err := provider.ProvideDb()
	if err != nil {
		panic(err)
	}
	return database.NewDatabase(sqlDb)
}

func AddUser(t *testing.T, user *model.User) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.AddUser(user)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertUserExists(t, database, user.Username, user)
}

func AddUserBatch(t *testing.T, users []*model.User) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.BatchAddUsers(users)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		integration_test_assert.AssertUserExists(t, database, user.Username, user)
	}
}

func AddPost(t *testing.T, post *model.Post) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.AddPost(post)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertPostExists(t, database, post.PostId, post)
}

func AddPostBatch(t *testing.T, posts []*model.Post) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.BatchAddPosts(posts)
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		integration_test_assert.AssertPostExists(t, database, post.PostId, post)
	}
}

func AddLikePost(t *testing.T, likePost *model.LikePost) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.AddLikePost(likePost)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertLikePostExists(t, database, likePost)
}

func AddLikePostBatch(t *testing.T, likePosts []*model.LikePost) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.BatchAddLikePosts(likePosts)
	if err != nil {
		panic(err)
	}

	for _, likePost := range likePosts {
		integration_test_assert.AssertLikePostExists(t, database, likePost)
	}
}

func AddSuperlikePost(t *testing.T, superlikePost *model.SuperlikePost) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.AddSuperlikePost(superlikePost)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertSuperlikePostExists(t, database, superlikePost)
}

func AddSuperlikePostBatch(t *testing.T, superlikePosts []*model.SuperlikePost) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.BatchAddSuperlikePosts(superlikePosts)
	if err != nil {
		panic(err)
	}

	for _, superlikePost := range superlikePosts {
		integration_test_assert.AssertSuperlikePostExists(t, database, superlikePost)
	}
}

func AddReview(t *testing.T, review *model.Review) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.AddReview(review)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertReviewExists(t, database, review.ReviewId, review)
}

func AddReviewBatch(t *testing.T, reviews []*model.Review) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.BatchAddReviews(reviews)
	if err != nil {
		panic(err)
	}

	for _, review := range reviews {
		integration_test_assert.AssertReviewExists(t, database, review.ReviewId, review)
	}
}

func AddFollow(t *testing.T, follow *model.Follow) {
	database := CreateTestDatabase()
	defer database.Client.Close()

	err := database.Client.AddFollow(follow)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertFollowExists(t, database, follow)
}
