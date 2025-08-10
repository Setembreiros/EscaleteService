package db_test

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
	integration_test_action "escalateservice/test/integration_test_common/action"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
)

var db *database.Database

func setUp(t *testing.T) {
	db = integration_test_arrange.CreateTestDatabase()
}

func tearDown() {
	db.Client.Clean()
}

func TestUpdatePostScoresProcedure_Load(t *testing.T) {
	setUp(t)
	defer tearDown()
	t.Run("Encher base de datos", func(t *testing.T) {
		startArrange := time.Now()
		users := CreateRandomUsersInBatches(t, 1000, 1000)
		posts := CreateRandomPostsInBatches(t, users, 3000, 3000)
		var wg sync.WaitGroup
		wg.Add(4)
		go func() {
			defer wg.Done()
			CreateRandomLikePostsInBatches(t, posts, users, 100000, 10000) // 100k likes
		}()
		go func() {
			defer wg.Done()
			CreateRandomSuperlikePostsInBatches(t, posts, users, 10000, 10000) // 10k superlikes
		}()
		go func() {
			defer wg.Done()
			CreateRandomReviewsInBatches(t, posts, users, 5000, 5000) // 5k reviews
		}()
		go func() {
			defer wg.Done()
			CreateRandomFollowsInBatches(t, users, 10000, 10000) // 5k reviews
		}()
		wg.Wait()
		elapsedArrange := time.Since(startArrange)
		elapsedArrangeSeconds := float64(elapsedArrange.Milliseconds()) / 1000.0
		t.Logf("Arrange completado en %fs", elapsedArrangeSeconds)
	})

	t.Run("ExecutarProcedemento update_post_scores", func(t *testing.T) {
		start := time.Now()

		// Executar varias veces para medir rendemento consistente
		for i := 0; i < 3; i++ {
			startAlone := time.Now()
			integration_test_action.CallProcedureUpdatePostScores(t, db)
			elapsedAlone := time.Since(startAlone)
			elapsedAloneSeconds := float64(elapsedAlone.Milliseconds()) / 1000.0
			t.Logf("Execución update_post_scores %d completada en %fs", i+1, elapsedAloneSeconds)
		}

		elapsed := time.Since(start)
		elapsedSeconds := float64(elapsed.Milliseconds()) / 1000.0
		avgTime := elapsedSeconds / 3
		t.Logf("update_post_scores tardou %f en executarse", avgTime)
	})
	t.Run("ExecutarProcedemento update_user_scores", func(t *testing.T) {
		start := time.Now()

		// Executar varias veces para medir rendemento consistente
		for i := 0; i < 3; i++ {
			startAlone := time.Now()
			integration_test_action.CallProcedureUpdateUserScores(t, db)
			elapsedAlone := time.Since(startAlone)
			elapsedAloneSeconds := float64(elapsedAlone.Milliseconds()) / 1000.0
			t.Logf("Execución update_user_scores %d completada en %fs", i+1, elapsedAloneSeconds)
		}

		elapsed := time.Since(start)
		elapsedSeconds := float64(elapsed.Milliseconds()) / 1000.0
		avgTime := elapsedSeconds / 3
		t.Logf("update_user_scores tardou %f en executarse", avgTime)
	})
}

func CreateRandomUsersInBatches(t *testing.T, total, batchSize int) []*model.User {
	users := make([]*model.User, 0, total)
	for i := 0; i < total; i += batchSize {
		currentBatch := batchSize
		if i+batchSize > total {
			currentBatch = total - i
		}

		batchUsers := make([]*model.User, currentBatch)
		for j := 0; j < currentBatch; j++ {
			batchUsers[j] = &model.User{
				Username: "user" + strconv.Itoa(i+j),
			}
		}

		integration_test_arrange.AddUserBatch(t, batchUsers)
		users = append(users, batchUsers...)
	}
	return users
}

func CreateRandomPostsInBatches(t *testing.T, users []*model.User, total, batchSize int) []*model.Post {
	posts := make([]*model.Post, 0, total)
	for i := 0; i < total; i += batchSize {
		currentBatch := batchSize
		if i+batchSize > total {
			currentBatch = total - i
		}
		batchPosts := make([]*model.Post, currentBatch)
		for j := 0; j < currentBatch; j++ {
			userId := rand.Intn(len(users))
			user := users[userId]
			batchPosts[j] = &model.Post{
				PostId:   "post" + strconv.Itoa(i+j),
				Username: user.Username,
			}
		}
		integration_test_arrange.AddPostBatch(t, batchPosts)
		posts = append(posts, batchPosts...)
	}
	return posts
}

