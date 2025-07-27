package add

import (
	"io"
	"log"
	"strings"

	"github.com/avanboxel/snippy/internal/application/commands"
	"github.com/avanboxel/snippy/internal/infrastructure/db"
	"github.com/spf13/cobra"
)

var lang string
var tags string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a code snippet to Snippy",
	Args:  cobra.MaximumNArgs(1),
	Long:  `All software has versions. This is Snippy's`,
	Run:   run,
}

func Init() *cobra.Command {
	c := addCmd
	c.PersistentFlags().StringVarP(&lang, "lang", "l", "go", "Language (optional)")
	c.PersistentFlags().StringVarP(&tags, "tags", "t", "", "Tags (comma-separated, optional)")
	return c
}

func run(cmd *cobra.Command, args []string) {
	var inputReader io.Reader = cmd.InOrStdin()
	var code string

	if len(args) > 0 {
		code = args[0]
	} else {
		code = processInput(inputReader)
	}

	db, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err.Error())
	}

	commands.CreateSnippet(
		db,
		commands.CreateSnippetCommand{
			Code: code,
			Tags: strings.Split(tags, ","),
			Lang: lang,
		},
	)
}

func processInput(r io.Reader) string {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		log.Fatal("Unable to read input file")
	}
	return buf.String()
}
