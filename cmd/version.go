package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Snippy",
	Long:  `All software has versions. This is Snippy's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Snippy - Simple Snippet Manager v0.0.4")
	},
}
