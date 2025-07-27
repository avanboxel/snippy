package db

import (
	"fmt"
	"strings"

	"github.com/avanboxel/snippy/internal/domain/models"
)

func (s *SQLiteDB) SaveSnippet(snippet *models.Snippet) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		"INSERT INTO snippets (code, language) VALUES (?, ?)",
		snippet.Code, snippet.Language,
	)
	if err != nil {
		return fmt.Errorf("failed to insert snippet: %w", err)
	}

	snippetID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	snippet.Id = int(snippetID)

	for _, tag := range snippet.Tags {
		_, err := tx.Exec(
			"INSERT INTO snippet_tags (snippet_id, tag) VALUES (?, ?)",
			snippetID, tag,
		)
		if err != nil {
			return fmt.Errorf("failed to insert tag: %w", err)
		}
	}

	return tx.Commit()
}

func (s *SQLiteDB) GetSnippet(id int) (*models.Snippet, error) {
	snippet := &models.Snippet{Id: id}

	err := s.db.QueryRow(
		"SELECT code, language FROM snippets WHERE id = ?", id,
	).Scan(&snippet.Code, &snippet.Language)
	if err != nil {
		return nil, fmt.Errorf("failed to get snippet: %w", err)
	}

	tags, err := s.getSnippetTags(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get snippet tags: %w", err)
	}
	snippet.Tags = tags

	return snippet, nil
}

func (s *SQLiteDB) getSnippetTags(snippetID int) ([]string, error) {
	rows, err := s.db.Query(
		"SELECT tag FROM snippet_tags WHERE snippet_id = ?", snippetID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, rows.Err()
}

func (s *SQLiteDB) ListSnippets() ([]models.Snippet, error) {
	rows, err := s.db.Query("SELECT id, code, language FROM snippets ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("failed to list snippets: %w", err)
	}
	defer rows.Close()

	var snippetList []models.Snippet
	for rows.Next() {
		snippet := models.Snippet{}
		if err := rows.Scan(&snippet.Id, &snippet.Code, &snippet.Language); err != nil {
			return nil, fmt.Errorf("failed to scan snippet: %w", err)
		}

		tags, err := s.getSnippetTags(snippet.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to get tags for snippet %d: %w", snippet.Id, err)
		}
		snippet.Tags = tags

		snippetList = append(snippetList, snippet)
	}

	return snippetList, rows.Err()
}

func (s *SQLiteDB) DeleteSnippet(id int) error {
	result, err := s.db.Exec("DELETE FROM snippets WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete snippet: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("snippet with id %d not found", id)
	}

	return nil
}

func (s *SQLiteDB) SearchSnippets(
	code string,
	lang string,
	tags []string,
) ([]models.Snippet, error) {
	args := []any{}
	whereString := ""
	if code != "" {
		args = append(args, "%"+code+"%")
		whereString = whereString + " code LIKE ? "
	}
	if lang != "" {
		args = append(args, "%"+lang+"%")
		whereString = whereString + " lang LIKE ? "
	}
	if len(tags) > 0 {
		args = append(args, strings.Join(tags, ","))
		whereString = whereString + " st.tag IN (?) "
	}

	rows, err := s.db.Query(`
		SELECT DISTINCT s.id, s.code, s.language 
		FROM snippets s 
		LEFT JOIN snippet_tags st ON s.id = st.snippet_id 
		WHERE (`+whereString+`)
		ORDER BY s.created_at DESC
	`, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search snippets: %w", err)
	}
	defer rows.Close()

	var snippetList []models.Snippet
	for rows.Next() {
		snippet := models.Snippet{}
		if err := rows.Scan(&snippet.Id, &snippet.Code, &snippet.Language); err != nil {
			return nil, fmt.Errorf("failed to scan snippet: %w", err)
		}

		tags, err := s.getSnippetTags(snippet.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to get tags for snippet %d: %w", snippet.Id, err)
		}
		snippet.Tags = tags

		snippetList = append(snippetList, snippet)
	}

	return snippetList, rows.Err()
}

func (s *SQLiteDB) Close() error {
	return s.db.Close()
}
