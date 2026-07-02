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
	CoverImage      string    `json:"cover_image"`
	PassingScore    int64     `json:"passing_score"`
	TotalModules    int64     `json:"total_modules"`
	IsCourse        bool      `json:"is_course"`
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
		cover := ""
		if r.CoverImage.Valid {
			cover = r.CoverImage.String
		}
		items = append(items, EducationItem{
			ID:              r.ID,
			Title:           r.Title,
			Category:        r.Category,
			Body:            r.Body,
			AgeGroup:        age,
			DurationMinutes: nullToInt64(r.DurationMinutes),
			IsPublished:     r.IsPublished,
			CoverImage:      cover,
			PassingScore:    r.PassingScore,
			TotalModules:    r.TotalModules,
			IsCourse:        r.IsCourse,
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
	cover := ""
	if r.CoverImage.Valid {
		cover = r.CoverImage.String
	}
	return &EducationItem{
		ID:              r.ID,
		Title:           r.Title,
		Category:        r.Category,
		Body:            r.Body,
		AgeGroup:        age,
		DurationMinutes: nullToInt64(r.DurationMinutes),
		IsPublished:     r.IsPublished,
		CoverImage:      cover,
		PassingScore:    r.PassingScore,
		TotalModules:    r.TotalModules,
		IsCourse:        r.IsCourse,
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

// CreateMedia inserts a new gallery item.
func (s *StaticService) CreateMedia(ctx context.Context, title, fileURL, mediaType, caption string, uploadedBy int64, isPublished bool) (*MediaItem, error) {
	var uploader sql.NullInt64
	if uploadedBy > 0 {
		uploader = sql.NullInt64{Int64: uploadedBy, Valid: true}
	}
	r, err := s.querier.CreateMedia(ctx, queries.CreateMediaParams{
		Title:       title,
		FileUrl:     fileURL,
		MediaType:   mediaType,
		Caption:     sql.NullString{String: caption, Valid: caption != ""},
		UploadedBy:  uploader,
		IsPublished: isPublished,
	})
	if err != nil {
		return nil, err
	}
	captionStr := ""
	if r.Caption.Valid {
		captionStr = r.Caption.String
	}
	return &MediaItem{
		ID:          r.ID,
		Title:       r.Title,
		FileURL:     r.FileUrl,
		MediaType:   r.MediaType,
		ActivityID:  r.ActivityID.Int64,
		DistrictID:  r.DistrictID.Int64,
		Caption:     captionStr,
		UploadedBy:  r.UploadedBy.Int64,
		UploadedAt:  r.UploadedAt,
		IsPublished: r.IsPublished,
	}, nil
}

// UpdateMedia updates a gallery item.
func (s *StaticService) UpdateMedia(ctx context.Context, id int64, title, fileURL, mediaType, caption string, isPublished bool) (*MediaItem, error) {
	r, err := s.querier.UpdateMedia(ctx, queries.UpdateMediaParams{
		Title:       title,
		FileUrl:     fileURL,
		MediaType:   mediaType,
		Caption:     sql.NullString{String: caption, Valid: caption != ""},
		IsPublished: isPublished,
		ID:          id,
	})
	if err != nil {
		return nil, err
	}
	captionStr := ""
	if r.Caption.Valid {
		captionStr = r.Caption.String
	}
	return &MediaItem{
		ID:          r.ID,
		Title:       r.Title,
		FileURL:     r.FileUrl,
		MediaType:   r.MediaType,
		ActivityID:  r.ActivityID.Int64,
		DistrictID:  r.DistrictID.Int64,
		Caption:     captionStr,
		UploadedBy:  r.UploadedBy.Int64,
		UploadedAt:  r.UploadedAt,
		IsPublished: r.IsPublished,
	}, nil
}

// DeleteMedia deletes a gallery item.
func (s *StaticService) DeleteMedia(ctx context.Context, id int64) error {
	return s.querier.DeleteMedia(ctx, id)
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

// CreateDocumentRequest — input for creating a document.
type CreateDocumentRequest struct {
	Title       string `json:"title"`
	FileURL     string `json:"file_url"`
	Category    string `json:"category"`
	Version     int64  `json:"version"`
	FileSize    int64  `json:"file_size"`
	Description string `json:"description"`
	UploadedBy  int64  `json:"uploaded_by"`
}

// UpdateDocumentRequest — input for updating a document.
type UpdateDocumentRequest struct {
	Title       string `json:"title"`
	FileURL     string `json:"file_url"`
	Category    string `json:"category"`
	Version     int64  `json:"version"`
	FileSize    int64  `json:"file_size"`
	Description string `json:"description"`
}

// CreateDocument creates a new document record.
func (s *StaticService) CreateDocument(ctx context.Context, req CreateDocumentRequest) (*DocumentItem, error) {
	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("judul dokumen wajib diisi")
	}
	if strings.TrimSpace(req.FileURL) == "" {
		return nil, errors.New("file wajib diupload")
	}
	if strings.TrimSpace(req.Category) == "" {
		req.Category = "SOP"
	}
	if req.Version == 0 {
		req.Version = 1
	}

	id, err := s.querier.CreateDocument(ctx, queries.CreateDocumentParams{
		Title:       req.Title,
		FileUrl:     req.FileURL,
		Category:    req.Category,
		Version:     req.Version,
		FileSize:    sql.NullInt64{Int64: req.FileSize, Valid: req.FileSize > 0},
		Description: sql.NullString{String: req.Description, Valid: strings.TrimSpace(req.Description) != ""},
		UploadedBy:  sql.NullInt64{Int64: req.UploadedBy, Valid: req.UploadedBy > 0},
	})
	if err != nil {
		return nil, err
	}
	return s.GetDocumentByID(ctx, id)
}

// UpdateDocument updates an existing document.
func (s *StaticService) UpdateDocument(ctx context.Context, id int64, req UpdateDocumentRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("judul dokumen wajib diisi")
	}
	if strings.TrimSpace(req.FileURL) == "" {
		return errors.New("file wajib diupload")
	}
	_, err := s.querier.UpdateDocument(ctx, queries.UpdateDocumentParams{
		Title:       req.Title,
		FileUrl:     req.FileURL,
		Category:    req.Category,
		Version:     req.Version,
		FileSize:    sql.NullInt64{Int64: req.FileSize, Valid: req.FileSize > 0},
		Description: sql.NullString{String: req.Description, Valid: strings.TrimSpace(req.Description) != ""},
		ID:          id,
	})
	return err
}

// DeleteDocument removes a document.
func (s *StaticService) DeleteDocument(ctx context.Context, id int64) error {
	_, err := s.querier.DeleteDocument(ctx, id)
	return err
}

// GetDocumentByID returns a single document.
func (s *StaticService) GetDocumentByID(ctx context.Context, id int64) (*DocumentItem, error) {
	r, err := s.querier.GetDocumentByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	desc := ""
	if r.Description.Valid {
		desc = r.Description.String
	}
	return &DocumentItem{
		ID:          r.ID,
		Title:       r.Title,
		FileURL:     r.FileUrl,
		Category:    r.Category,
		Version:     r.Version,
		FileSize:    nullToInt64(r.FileSize),
		Description: desc,
		UploadedBy:  r.UploadedBy.Int64,
		UploadedAt:  r.UploadedAt,
	}, nil
}
