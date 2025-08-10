package integration_test_action

import (
	"escalateservice/cmd/startup"
	database "escalateservice/internal/db"
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

func CallProcedureUpdatePostScores(t *testing.T, db *database.Database) {
	err := db.Client.CallProcedure("update_post_scores")
	if err != nil {
		panic(err)
	}
}
