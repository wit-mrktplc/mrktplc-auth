package authorized_domains

import (
	"os"
	"slices"
	"strings"
)

// AuthorizedDomains is a struct that holds the necessary dependencies for the AuthorizedDomains
type AuthorizedDomains struct {
	EnabledDomains []string
}

// NewAuthorizedDomains creates a new instance of the AuthorizedDomains
func NewAuthorizedDomains() *AuthorizedDomains {
	return &AuthorizedDomains{
		EnabledDomains: strings.Split(os.Getenv("AUTHORIZED_DOMAINS"), ","),
	}
}

// IsEnabled checks if a AuthorizedDomains is enabled
func (d *AuthorizedDomains) IsEnabled(authorized_domains string) bool {
	return slices.Contains(d.EnabledDomains, authorized_domains)
}