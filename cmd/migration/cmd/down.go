package cmd

import (
	"fmt"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"

	// migration source adapter
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// database driver
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"github.com/sumelms/microservice-course/pkg/config"
	"github.com/sumelms/microservice-course/pkg/database/postgres"
)

var (
	downCmd = &cobra.Command{
		Use:   "down",
		Short: "migrations down",
		Run: func(cmd *cobra.Command, args []string) {

			cfg, err := config.NewConfig(configPath)
			if err != nil {
				log.Fatalf("error loading the config file: %s", err.Error())
			}
			db, err := postgres.Connect(cfg.Database)
			if err != nil {
				log.Fatalf("error connecting to the database: %s", err.Error())
			}
			driver, err := migratePostgres.WithInstance(db.DB, &migratePostgres.Config{})
			if err != nil {
				log.Fatalf("WithInstance err %s", err.Error())
			}
			m, err := migrate.NewWithDatabaseInstance(
				fmt.Sprintf("file://%s", folderPath),
				"postgres", driver)
			if err != nil {
				log.Fatalf("NewWithDatabaseInstance err %s", err.Error())
			}
			numSteps *= -1
			if numSteps < 0 {
				if err := m.Steps(numSteps); err != nil {
					log.Fatalf("m.Steps(%d) error: %s", numSteps, err.Error())
				}
				return
			}
			if err := m.Down(); err != nil {
				log.Fatalf("m.Down() error: %s", err.Error())
			}
		},
	}
)

func init() { //nolint: gochecknoinits
	downCmd.Flags().IntVar(&numSteps, "steps", 0, "num of migrations to down")
}
