package cmd

import (
	"fmt"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"

	// migration source adapter.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// database driver.
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/sumelms/microservice-course/pkg/config"
	"github.com/sumelms/microservice-course/pkg/database/postgres"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "migrations up",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.NewConfig(configPath)
		if err != nil {
			log.Fatal(err)
		}
		db, err := postgres.Connect(cfg.Database)
		if err != nil {
			log.Fatalf("error connecting to the database: %s", err.Error())
		}
		driver, err := migratePostgres.WithInstance(db.DB, &migratePostgres.Config{})
		if err != nil {
			log.Fatal(err)
		}
		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", folderPath),
			"postgres", driver)
		if err != nil {
			log.Fatal(err)
		}
		if numSteps > 0 {
			if err := m.Steps(numSteps); err != nil {
				fmt.Printf("error: %v", err)
			}

			return
		}
		m.Up() //nolint: errcheck
	},
}

func init() { //nolint: gochecknoinits
	upCmd.Flags().IntVar(&numSteps, "steps", 0, "num of migrations to up")
}
