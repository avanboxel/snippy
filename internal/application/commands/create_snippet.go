package commands

import (
	"github.com/avanboxel/snippy/internal/domain/models"
	"github.com/avanboxel/snippy/internal/domain/repositories"
)

type CreateSnippetCommand struct {
	Code string
	Tags []string
	Lang string
}

func CreateSnippet(r repositories.SnippetRepository, c CreateSnippetCommand) {
	s := models.NewSnippet(c.Code, c.Tags, c.Lang)
	r.SaveSnippet(s)
}
