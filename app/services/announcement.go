package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// AnnouncementService handles CRUD for renjana_announcements.
type AnnouncementService struct {
	querier *queries.Querier
}

func NewAnnouncementService(querier *queries.Querier) *AnnouncementService {
	return &AnnouncementService{querier: querier}
}

// ---------------------------------------------------------------------------
// DTOs
// ---------------------------------------------------------------------------

// AnnouncementListItem — one row in the CRUD table.
type AnnouncementListItem struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Slug        string    `json:"slug"`
	Body        string    `json:"body"`
	CoverURL    string    `json:"cover_url"`
	AuthorID    int64     `json:"author_id"`
	PublishedAt time.Time `json:"published_at"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
}

// AnnouncementDetail — full record (for show/edit).
type AnnouncementDetail struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Slug        string    `json:"slug"`
	Body        string    `json:"body"`
	CoverURL    string    `json:"cover_url"`
	AuthorID    int64     `json:"author_id"`
	PublishedAt time.Time `json:"published_at"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateAnnouncementRequest — input for create.
type CreateAnnouncementRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	CoverURL    string `json:"cover_url"`
	AuthorID    int64  `json:"author_id"`
	PublishedAt string `json:"published_at"` // YYYY-MM-DD or RFC3339
	IsPublished bool   `json:"is_published"`
}

// UpdateAnnouncementRequest — input for update.
type UpdateAnnouncementRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	CoverURL    string `json:"cover_url"`
	PublishedAt string `json:"published_at"`
	IsPublished bool   `json:"is_published"`
}

// ---------------------------------------------------------------------------
// Errors
// ---------------------------------------------------------------------------

var (
	ErrAnnouncementNotFound = errors.New("announcement not found")
)

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func parsePublishedAt(s string, fallback time.Time) time.Time {
	if s == "" {
		return fallback
	}
	// Try common formats
	for _, layout := range []string{time.RFC3339, "2006-01-02", "2006-01-02 15:04:05"} {
		if t, err := time.Parse(layout, s); err == nil {
			return t
		}
	}
	return fallback
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	// remove non-alphanumeric and non-dash
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			out = append(out, r)
		}
	}
	return string(out)
}

// ---------------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------------

// List returns a paginated, filtered list of announcements.
func (s *AnnouncementService) List(
	ctx context.Context,
	search string,
	category string,
	isPublishedStr string,
	page, perPage int,
) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var searchArg, categoryArg, publishedArg interface{}
	if strings.TrimSpace(search) == "" {
		searchArg = ""
	} else {
		searchArg = search
	}
	if strings.TrimSpace(category) == "" {
		categoryArg = ""
	} else {
		categoryArg = category
	}
	if strings.TrimSpace(isPublishedStr) == "" {
		publishedArg = ""
	} else {
		publishedArg = isPublishedStr
	}

	rows, err := s.querier.ListAnnouncementsPaginated(ctx, queries.ListAnnouncementsPaginatedParams{
		Column1: searchArg,
		Column2: categoryArg,
		Column3: publishedArg,
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]AnnouncementListItem, 0, len(rows))
	for _, r := range rows {
		slug := ""
		if r.Slug.Valid {
			slug = r.Slug.String
		}
		body := ""
		if r.Body.Valid {
			body = r.Body.String
		}
		cover := ""
		if r.CoverUrl.Valid {
			cover = r.CoverUrl.String
		}
		var authorID int64
		if r.AuthorID.Valid {
			authorID = r.AuthorID.Int64
		}
		items = append(items, AnnouncementListItem{
			ID:          r.ID,
			Title:       r.Title,
			Content:     r.Content,
			Category:    r.Category,
			Slug:        slug,
			Body:        body,
			CoverURL:    cover,
			AuthorID:    authorID,
			PublishedAt: r.PublishedAt,
			IsPublished: r.IsPublished,
			CreatedAt:   r.CreatedAt,
		})
	}

	total, err := s.querier.CountAnnouncementsFiltered(ctx, queries.CountAnnouncementsFilteredParams{
		Column1: searchArg,
		Column2: categoryArg,
		Column3: publishedArg,
	})
	if err != nil {
		return nil, err
	}

	return BuildPagination(items, page, perPage, total), nil
}

