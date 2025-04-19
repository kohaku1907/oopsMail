package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/kohaku1907/oopsmail/internal/api"
	"github.com/kohaku1907/oopsmail/internal/mailbox"
	"github.com/kohaku1907/oopsmail/internal/smtp"
	"github.com/kohaku1907/oopsmail/internal/storage"
	"github.com/kohaku1907/oopsmail/internal/web"
)

func main() {
	// Initialize storage
	redisStorage := storage.NewRedisStorage()

	// Initialize mailbox service
	mailboxService := mailbox.NewService(redisStorage)

	// Initialize SMTP server
	smtpServer := smtp.NewServer(mailboxService)
	go func() {
		if err := smtpServer.Start(":1025"); err != nil {
			log.Fatalf("Failed to start SMTP server: %v", err)
		}
	}()

	// Initialize web handler
	webHandler, err := web.NewHandler()
	if err != nil {
		log.Fatalf("Failed to initialize web handler: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./internal/web/static")

	// Web routes
	router.GET("/", func(c *gin.Context) {
		webHandler.Home(c.Writer, c.Request)
	})
	router.GET("/create", func(c *gin.Context) {
		webHandler.CreateMailbox(c.Writer, c.Request)
	})
	router.GET("/view", func(c *gin.Context) {
		webHandler.ViewEmails(c.Writer, c.Request)
	})

	// API routes
	apiServer := api.NewServer(mailboxService)
	router.POST("/api/mailbox", apiServer.CreateMailbox)
	router.GET("/api/mailbox/:id", apiServer.GetEmails)

	// Start HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")

	// Graceful shutdown
	smtpServer.Stop()
}
