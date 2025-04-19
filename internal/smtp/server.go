package smtp

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/kohaku1907/oopsmail/internal/mailbox"
	"github.com/kohaku1907/oopsmail/internal/storage"
)

type Server struct {
	backend *Backend
	server  *smtp.Server
}

func NewServer(mailboxService *mailbox.Service) *Server {
	backend := &Backend{
		mailboxService: mailboxService,
	}

	server := smtp.NewServer(backend)
	server.Addr = ":1025"
	server.Domain = "oopsmail.com"
	server.AllowInsecureAuth = true

	return &Server{
		backend: backend,
		server:  server,
	}
}

func (s *Server) Start(addr string) error {
	s.server.Addr = addr
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Close()
}

type Backend struct {
	mailboxService *mailbox.Service
}

func (b *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{
		backend: b,
	}, nil
}

type Session struct {
	backend *Backend
	from    string
	to      string
}

func (s *Session) AuthPlain(username, password string) error {
	return nil // Allow all auth attempts
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	s.to = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	// Parse email content
	// This is a simplified version - in a real implementation,
	// you'd want to use a proper email parser
	content, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// Extract mailbox ID from the "to" address
	// Format: <id>@oopsmail.com
	parts := strings.Split(s.to, "@")
	if len(parts) != 2 {
		return fmt.Errorf("invalid email address format")
	}
	mailboxID := parts[0]

	// Create email object
	email := &storage.Email{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		From:      s.from,
		To:        s.to,
		Subject:   "No Subject", // You'd parse this from the email content
		Body:      string(content),
		CreatedAt: time.Now(),
	}

	// Store the email
	ctx := context.Background()
	return s.backend.mailboxService.StoreEmail(ctx, mailboxID, email)
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
