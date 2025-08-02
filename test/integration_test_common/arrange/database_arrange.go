package integration_test_arrange

import (
	"escalateservice/cmd/provider"
	database "escalateservice/internal/db"
)

func CreateTestDatabase() *database.Database {
	provider := provider.NewProvider("test", "postgres://postgres:artis@localhost:5432/artis?search_path=public&sslmode=disable")
	sqlDb, err := provider.ProvideDb()
	if err != nil {
		panic(err)
	}
	return database.NewDatabase(sqlDb)
}
