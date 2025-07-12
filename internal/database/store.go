package database

import (
	"database/sql"
	"time"
)

// Store encapsulates the database connection.
type Store struct {
	DB *sql.DB
}

// NewStore creates a new Store with a database connection.
func NewStore(db *sql.DB) *Store {
	return &Store{DB: db}
}

// GetMessages retrieves all messages from the database.
func (s *Store) GetMessages() ([]Message, error) {
	rows, err := s.DB.Query("SELECT id, content, created_at FROM messages ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// CreateMessage adds a new message to the database.
func (s *Store) CreateMessage(content string) (*Message, error) {
	now := time.Now().UTC()
	result, err := s.DB.Exec("INSERT INTO messages (content, created_at) VALUES (?, ?)", content, now)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:        id,
		Content:   content,
		CreatedAt: now,
	}, nil
}
