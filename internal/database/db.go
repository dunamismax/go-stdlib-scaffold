package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect initializes the database connection.
func Connect() error {
	dbPath := os.Getenv("APP_DB_PATH")
	if dbPath == "" {
		dbPath = "app.db" // Default path
	}

	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return createMessagesTable()
}

func createMessagesTable() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			created_at DATETIME NOT NULL
		)
	`)
	return err
}

// GetMessages retrieves all messages from the database.
func GetMessages() ([]Message, error) {
	rows, err := DB.Query("SELECT id, content, created_at FROM messages ORDER BY created_at DESC")
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
func CreateMessage(content string) (*Message, error) {
	now := time.Now()
	result, err := DB.Exec("INSERT INTO messages (content, created_at) VALUES (?, ?)", content, now)
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
