package storage

import (
	"context"
	"time"
)

// Email represents a received email
type Email struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// Storage defines the interface for storing and retrieving emails
type Storage interface {
	// CreateMailbox creates a new mailbox with the given ID and expiration time
	CreateMailbox(ctx context.Context, id string, expiration time.Duration) error

	// StoreEmail stores an email in the specified mailbox
	StoreEmail(ctx context.Context, mailboxID string, email *Email) error

	// GetEmails retrieves all emails for a given mailbox
	GetEmails(ctx context.Context, mailboxID string) ([]*Email, error)

	// DeleteMailbox removes a mailbox and all its emails
	DeleteMailbox(ctx context.Context, mailboxID string) error

	// MailboxExists checks if a mailbox exists
	MailboxExists(ctx context.Context, mailboxID string) (bool, error)
}
