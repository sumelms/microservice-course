package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath string
	folderPath string
	numSteps   int

	rootCmd = &cobra.Command{
		Use:   "migration",
		Short: "A migrations manager",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "./config/config.yml", "config file")
	rootCmd.PersistentFlags().StringVar(&folderPath, "folder", "./db/migrations", "migrations folder")
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(createCmd)
}
