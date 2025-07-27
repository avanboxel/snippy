package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLite() (*SQLiteDB, error) {
	dbPath := "snippy.sql"
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping SQLite: %w", err)
	}

	sqlite := &SQLiteDB{db: db}
	if err := sqlite.createTables(); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return sqlite, nil
}

func (s *SQLiteDB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS snippets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT NOT NULL,
		language TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS snippet_tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		snippet_id INTEGER NOT NULL,
		tag TEXT NOT NULL,
		FOREIGN KEY (snippet_id) REFERENCES snippets(id) ON DELETE CASCADE
	);

	CREATE INDEX IF NOT EXISTS idx_snippet_tag ON snippet_tags(snippet_id, tag);
	CREATE INDEX IF NOT EXISTS idx_snippet_created ON snippets(created_at);

	CREATE TRIGGER IF NOT EXISTS update_snippet_timestamp 
	AFTER UPDATE ON snippets
	BEGIN
		UPDATE snippets SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
	END;
	`

	_, err := s.db.Exec(query)
	return err
}
