package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// OrganizationService handles the single-row RENJANA organization config.
type OrganizationService struct {
	querier *queries.Querier
}

func NewOrganizationService(querier *queries.Querier) *OrganizationService {
	return &OrganizationService{querier: querier}
}

// ---------------------------------------------------------------------------
// DTOs
// ---------------------------------------------------------------------------

type Organization struct {
	ID                  int64     `json:"id"`
	Vision              string    `json:"vision"`
	Mission             string    `json:"mission"`
	History             string    `json:"history"`
	Structure           string    `json:"structure"`
	ContactEmail        string    `json:"contact_email"`
	ContactPhone        string    `json:"contact_phone"`
	Address             string    `json:"address"`
	SocialInstagram     string    `json:"social_instagram"`
	SocialTikTok        string    `json:"social_tiktok"`
	SocialYouTube       string    `json:"social_youtube"`
	SocialInstagramUrl  string    `json:"social_instagram_url"`
	SocialInstagramName string    `json:"social_instagram_name"`
	SocialTiktokUrl     string    `json:"social_tiktok_url"`
	SocialTiktokName    string    `json:"social_tiktok_name"`
	SocialYoutubeUrl    string    `json:"social_youtube_url"`
	SocialYoutubeName   string    `json:"social_youtube_name"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type UpdateOrganizationRequest struct {
	Vision              string `json:"vision"`
	Mission             string `json:"mission"`
	History             string `json:"history"`
	Structure           string `json:"structure"`
	ContactEmail        string `json:"contact_email"`
	ContactPhone        string `json:"contact_phone"`
	Address             string `json:"address"`
	SocialInstagram     string `json:"social_instagram"`
	SocialTikTok        string `json:"social_tiktok"`
	SocialYouTube       string `json:"social_youtube"`
	SocialInstagramUrl  string `json:"social_instagram_url"`
	SocialInstagramName string `json:"social_instagram_name"`
	SocialTiktokUrl     string `json:"social_tiktok_url"`
	SocialTiktokName    string `json:"social_tiktok_name"`
	SocialYoutubeUrl    string `json:"social_youtube_url"`
	SocialYoutubeName   string `json:"social_youtube_name"`
}

var ErrOrganizationNotFound = errors.New("organization not found")

// Get returns the single RENJANA org record (id=1).
func (s *OrganizationService) Get(ctx context.Context) (*Organization, error) {
	r, err := s.querier.GetOrganization(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Return empty org if not seeded
			return &Organization{ID: 1}, nil
		}
		return nil, err
	}
	return &Organization{
		ID:                  r.ID,
		Vision:              nullString(r.Vision),
		Mission:             nullString(r.Mission),
		History:             nullString(r.History),
		Structure:           nullString(r.Structure),
		ContactEmail:        nullString(r.ContactEmail),
		ContactPhone:        nullString(r.ContactPhone),
		Address:             nullString(r.Address),
		SocialInstagram:     nullString(r.SocialInstagram),
		SocialTikTok:        nullString(r.SocialTiktok),
		SocialYouTube:       nullString(r.SocialYoutube),
		SocialInstagramUrl:  nullString(r.SocialInstagramUrl),
		SocialInstagramName: nullString(r.SocialInstagramName),
		SocialTiktokUrl:     nullString(r.SocialTiktokUrl),
		SocialTiktokName:    nullString(r.SocialTiktokName),
		SocialYoutubeUrl:    nullString(r.SocialYoutubeUrl),
		SocialYoutubeName:   nullString(r.SocialYoutubeName),
		UpdatedAt:           r.UpdatedAt,
	}, nil
}

// Update creates or updates the single RENJANA org record.
func (s *OrganizationService) Update(ctx context.Context, req UpdateOrganizationRequest) error {
	_, err := s.querier.UpsertOrganization(ctx, queries.UpsertOrganizationParams{
		Vision:              sql.NullString{String: req.Vision, Valid: req.Vision != ""},
		Mission:             sql.NullString{String: req.Mission, Valid: req.Mission != ""},
		History:             sql.NullString{String: req.History, Valid: req.History != ""},
		Structure:           sql.NullString{String: req.Structure, Valid: req.Structure != ""},
		ContactEmail:        sql.NullString{String: req.ContactEmail, Valid: req.ContactEmail != ""},
		ContactPhone:        sql.NullString{String: req.ContactPhone, Valid: req.ContactPhone != ""},
		Address:             sql.NullString{String: req.Address, Valid: req.Address != ""},
		SocialInstagram:     sql.NullString{String: req.SocialInstagram, Valid: req.SocialInstagram != ""},
		SocialTiktok:        sql.NullString{String: req.SocialTikTok, Valid: req.SocialTikTok != ""},
		SocialYoutube:       sql.NullString{String: req.SocialYouTube, Valid: req.SocialYouTube != ""},
		SocialInstagramUrl:  sql.NullString{String: req.SocialInstagramUrl, Valid: req.SocialInstagramUrl != ""},
		SocialInstagramName: sql.NullString{String: req.SocialInstagramName, Valid: req.SocialInstagramName != ""},
		SocialTiktokUrl:     sql.NullString{String: req.SocialTiktokUrl, Valid: req.SocialTiktokUrl != ""},
		SocialTiktokName:    sql.NullString{String: req.SocialTiktokName, Valid: req.SocialTiktokName != ""},
		SocialYoutubeUrl:    sql.NullString{String: req.SocialYoutubeUrl, Valid: req.SocialYoutubeUrl != ""},
		SocialYoutubeName:   sql.NullString{String: req.SocialYoutubeName, Valid: req.SocialYoutubeName != ""},
	})
	return err
}

func nullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
