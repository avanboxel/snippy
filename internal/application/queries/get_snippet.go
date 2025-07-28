package queries

import (
	"github.com/avanboxel/snippy/internal/domain/models"
	"github.com/avanboxel/snippy/internal/domain/repositories"
)

type GetSnippetsQuery struct {
	Id   int
	Code string
	Lang string
	Tags []string
}

func GetSnippets(r repositories.SnippetRepository, q GetSnippetsQuery) []models.Snippet {
	if q.Id == 0 && q.Code == "" && len(q.Tags) == 0 && q.Lang == "" {
		l, _ := r.ListSnippets()
		return l
	}

	if q.Id != 0 {
		s, _ := r.GetSnippet(q.Id)
		if s == nil {
			return []models.Snippet{}
		}
		return []models.Snippet{*s}
	}

	l, _ := r.SearchSnippets(q.Code, q.Lang, q.Tags)
	return l
}
