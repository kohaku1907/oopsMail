package mailbox

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/kohaku1907/oopsmail/internal/storage"
)

type Service struct {
	store storage.Storage
}

func NewService(store storage.Storage) *Service {
	return &Service{
		store: store,
	}
}

// CreateMailbox creates a new temporary mailbox with a random ID
func (s *Service) CreateMailbox(ctx context.Context) (string, error) {
	id, err := generateRandomID()
	if err != nil {
		return "", err
	}

	// Set mailbox to expire after 1 hour
	expiration := time.Hour
	if err := s.store.CreateMailbox(ctx, id, expiration); err != nil {
		return "", err
	}

	return id, nil
}

// StoreEmail stores a new email in the specified mailbox
func (s *Service) StoreEmail(ctx context.Context, mailboxID string, email *storage.Email) error {
	exists, err := s.store.MailboxExists(ctx, mailboxID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrMailboxNotFound
	}

	return s.store.StoreEmail(ctx, mailboxID, email)
}

// GetEmails retrieves all emails for a given mailbox
func (s *Service) GetEmails(ctx context.Context, mailboxID string) ([]*storage.Email, error) {
	exists, err := s.store.MailboxExists(ctx, mailboxID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrMailboxNotFound
	}

	return s.store.GetEmails(ctx, mailboxID)
}

// generateRandomID creates a random string to use as a mailbox ID
func generateRandomID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

var ErrMailboxNotFound = fmt.Errorf("mailbox not found")
