package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	// migration source adapter
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// database driver
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var (
	name string

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "migrations create",
		Run: func(cmd *cobra.Command, args []string) {
			baseName := fmt.Sprintf("%s/%d_%s", folderPath, time.Now().Unix(), name)
			upName := fmt.Sprintf("%s.up.sql", baseName)
			downName := fmt.Sprintf("%s.down.sql", baseName)
			createFile(upName)
			createFile(downName)
		},
	}
)

func createFile(fname string) {
	if _, err := os.Create(fname); err != nil {
		log.Fatalf("error creating migration file %s: %s", fname, err)
	}
}

func init() { //nolint: gochecknoinits
	createCmd.Flags().StringVar(&name, "name", "", "name of the migration")
	createCmd.MarkFlagRequired("name") //nolint: errcheck
}
