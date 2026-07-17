package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
)

type AuthService struct {
	querier       *queries.Querier
	sessionSecret string
	argon2Params  Argon2Params
	oauthConfig   *oauth2.Config
}

type AuthServiceConfig struct {
	SessionSecret      string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	Argon2Params       Argon2Params
}

func NewAuthService(querier *queries.Querier, cfg AuthServiceConfig) *AuthService {
	argonParams := cfg.Argon2Params
	if argonParams.Memory == 0 {
		argonParams.Memory = DefaultArgon2Params.Memory
	}
	if argonParams.Iterations == 0 {
		argonParams.Iterations = DefaultArgon2Params.Iterations
	}
	if argonParams.Parallelism == 0 {
		argonParams.Parallelism = DefaultArgon2Params.Parallelism
	}
	if argonParams.SaltLength == 0 {
		argonParams.SaltLength = DefaultArgon2Params.SaltLength
	}
	if argonParams.KeyLength == 0 {
		argonParams.KeyLength = DefaultArgon2Params.KeyLength
	}
	return &AuthService{
		querier:       querier,
		sessionSecret: cfg.SessionSecret,
		argon2Params:  argonParams,
		oauthConfig: &oauth2.Config{
			ClientID:     cfg.GoogleClientID,
			ClientSecret: cfg.GoogleClientSecret,
			RedirectURL:  cfg.GoogleRedirectURL,
			Scopes:       []string{"email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

// GetOAuthConfig returns the OAuth config for Google
func (s *AuthService) GetOAuthConfig() *oauth2.Config {
	return s.oauthConfig
}

// ProcessGoogleToken exchanges the OAuth code for a token and returns user info
func (s *AuthService) ProcessGoogleToken(ctx context.Context, code string) (*models.User, error) {
	// Exchange code for token
	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// Get user info from Google
	oauthClient := s.oauthConfig.Client(ctx, token)
	resp, err := oauthClient.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, ErrInvalidToken
	}
	defer resp.Body.Close()

	var googleUser struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Picture  string `json:"picture"`
		Verified bool   `json:"verified_email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, ErrInvalidToken
	}

	// Check if user exists by Google ID
	user, err := s.querier.GetUserByGoogleID(ctx, googleUser.ID)
	if err == nil {
		// Migrate external avatar to local if needed
		if user.Avatar != "" && !strings.HasPrefix(user.Avatar, "/storage/") {
			if localPath, dlErr := s.downloadAndSaveAvatar(ctx, user.Avatar, googleUser.ID); dlErr != nil {
				slog.Warn("failed to download avatar for existing user", "user_id", user.ID, "error", dlErr)
			} else if upErr := s.querier.UpdateUserAvatar(ctx, user.ID, localPath); upErr != nil {
				slog.Warn("failed to update avatar path", "user_id", user.ID, "error", upErr)
			} else {
				user.Avatar = localPath
			}
		}
		return user, nil
	}
	if !errors.Is(err, queries.ErrUserNotFound) {
		return nil, err
	}

	// Check if user exists by email
	user, err = s.querier.GetUserByEmail(ctx, googleUser.Email)
	if err == nil {
		// Link Google ID to existing account
		user.GoogleID = sql.NullString{String: googleUser.ID, Valid: true}

		// Migrate external avatar to local if needed
		if user.Avatar != "" && !strings.HasPrefix(user.Avatar, "/storage/") {
			if localPath, dlErr := s.downloadAndSaveAvatar(ctx, user.Avatar, googleUser.ID); dlErr != nil {
				slog.Warn("failed to download avatar for existing user", "user_id", user.ID, "error", dlErr)
			} else {
				user.Avatar = localPath
			}
		}

		if err := s.querier.UpdateUser(ctx, user); err != nil {
			return nil, err
		}
		return user, nil
	}

	// Download avatar to local storage
	localAvatar := googleUser.Picture
	if localPath, dlErr := s.downloadAndSaveAvatar(ctx, googleUser.Picture, googleUser.ID); dlErr != nil {
		slog.Warn("failed to download Google avatar", "email", googleUser.Email, "error", dlErr)
	} else {
		localAvatar = localPath
	}

	// Create new user
	newUser := &models.User{
		Email: googleUser.Email,
		Name:  googleUser.Name,
		GoogleID: sql.NullString{
			String: googleUser.ID,
			Valid:  true,
		},
		Avatar:        localAvatar,
		EmailVerified: googleUser.Verified,
		Role:          models.RoleRelawan,
	}

	if err := s.querier.CreateUserWithGoogleID(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// Register creates a new user with email/password
func (s *AuthService) Register(name, email, password string) (*models.User, error) {
	// Check if user already exists
	_, err := s.querier.GetUserByEmail(context.Background(), email)
	if err == nil {
		return nil, queries.ErrUserAlreadyExists
	}
	if !errors.Is(err, queries.ErrUserNotFound) {
		return nil, err
	}

	// Hash password with configured argon2id params
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Email: email,
		Name:  name,
		Password: sql.NullString{
			String: hashedPassword,
			Valid:  true,
		},
		Role:          models.RoleRelawan,
		EmailVerified: false,
	}

	if err := s.querier.CreateUser(context.Background(), user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user with email/password
func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.querier.GetUserByEmail(context.Background(), email)
	if err != nil {
		if errors.Is(err, queries.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// Check password - user must have a password (not OAuth-only user)
	if !user.Password.Valid {
		return nil, ErrInvalidCredentials
	}

	if !checkPassword(user.Password.String, password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(id int64) (*models.User, error) {
	return s.querier.GetUserByID(context.Background(), id)
}

// hashPassword hashes a password using argon2id with the configured params.
func (s *AuthService) hashPassword(password string) (string, error) {
	return generateFromPassword(password, s.argon2Params)
}

// downloadAndSaveAvatar downloads an external avatar image and saves it to local storage.
// Returns the local URL path (/storage/avatars/<filename>) or an error.
func (s *AuthService) downloadAndSaveAvatar(ctx context.Context, pictureURL, googleID string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pictureURL, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	// Google avatar selalu JPEG
	filename := googleID + ".jpg"

	// Ensure directory exists
	avatarDir := "./storage/avatars"
	if err := os.MkdirAll(avatarDir, 0750); err != nil {
		return "", fmt.Errorf("mkdir: %w", err)
	}

	filePath := filepath.Join(avatarDir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return "", fmt.Errorf("write file: %w", err)
	}

	slog.Info("saved avatar locally", "filename", filename, "source", pictureURL)
	return "/storage/avatars/" + filename, nil
}

// GetOAuthURL returns the OAuth URL for Google login
func (s *AuthService) GetOAuthURL(state string) string {
	return s.oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// ValidateState validates the OAuth state parameter
func (s *AuthService) ValidateState(state, expected string) bool {
	return state == expected
}
