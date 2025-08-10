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

func TestUpdateUserScoresProcedure_OneUser(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	post := &model.Post{PostId: "post1", Username: user1.Username, Score: 225}
	integration_test_arrange.AddPost(t, post)
	follow := &model.Follow{Follower: user2.Username, Followee: user1.Username}
	integration_test_arrange.AddFollow(t, follow)

	integration_test_action.CallProcedureUpdateUserScores(t, db)

	// Score esperado:
	// BaseScore: 800 = 800
	// PostScore: 225 = 225
	// Follows: 1 * 10 = 10
	// Total esperado = 800 + 10 + 225 = 1035 expectedScore := 1035
	expectedScore := 1035.0
	integration_test_assert.AssertUserScore(t, db, user1.Username, expectedScore)
}

func TestUpdateUserScoresProcedure_MultipleUser(t *testing.T) {
	setUp(t)
	defer tearDown()
	user1 := &model.User{Username: "user1"}
	integration_test_arrange.AddUser(t, user1)
	user2 := &model.User{Username: "user2"}
	integration_test_arrange.AddUser(t, user2)
	user3 := &model.User{Username: "user3"}
	integration_test_arrange.AddUser(t, user3)
	// User 1 -> follows: 1, post scores: 105, 18
	// User 2 -> follows: 0, post scores: 101
	// User 3 -> follows: 2, post scores: 25, 155, 666, 777
	post1 := &model.Post{PostId: "post1", Username: user1.Username, Score: 105}
	integration_test_arrange.AddPost(t, post1)
	post2 := &model.Post{PostId: "post2", Username: user3.Username, Score: 25}
	integration_test_arrange.AddPost(t, post2)
	post3 := &model.Post{PostId: "post3", Username: user3.Username, Score: 155}
	integration_test_arrange.AddPost(t, post3)
	post4 := &model.Post{PostId: "post4", Username: user1.Username, Score: 18}
	integration_test_arrange.AddPost(t, post4)
	post5 := &model.Post{PostId: "post5", Username: user2.Username, Score: 101}
	integration_test_arrange.AddPost(t, post5)
	post6 := &model.Post{PostId: "post6", Username: user3.Username, Score: 666}
	integration_test_arrange.AddPost(t, post6)
	post7 := &model.Post{PostId: "post7", Username: user3.Username, Score: 777}
	integration_test_arrange.AddPost(t, post7)
	follow1 := &model.Follow{Follower: user2.Username, Followee: user1.Username}
	integration_test_arrange.AddFollow(t, follow1)
	follow2 := &model.Follow{Follower: user2.Username, Followee: user3.Username}
	integration_test_arrange.AddFollow(t, follow2)
	follow3 := &model.Follow{Follower: user1.Username, Followee: user3.Username}
	integration_test_arrange.AddFollow(t, follow3)

	integration_test_action.CallProcedureUpdateUserScores(t, db)

	// Score esperado:
	// BaseScore: 800 = 800
	// PostScore: 105 + 18 / 2 = 123 /2 = 61.5
	// Follows: 1 * 10 = 10
	// Total esperado = 800 + 10 + 61.5 = 871.5 expectedScore := 871.5
	expectedScore := 871.5
	integration_test_assert.AssertUserScore(t, db, user1.Username, expectedScore)
	// Score esperado:
	// BaseScore: 800 = 800
	// PostScore: 101 = 101
	// Follows: 0 * 10 = 0
	// Total esperado = 800 + 0 + 101 = 901 expectedScore := 901
	expectedScore = 901.0
	integration_test_assert.AssertUserScore(t, db, user2.Username, expectedScore)
	// Score esperado:
	// BaseScore: 800 = 800
	// PostScore: 25 + 155 + 666 + 777 / 4 = 1623 /4 = 405.75
	// Follows: 2 * 10 = 20
	// Total esperado = 800 + 20 + 405.75 = 1225.75 expectedScore := 1225.75
	expectedScore = 1225.75
	integration_test_assert.AssertUserScore(t, db, user3.Username, expectedScore)
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
