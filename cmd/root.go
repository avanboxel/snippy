package cmd

import (
	"os"

	"github.com/avanboxel/snippy/cmd/add"
	"github.com/avanboxel/snippy/cmd/clean"
	"github.com/avanboxel/snippy/cmd/list"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snippy",
	Short: "A simple offline snippet manager",
	Long:  `A simple offline snippet manager`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(add.Init())
	rootCmd.AddCommand(list.Init())
	rootCmd.AddCommand(clean.Init())
}
