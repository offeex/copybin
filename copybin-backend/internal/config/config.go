package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
)

// Config holds all configuration for the application
type Config struct {
	Discord   oauth2.Config
	JWTSecret string
	DBURL     string
}

// New creates and returns a new Config
func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// JWT
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not set in the environment")
	}

	// Discord
	discordClientID := os.Getenv("DISCORD_CLIENT_ID")
	discordClientSecret := os.Getenv("DISCORD_CLIENT_SECRET")
	if discordClientID == "" || discordClientSecret == "" {
		return nil, fmt.Errorf("DISCORD_CLIENT_ID or DISCORD_CLIENT_SECRET is not set in the environment")
	}

	discordConfig := oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/discord/callback",
		ClientID:     discordClientID,
		ClientSecret: discordClientSecret,
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}

	// Database URL
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL is not set in the environment")
	}

	return &Config{
		Discord:   discordConfig,
		JWTSecret: jwtSecret,
		DBURL:     dbURL,
	}, nil
}
