package handlers

import (
	"strings"

	"mrktplc-auth/internal/authorized_domains"
	"mrktplc-auth/internal/token"

	"github.com/gin-gonic/gin"
)

// AuthHandler is a struct that holds the necessary dependencies for the AuthHandler
type AuthHandler struct {
	authorized_domains *authorized_domains.AuthorizedDomains
	token *token.Token
	EnterRequest
}

// NewAuthHandler creates a new instance of the AuthHandler
func NewAuthHandler(d *authorized_domains.AuthorizedDomains, t *token.Token) *AuthHandler {
	return &AuthHandler{
		authorized_domains: d,
		token: t,
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

	if !ah.authorized_domains.IsEnabled(domain) {
		c.JSON(400, gin.H{"error": "Invalid email domain"})
		return
	}

	jwt, err := ah.token.CreateToken(req.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(200, gin.H{"message": "Hello, " + req.Email, "token": jwt})
}