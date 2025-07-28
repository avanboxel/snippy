package clean

import (
	"log"
	"strings"

	"github.com/avanboxel/snippy/internal/application/commands"
	"github.com/avanboxel/snippy/internal/infrastructure/db"
	"github.com/spf13/cobra"
)

var lang string
var tags string
var search string
var id int

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans all code snippet saved with Snippy",
	Args:  cobra.MaximumNArgs(0),
	Long:  "Cleans all code snippet saved with Snippy",
	Run:   run,
}

func Init() *cobra.Command {
	c := cleanCmd
	c.PersistentFlags().StringVarP(&lang, "lang", "l", "", "Language (optional)")
	c.PersistentFlags().StringVarP(&tags, "tags", "t", "", "Tags (comma-separated, optional)")
	c.PersistentFlags().StringVarP(&search, "search", "s", "", "Search by part of snippet (optional)")
	c.PersistentFlags().IntVarP(&id, "id", "i", 0, "Id (optional)")
	return c
}

func run(cmd *cobra.Command, args []string) {
	c := commands.CleanSnippetsCommand{}
	if tags != "" {
		c.Tags = strings.Split(tags, ",")
	}
	c.Code = search
	c.Lang = lang
	c.Id = id

	db, err := db.NewSQLite()
	if err != nil {
		log.Fatal("Unable to access db.")
	}

	commands.CleanSnippets(db, c)
}