func CreateRandomLikePostsInBatches(t *testing.T, posts []*model.Post, users []*model.User, total, batchSize int) {
	var likeCache = make(map[string]bool)
	for i := 0; i < total; i += batchSize {
		currentBatch := batchSize
		if i+batchSize > total {
			currentBatch = total - i
		}
		batchLikePost := make([]*model.LikePost, currentBatch)
		for j := 0; j < currentBatch; j++ {
			for {
				postId := rand.Intn(len(posts))
				userId := rand.Intn(len(users))
				post := posts[postId]
				user := users[userId]
				likeCacheKey := user.Username + "_" + post.PostId
				if _, exists := likeCache[likeCacheKey]; exists {
					continue // Skip if this like already exists
				}
				likeCache[likeCacheKey] = true
				batchLikePost[j] = &model.LikePost{
					PostId:   post.PostId,
					Username: user.Username,
				}
				break
			}
		}
		integration_test_arrange.AddLikePostBatch(t, batchLikePost)
	}
}

func CreateRandomSuperlikePostsInBatches(t *testing.T, posts []*model.Post, users []*model.User, total, batchSize int) {
	var superlikeCache = make(map[string]bool)
	for i := 0; i < total; i += batchSize {
		currentBatch := batchSize
		if i+batchSize > total {
			currentBatch = total - i
		}
		batchSuperlikePost := make([]*model.SuperlikePost, currentBatch)
		for j := 0; j < currentBatch; j++ {
			for {
				postId := rand.Intn(len(posts))
				userId := rand.Intn(len(users))
				post := posts[postId]
				user := users[userId]
				superlikeCacheKey := user.Username + "_" + post.PostId
				if _, exists := superlikeCache[superlikeCacheKey]; exists {
					continue // Skip if this superlike already exists
				}
				superlikeCache[superlikeCacheKey] = true
				batchSuperlikePost[j] = &model.SuperlikePost{
					PostId:   post.PostId,
					Username: user.Username,
				}
				break
			}
		}
		integration_test_arrange.AddSuperlikePostBatch(t, batchSuperlikePost)
	}
}

func CreateRandomReviewsInBatches(t *testing.T, posts []*model.Post, users []*model.User, total, batchSize int) {
	var reviewCache = make(map[string]bool)
	for i := 0; i < total; i += batchSize {
		currentBatch := batchSize
		if i+batchSize > total {
			currentBatch = total - i
		}
		batchReview := make([]*model.Review, currentBatch)
		for j := 0; j < currentBatch; j++ {
			for {
				postId := rand.Intn(len(posts))
				userId := rand.Intn(len(users))
				post := posts[postId]
				user := users[userId]
				reviewCacheKey := user.Username + "_" + post.PostId
				if _, exists := reviewCache[reviewCacheKey]; exists {
					continue // Skip if this review already exists
				}
				reviewCache[reviewCacheKey] = true
				batchReview[j] = &model.Review{
					ReviewId: uint64(i + j + 1),
					PostId:   post.PostId,
					Reviewer: user.Username,
					Rating:   rand.Intn(5) + 1, // Rating between 1 and 5
				}
				break // Exit the loop once a unique review is created
			}
		}
		integration_test_arrange.AddReviewBatch(t, batchReview)
	}
}

func CreateRandomFollowsInBatches(t *testing.T, users []*model.User, total, batchSize int) {
	var followCache = make(map[string]bool)
	for i := 0; i < total; i += batchSize {
		currentBatch := batchSize
		if i+batchSize > total {
			currentBatch = total - i
		}
		batchFollow := make([]*model.Follow, currentBatch)
		for j := 0; j < currentBatch; j++ {
			for {
				followerId := rand.Intn(len(users))
				followeeId := rand.Intn(len(users))
				if followerId == followeeId {
					continue
				}
				follower := users[followerId]
				followee := users[followeeId]
				followCacheKey := follower.Username + "_" + followee.Username
				if _, exists := followCache[followCacheKey]; exists {
					continue // Skip if this review already exists
				}
				followCache[followCacheKey] = true
				batchFollow[j] = &model.Follow{
					Follower: follower.Username,
					Followee: followee.Username, // Rating between 1 and 5
				}
				break // Exit the loop once a unique review is created
			}
		}
		integration_test_arrange.AddFollowBatch(t, batchFollow)
	}
}
