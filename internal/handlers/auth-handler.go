package handlers

import (
	"fmt"
	"strings"

	"mrktplc-auth/internal/authorized_domains"

	"github.com/gin-gonic/gin"
)

// AuthHandler is a struct that holds the necessary dependencies for the AuthHandler
type AuthHandler struct {
	authorized_domains *authorized_domains.AuthorizedDomains
	EnterRequest
}

// NewAuthHandler creates a new instance of the AuthHandler
func NewAuthHandler(d *authorized_domains.AuthorizedDomains) *AuthHandler {
	return &AuthHandler{
		authorized_domains: d,
	}
}

type EnterRequest struct {
	Email string `json:"email"`
}

// Enter is the handler for the /auth/enter route
func (ah *AuthHandler) Enter(c *gin.Context) {
	// get email from request body
	var req EnterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	domain := strings.Split(req.Email, "@")[1]

	fmt.Println(domain)

	if !ah.authorized_domains.IsEnabled(domain) {
		c.JSON(400, gin.H{"error": "Invalid email domain"})
		return
	}

	c.JSON(200, gin.H{"message": "Hello, " + req.Email})
}