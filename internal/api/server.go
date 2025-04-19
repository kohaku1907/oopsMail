package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kohaku1907/oopsmail/internal/mailbox"
)

type Server struct {
	mailboxService *mailbox.Service
	router         *gin.Engine
}

func NewServer(mailboxService *mailbox.Service) *Server {
	router := gin.Default()

	server := &Server{
		mailboxService: mailboxService,
		router:         router,
	}

	// Register routes
	router.POST("/mailbox", server.CreateMailbox)
	router.GET("/mailbox/:id", server.GetEmails)

	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Stop() error {
	// Add any cleanup logic here
	return nil
}

type createMailboxResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (s *Server) CreateMailbox(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := s.mailboxService.CreateMailbox(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := createMailboxResponse{
		ID:        id,
		Email:     id + "@oopsmail.com",
		ExpiresAt: time.Now().Add(time.Hour),
	}

	c.JSON(http.StatusCreated, response)
}

func (s *Server) GetEmails(c *gin.Context) {
	ctx := c.Request.Context()
	mailboxID := c.Param("id")

	emails, err := s.mailboxService.GetEmails(ctx, mailboxID)
	if err != nil {
		if err == mailbox.ErrMailboxNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "mailbox not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emails)
}
