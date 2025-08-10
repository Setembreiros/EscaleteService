package db_integration_test

import (
	"testing"

	database "escalateservice/internal/db"
	model "escalateservice/internal/model/domain"
	integration_test_action "escalateservice/test/integration_test_common/action"
	integration_test_arrange "escalateservice/test/integration_test_common/arrange"
	integration_test_assert "escalateservice/test/integration_test_common/assert"
)

var db *database.Database

func setUp(t *testing.T) {
	db = integration_test_arrange.CreateTestDatabase()
}

func tearDown() {
	db.Client.Clean()
}

func TestUpdatePostScoresProcedure_review0(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post := &model.Post{PostId: "post1", Username: user1.Username}
	integration_test_arrange.AddPost(t, post)
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})

	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 1,
		PostId:   post.PostId,
		Reviewer: user2.Username,
		Rating:   0,
	})

	integration_test_action.CallProcedureUpdatePostScores(t, db)

	// Score esperado:
	// Likes: 2 * 1 = 2
	// Superlikes: 1 * 10 = 10
	// Review rating 0: -200
	// Total esperado = 2 + 10 - 200 = -188	expectedScore := -188
	expectedScore := -188
	integration_test_assert.AssertPostScore(t, db, post.PostId, expectedScore)
}

func TestUpdatePostScoresProcedure_review1(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post := &model.Post{PostId: "post1", Username: user1.Username}
	integration_test_arrange.AddPost(t, post)
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})

	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 1,
		PostId:   post.PostId,
		Reviewer: user2.Username,
		Rating:   1,
	})

	integration_test_action.CallProcedureUpdatePostScores(t, db)

	// Score esperado:
	// Likes: 2 * 1 = 2
	// Superlikes: 1 * 10 = 10
	// Review rating 1: -100
	// Total esperado = 2 + 10 - 100 = -88	expectedScore := -88
	expectedScore := -88
	integration_test_assert.AssertPostScore(t, db, post.PostId, expectedScore)
}

func TestUpdatePostScoresProcedure_review2(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post := &model.Post{PostId: "post1", Username: user1.Username}
	integration_test_arrange.AddPost(t, post)
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})

	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 1,
		PostId:   post.PostId,
		Reviewer: user2.Username,
		Rating:   2,
	})

	integration_test_action.CallProcedureUpdatePostScores(t, db)

	// Score esperado:
	// Likes: 2 * 1 = 2
	// Superlikes: 1 * 10 = 10
	// Review rating 2: 0
	// Total esperado = 2 + 10 + 0 = 12	expectedScore := 12
	expectedScore := 12
	integration_test_assert.AssertPostScore(t, db, post.PostId, expectedScore)
}

func TestUpdatePostScoresProcedure_review3(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post := &model.Post{PostId: "post1", Username: user1.Username}
	integration_test_arrange.AddPost(t, post)
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})

	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 1,
		PostId:   post.PostId,
		Reviewer: user2.Username,
		Rating:   3,
	})

	integration_test_action.CallProcedureUpdatePostScores(t, db)

	// Score esperado:
	// Likes: 2 * 1 = 2
	// Superlikes: 1 * 10 = 10
	// Review rating 3: 100
	// Total esperado = 2 + 10 + 100 = 112	expectedScore := 112
	expectedScore := 112
	integration_test_assert.AssertPostScore(t, db, post.PostId, expectedScore)
}

func TestUpdatePostScoresProcedure_review4(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post := &model.Post{PostId: "post1", Username: user1.Username}
	integration_test_arrange.AddPost(t, post)
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})

	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 1,
		PostId:   post.PostId,
		Reviewer: user2.Username,
		Rating:   4,
	})

	integration_test_action.CallProcedureUpdatePostScores(t, db)

	// Score esperado:
	// Likes: 2 * 1 = 2
	// Superlikes: 1 * 10 = 10
	// Review rating 4: 200
	// Total esperado = 2 + 10 + 200 = 212	expectedScore := 212
	expectedScore := 212
	integration_test_assert.AssertPostScore(t, db, post.PostId, expectedScore)
}

func TestUpdatePostScoresProcedure_review5(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post := &model.Post{PostId: "post1", Username: user1.Username}
	integration_test_arrange.AddPost(t, post)
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post.PostId,
		Username: user2.Username,
	})

	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 1,
		PostId:   post.PostId,
		Reviewer: user2.Username,
		Rating:   5,
	})

	integration_test_action.CallProcedureUpdatePostScores(t, db)

	// Score esperado:
	// Likes: 2 * 1 = 2
	// Superlikes: 1 * 10 = 10
	// Review rating 5: 300
	// Total esperado = 2 + 10 + 300 = 312	expectedScore := 312
	expectedScore := 312
	integration_test_assert.AssertPostScore(t, db, post.PostId, expectedScore)
}

func TestUpdatePostScoresProcedure_MultiplePost(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post1 := &model.Post{PostId: "post1", Username: user1.Username}
	integration_test_arrange.AddPost(t, post1)
	post2 := &model.Post{PostId: "post2", Username: user2.Username}
	integration_test_arrange.AddPost(t, post2)
	post3 := &model.Post{PostId: "post3", Username: user1.Username}
	integration_test_arrange.AddPost(t, post3)
	// Post 1 -> likes: 2, superlikes: 1, review rating: 5
	// Post 2 -> likes: 0, superlikes: 2, review rating: 1, 2
	// Post 3 -> likes: 1, superlikes: 0, review rating: 5, 1
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post1.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post1.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddLikePost(t, &model.LikePost{
		PostId:   post3.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post1.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post2.PostId,
		Username: user1.Username,
	})
	integration_test_arrange.AddSuperlikePost(t, &model.SuperlikePost{
		PostId:   post2.PostId,
		Username: user2.Username,
	})
	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 1,
		PostId:   post1.PostId,
		Reviewer: user2.Username,
		Rating:   5,
	})
	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 2,
		PostId:   post2.PostId,
		Reviewer: user1.Username,
		Rating:   1,
	})
	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 3,
		PostId:   post2.PostId,
		Reviewer: user2.Username,
		Rating:   2,
	})
	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 4,
		PostId:   post3.PostId,
		Reviewer: user1.Username,
		Rating:   5,
	})
	integration_test_arrange.AddReview(t, &model.Review{
		ReviewId: 5,
		PostId:   post3.PostId,
		Reviewer: user2.Username,
		Rating:   1,
	})

	integration_test_action.CallProcedureUpdatePostScores(t, db)

	// Score esperado post1:
	// Likes: 2 * 1 = 2
	// Superlikes: 1 * 10 = 10
	// Review rating 5: 300
	// Total esperado = 2 + 10 + 300 = 312	expectedScore := 312
	expectedScore := 312
	integration_test_assert.AssertPostScore(t, db, post1.PostId, expectedScore)
	// Score esperado post2:
	// Likes: 0 * 1 = 0
	// Superlikes: 2 * 10 = 20
	// Review rating 1, 2: -100 + 0 = -100
	// Total esperado = 2 + 10 - 100 = -80	expectedScore := -80
	expectedScore = -80
	integration_test_assert.AssertPostScore(t, db, post2.PostId, expectedScore)
	// Score esperado post3:
	// Likes: 1 * 1 = 1
	// Superlikes: 0 * 10 = 0
	// Review rating 5, 1: 300 - 100: 200
	// Total esperado = 1 + 0 + 200 = 201
	expectedScore = 201
	integration_test_assert.AssertPostScore(t, db, post3.PostId, expectedScore)
}
