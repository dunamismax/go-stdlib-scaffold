package database

import "time"

// Message represents a message in the database.
type Message struct {
	ID        int64
	Content   string
	CreatedAt time.Time
}