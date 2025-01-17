package services

import (
	"context"
	"spotify-backup/internal/domain/models"
)

// PlaylistService defines the interface for playlist-related operations.
type PlaylistService interface {
	// GetPlaylist retrieves a playlist by its ID.
	GetPlaylist(ctx context.Context, id string) (*models.Playlist, error)

	// GetUserPlaylists retrieves all playlists for a given user.
	GetUserPlaylists(ctx context.Context, userID string) ([]models.Playlist, error)

	// ExportToCSV exports a playlist to a CSV file.
	ExportToCSV(ctx context.Context, playlist *models.Playlist) (string, error)
}
