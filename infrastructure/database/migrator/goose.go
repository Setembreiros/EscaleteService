package migrator

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq" // Cambia ao driver que uses
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type GooseClient struct {
	db      *sql.DB
	workdir string
	connStr string
}

func NewGooseClient(connStr string) (*GooseClient, error) {
	if connStr == "" {
		log.Error().Stack().Msg("No connection string provided")
		return nil, errors.New("no connection string provided")
	}

	// Abre DB
	db, err := sql.Open("postgres", strings.Trim(connStr, "\"")) // Asegúrate de cambiar "postgres" se usas outro motor
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to open database connection")
		return nil, err
	}

	// Verifica conexión
	if err := db.Ping(); err != nil {
		log.Error().Stack().Err(err).Msg("failed to ping database")
		return nil, err
	}

	// Directorio de migracións goose
	rootDir, err := findProjectRoot("EscalateService")
	if err != nil {
		log.Error().Err(err).Msg("Failed to locate EscalateService root")
		return nil, err
	}

	migrationPath := filepath.Join(rootDir, "infrastructure", "database", "migrator", "migrations")
	if fi, err := os.Stat(migrationPath); err != nil || !fi.IsDir() {
		cwd, _ := os.Getwd()
		log.Error().Stack().Msgf("migrations dir not found: %s, current directory: %s", migrationPath, cwd)
		return nil, fmt.Errorf("migrations dir not found: %s", migrationPath)
	}

	// Configura goose
	goose.SetBaseFS(os.DirFS(migrationPath))

	return &GooseClient{
		db:      db,
		workdir: migrationPath,
		connStr: connStr,
	}, nil
}

func (gc *GooseClient) ApplyMigrations(ctx context.Context) error {
	defer func() {
		if err := gc.db.Close(); err != nil {
			log.Error().Stack().Err(err).Msg("error closing database")
		}
	}()

	log.Info().Msg("Applying goose migrations...")

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(gc.db, "migrations"); err != nil {
		log.Error().Stack().Err(err).Msg("goose migration failed")
		return err
	}

	log.Info().Msg("Migrations applied successfully")
	return nil
}

func findProjectRoot(targetDirName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	for {
		base := filepath.Base(cwd)
		if base == targetDirName {
			return cwd, nil
		}

		parent := filepath.Dir(cwd)
		if parent == cwd {
			break // chegamos ao root do sistema de ficheiros
		}
		cwd = parent
	}

	return "", fmt.Errorf("directory %s not found in parent paths", targetDirName)
}
