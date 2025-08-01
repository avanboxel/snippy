package list

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/avanboxel/snippy/internal/application/queries"
	"github.com/avanboxel/snippy/internal/infrastructure/db"
	"github.com/spf13/cobra"
)

var lang string
var tags string
var search string
var id int

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all code snippet saved with Snippy",
	Args:  cobra.MaximumNArgs(0),
	Long:  "Lists all code snippet saved with Snippy",
	Run:   run,
}

func Init() *cobra.Command {
	c := listCmd
	c.PersistentFlags().StringVarP(&lang, "lang", "l", "", "Language (optional)")
	c.PersistentFlags().StringVarP(&tags, "tags", "t", "", "Tags (comma-separated, optional)")
	c.PersistentFlags().StringVarP(&search, "search", "s", "", "Search by part of snippet (optional)")
	c.PersistentFlags().IntVarP(&id, "id", "i", 0, "Id (optional)")
	return c
}

func run(cmd *cobra.Command, args []string) {
	q := queries.GetSnippetsQuery{}
	if tags != "" {
		q.Tags = strings.Split(tags, ",")
	}
	q.Code = search
	q.Lang = lang
	q.Id = id

	db, err := db.NewSQLite()
	if err != nil {
		log.Fatal("Unable to access db.")
	}

	s := queries.GetSnippets(db, q)
	if len(s) == 0 {
		fmt.Println("No snippets found.")
		return
	}

	for _, v := range s {
		fmt.Println(strings.Repeat("=", 20) + "<snippet>" + strings.Repeat("=", 20))
		fmt.Println("Id:\n\t" + strconv.Itoa(v.Id))
		fmt.Println("Language:\n\t" + v.Language)
		fmt.Println("Tags:\n\t" + strings.Join(v.Tags, ", "))
		fmt.Println("Code:")
		scanner := bufio.NewScanner(strings.NewReader(v.Code))
		for scanner.Scan() {
			fmt.Println("\t" + scanner.Text())
		}
		fmt.Println(strings.Repeat("=", 20) + "</snippet>" + strings.Repeat("=", 19))
	}
}