// Get returns a single announcement detail.
func (s *AnnouncementService) Get(ctx context.Context, id int64) (*AnnouncementDetail, error) {
	r, err := s.querier.GetAnnouncementByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAnnouncementNotFound
		}
		return nil, err
	}

	cat := r.Category
	slug := ""
	if r.Slug.Valid {
		slug = r.Slug.String
	}
	body := ""
	if r.Body.Valid {
		body = r.Body.String
	}
	cover := ""
	if r.CoverUrl.Valid {
		cover = r.CoverUrl.String
	}
	var authorID int64
	if r.AuthorID.Valid {
		authorID = r.AuthorID.Int64
	}

	return &AnnouncementDetail{
		ID:          r.ID,
		Title:       r.Title,
		Content:     r.Content,
		Category:    cat,
		Slug:        slug,
		Body:        body,
		CoverURL:    cover,
		AuthorID:    authorID,
		PublishedAt: r.PublishedAt,
		IsPublished: r.IsPublished,
		CreatedAt:   r.CreatedAt,
	}, nil
}

// Create inserts a new announcement.
func (s *AnnouncementService) Create(ctx context.Context, req CreateAnnouncementRequest) (*AnnouncementDetail, error) {
	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("judul wajib diisi")
	}
	if req.Category == "" {
		req.Category = "Pengumuman"
	}
	if req.Slug == "" {
		req.Slug = slugify(req.Title)
	}
	if req.Body == "" {
		req.Body = req.Content
	}
	pubAt := parsePublishedAt(req.PublishedAt, time.Now())

	id, err := s.querier.CreateAnnouncement(ctx, queries.CreateAnnouncementParams{
		Title:       req.Title,
		Content:     req.Content,
		Category:    req.Category,
		Slug:        sql.NullString{String: req.Slug, Valid: req.Slug != ""},
		Body:        sql.NullString{String: req.Body, Valid: req.Body != ""},
		CoverUrl:    sql.NullString{String: req.CoverURL, Valid: req.CoverURL != ""},
		AuthorID:    sql.NullInt64{Int64: req.AuthorID, Valid: req.AuthorID != 0},
		PublishedAt: pubAt,
		IsPublished: req.IsPublished,
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

// Update modifies an existing announcement.
func (s *AnnouncementService) Update(ctx context.Context, id int64, req UpdateAnnouncementRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("judul wajib diisi")
	}
	if req.Slug == "" {
		req.Slug = slugify(req.Title)
	}
	if req.Body == "" {
		req.Body = req.Content
	}
	pubAt := parsePublishedAt(req.PublishedAt, time.Now())

	rows, err := s.querier.UpdateAnnouncement(ctx, queries.UpdateAnnouncementParams{
		Title:       req.Title,
		Content:     req.Content,
		Category:    req.Category,
		Slug:        sql.NullString{String: req.Slug, Valid: req.Slug != ""},
		Body:        sql.NullString{String: req.Body, Valid: req.Body != ""},
		CoverUrl:    sql.NullString{String: req.CoverURL, Valid: req.CoverURL != ""},
		PublishedAt: pubAt,
		IsPublished: req.IsPublished,
		ID:          id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrAnnouncementNotFound
	}
	return nil
}

// Delete removes an announcement.
func (s *AnnouncementService) Delete(ctx context.Context, id int64) error {
	rows, err := s.querier.DeleteAnnouncement(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrAnnouncementNotFound
	}
	return nil
}

// ListByCategory returns announcements filtered by category (for portal views).
func (s *AnnouncementService) ListByCategory(ctx context.Context, category string, limit int) ([]AnnouncementListItem, error) {
	rows, err := s.querier.ListAnnouncementsPaginated(ctx, queries.ListAnnouncementsPaginatedParams{
		Column1: "",
		Column2: sql.NullString{String: category, Valid: category != ""},
		Column3: sql.NullString{String: "1", Valid: true}, // only published
		Limit:   int64(limit),
		Offset:  0,
	})
	if err != nil {
		return nil, err
	}

	items := make([]AnnouncementListItem, 0, len(rows))
	for _, r := range rows {
		slug := ""
		if r.Slug.Valid {
			slug = r.Slug.String
		}
		body := ""
		if r.Body.Valid {
			body = r.Body.String
		}
		cover := ""
		if r.CoverUrl.Valid {
			cover = r.CoverUrl.String
		}
		var authorID int64
		if r.AuthorID.Valid {
			authorID = r.AuthorID.Int64
		}
		items = append(items, AnnouncementListItem{
			ID:          r.ID,
			Title:       r.Title,
			Content:     r.Content,
			Category:    r.Category,
			Slug:        slug,
			Body:        body,
			CoverURL:    cover,
			AuthorID:    authorID,
			PublishedAt: r.PublishedAt,
			IsPublished: r.IsPublished,
			CreatedAt:   r.CreatedAt,
		})
	}
	return items, nil
}
