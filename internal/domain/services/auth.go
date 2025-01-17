package services

import (
	"context"
	"spotify-backup/internal/domain/models"
)

// AuthService defines the interface for authentication-related operations.
type AuthService interface {
	// GetAuthURL returns the URL for the authentication endpoint.
	GetAuthURL() string

	// ExchangeCode exchanges an authorization code for an access token.
	ExchangeCode(ctx context.Context, code string) (string, error)

	// GetUserProfile retrieves the user's profile using the provided access token.
	GetUserProfile(ctx context.Context, token string) (*models.User, error)
}
