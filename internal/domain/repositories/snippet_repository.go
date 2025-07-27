package repositories

import "github.com/avanboxel/snippy/internal/domain/models"

type SnippetRepository interface {
	SaveSnippet(snippet *models.Snippet) error
	GetSnippet(id int) (*models.Snippet, error)
	ListSnippets() ([]models.Snippet, error)
	DeleteSnippet(id int) error
	SearchSnippets(code string, lang string, tags []string) ([]models.Snippet, error)
	Close() error
}
