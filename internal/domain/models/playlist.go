package models

// Track represents a music track in a playlist
type Track struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Duration int    `json:"duration_ms"` // Duration in milliseconds
	URI      string `json:"uri"`
}

// Playlist represents a collection of tracks created by a user
type Playlist struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Owner       string  `json:"owner"`
	Public      bool    `json:"public"` // Indicates if the playlist is public
	Tracks      []Track `json:"tracks"`
}

// User represents a user in the system
type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}
