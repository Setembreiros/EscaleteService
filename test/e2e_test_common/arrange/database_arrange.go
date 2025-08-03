package e2e_test_arrange

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

	err := database.Client.AddUser(user)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertUserExists(t, database, user.Username, user)
}

func AddPost(t *testing.T, post *model.Post) {
	database := CreateTestDatabase()

	err := database.Client.AddPost(post)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertPostExists(t, database, post.PostId, post)
}
