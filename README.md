# Spotify Playlist Backup Service

A robust Go application that allows users to backup their Spotify playlists and export them to CSV format. Built with clean architecture principles and modern Go practices.

## 🎵 Features

- **Playlist Backup**: Download any public Spotify playlist
- **Batch Export**: Export multiple playlists at once
- **OAuth Authentication**: Secure login with Spotify credentials
- **CSV Export**: Download playlists in CSV format for easy importing elsewhere
- **User-Friendly Interface**: Clean web interface for easy playlist management
- **Private Playlist Support**: Access and backup private playlists (requires user authentication)

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- A Spotify Developer Account
- Docker (optional, for containerized deployment)

### Environment Setup

1. Create a Spotify application in your [Spotify Developer Dashboard](https://developer.spotify.com/dashboard)

2. Clone the repository:
```bash
git clone https://github.com/letsmakecakes/spotify-backup.git
cd spotify-backup
```

3. Copy the example environment file:
```bash
cp .env.example .env
```

4. Update the `.env` file with your Spotify credentials:
```env
SPOTIFY_CLIENT_ID=your_client_id
SPOTIFY_CLIENT_SECRET=your_client_secret
SPOTIFY_REDIRECT_URI=http://localhost:8080/callback
PORT=8080
```

### Running the Application

#### Local Development
```bash
# Install dependencies
go mod download

# Run the application
go run cmd/server/main.go
```

#### Using Docker
```bash
# Build the Docker image
docker build -t spotify-backup .

# Run the container
docker run -p 8080:8080 --env-file .env spotify-backup
```

The application will be available at `http://localhost:8080`

## 📁 Project Structure

```
spotify-backup/
├── cmd/                    # Application entrypoints
├── internal/              # Private application code
│   ├── api/              # HTTP handlers and routing
│   ├── domain/           # Business logic and models
│   ├── platform/         # External service integrations
│   └── storage/          # Data persistence
├── pkg/                   # Public packages
├── web/                   # Frontend assets
└── ...
```

## 🔒 Authentication

The application uses Spotify's OAuth 2.0 flow for authentication:

1. Users click "Login with Spotify"
2. They're redirected to Spotify's login page
3. After successful authentication, they're redirected back to the application
4. The application receives an access token to make API requests

## 🛠️ Usage

1. Visit the homepage at `http://localhost:8080`
2. Log in with your Spotify account (if accessing private playlists)
3. Enter a playlist URL or ID in the input field
4. Click "Backup" to download the playlist as CSV

### Backing Up Public Playlists
- No login required
- Simply paste the playlist URL or ID
- Click "Download CSV"

### Backing Up Private Playlists
- Login required
- Click "Backup All My Playlists" to export all your playlists
- Individual playlists can be selected for backup

## 🧪 Testing

Run the test suite:
```bash
go test ./...
```

Run with coverage:
```bash
go test ./... -cover
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Spotify Web API Documentation
- The Go Community
- All Contributors

## 📮 Contact

Your Name - [@letsmakeckes](https://twitter.com/letsmakecakes)

Project Link: [https://github.com/letsmakecakes/spotify-backup](https://github.com/yourusername/spotify-backup)

## 🚧 Roadmap

- [ ] Support for other music streaming services
- [ ] Playlist import functionality
- [ ] Automated backup scheduling
- [ ] Multiple export formats (JSON, XML)
- [ ] Playlist difference detection
- [ ] Collaborative playlist handling