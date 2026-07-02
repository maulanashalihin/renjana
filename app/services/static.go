package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// StaticService handles read-only operations for content tables
// (education, media, documents, innovations).
type StaticService struct {
	querier *queries.Querier
}

func NewStaticService(querier *queries.Querier) *StaticService {
	return &StaticService{querier: querier}
}

// ---------------------------------------------------------------------------
// Education
// ---------------------------------------------------------------------------

type EducationItem struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Category        string    `json:"category"`
	Body            string    `json:"body"`
	AgeGroup        string    `json:"age_group"`
	DurationMinutes int64     `json:"duration_minutes"`
	IsPublished     bool      `json:"is_published"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func nullToInt64(n sql.NullInt64) int64 {
	if n.Valid {
		return n.Int64
	}
	return 0
}

// ListEducation returns paginated education articles.
func (s *StaticService) ListEducation(ctx context.Context, category string, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var catArg interface{}
	if strings.TrimSpace(category) == "" {
		catArg = ""
	} else {
		catArg = category
	}

	rows, err := s.querier.ListEducationPaginated(ctx, queries.ListEducationPaginatedParams{
		Column1: catArg,
		Column2: int64(1), // only published
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]EducationItem, 0, len(rows))
	for _, r := range rows {
		age := ""
		if r.AgeGroup.Valid {
			age = r.AgeGroup.String
		}
		items = append(items, EducationItem{
			ID:              r.ID,
			Title:           r.Title,
			Category:        r.Category,
			Body:            r.Body,
			AgeGroup:        age,
			DurationMinutes: nullToInt64(r.DurationMinutes),
			IsPublished:     r.IsPublished,
			CreatedAt:       r.CreatedAt,
			UpdatedAt:       r.UpdatedAt,
		})
	}

	total, err := s.querier.CountEducationFiltered(ctx, queries.CountEducationFilteredParams{
		Column1: catArg,
		Column2: int64(1),
	})
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}

// GetEducation returns a single article.
func (s *StaticService) GetEducation(ctx context.Context, id int64) (*EducationItem, error) {
	r, err := s.querier.GetEducationByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	age := ""
	if r.AgeGroup.Valid {
		age = r.AgeGroup.String
	}
	return &EducationItem{
		ID:              r.ID,
		Title:           r.Title,
		Category:        r.Category,
		Body:            r.Body,
		AgeGroup:        age,
		DurationMinutes: nullToInt64(r.DurationMinutes),
		IsPublished:     r.IsPublished,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}, nil
}

// ---------------------------------------------------------------------------
// Media (Galeri)
// ---------------------------------------------------------------------------

type MediaItem struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	FileURL     string    `json:"file_url"`
	MediaType   string    `json:"media_type"`
	ActivityID  int64     `json:"activity_id"`
	DistrictID  int64     `json:"district_id"`
	Caption     string    `json:"caption"`
	UploadedBy  int64     `json:"uploaded_by"`
	UploadedAt  time.Time `json:"uploaded_at"`
	IsPublished bool      `json:"is_published"`
}

// ListMedia returns paginated gallery items.
func (s *StaticService) ListMedia(ctx context.Context, mediaType string, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var typeArg interface{}
	if strings.TrimSpace(mediaType) == "" {
		typeArg = ""
	} else {
		typeArg = mediaType
	}

	rows, err := s.querier.ListMediaPaginated(ctx, queries.ListMediaPaginatedParams{
		Column1: typeArg,
		Column2: int64(1),
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]MediaItem, 0, len(rows))
	for _, r := range rows {
		caption := ""
		if r.Caption.Valid {
			caption = r.Caption.String
		}
		items = append(items, MediaItem{
			ID:          r.ID,
			Title:       r.Title,
			FileURL:     r.FileUrl,
			MediaType:   r.MediaType,
			ActivityID:  r.ActivityID.Int64,
			DistrictID:  r.DistrictID.Int64,
			Caption:     caption,
			UploadedBy:  r.UploadedBy.Int64,
			UploadedAt:  r.UploadedAt,
			IsPublished: r.IsPublished,
		})
	}

	total, err := s.querier.CountMediaFiltered(ctx, queries.CountMediaFilteredParams{
		Column1: typeArg,
		Column2: int64(1),
	})
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}

// ---------------------------------------------------------------------------
// Documents
// ---------------------------------------------------------------------------

type DocumentItem struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	FileURL     string    `json:"file_url"`
	Category    string    `json:"category"`
	Version     int64     `json:"version"`
	FileSize    int64     `json:"file_size"`
	Description string    `json:"description"`
	UploadedBy  int64     `json:"uploaded_by"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

// ListDocuments returns paginated documents.
func (s *StaticService) ListDocuments(ctx context.Context, category string, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var catArg interface{}
	if strings.TrimSpace(category) == "" {
		catArg = ""
	} else {
		catArg = category
	}

	rows, err := s.querier.ListDocumentsPaginated(ctx, queries.ListDocumentsPaginatedParams{
		Column1: catArg,
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]DocumentItem, 0, len(rows))
	for _, r := range rows {
		desc := ""
		if r.Description.Valid {
			desc = r.Description.String
		}
		items = append(items, DocumentItem{
			ID:          r.ID,
			Title:       r.Title,
			FileURL:     r.FileUrl,
			Category:    r.Category,
			Version:     r.Version,
			FileSize:    nullToInt64(r.FileSize),
			Description: desc,
			UploadedBy:  r.UploadedBy.Int64,
			UploadedAt:  r.UploadedAt,
		})
	}

	total, err := s.querier.CountDocumentsFiltered(ctx, catArg)
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}


