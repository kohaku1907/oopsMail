package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kohaku1907/oopsmail/internal/api"
	"github.com/kohaku1907/oopsmail/internal/mailbox"
	"github.com/kohaku1907/oopsmail/internal/smtp"
	"github.com/kohaku1907/oopsmail/internal/storage"
)

func main() {
	// Initialize storage
	store := storage.NewRedisStorage()

	// Initialize mailbox service
	mailboxService := mailbox.NewService(store)

	// Initialize SMTP server
	smtpServer := smtp.NewServer(mailboxService)
	go func() {
		if err := smtpServer.Start(":1025"); err != nil {
			log.Fatalf("Failed to start SMTP server: %v", err)
		}
	}()

	// Initialize HTTP server
	httpServer := api.NewServer(mailboxService)
	go func() {
		if err := httpServer.Start(":8080"); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Graceful shutdown
	smtpServer.Stop()
	httpServer.Stop()
}
