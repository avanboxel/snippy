package commands

import (
	"github.com/avanboxel/snippy/internal/application/queries"
	"github.com/avanboxel/snippy/internal/domain/repositories"
)

type CleanSnippetsCommand struct {
	Code string
	Lang string
	Tags []string
	Id   int
}

func CleanSnippets(r repositories.SnippetRepository, c CleanSnippetsCommand) {
	if c.Id != 0 {
		r.DeleteSnippet(c.Id)
		return
	}
	s := queries.GetSnippets(r, queries.GetSnippetsQuery{
		Code: c.Code,
		Lang: c.Lang,
		Tags: c.Tags,
	})
	for _, v := range s {
		r.DeleteSnippet(v.Id)
	}
}
