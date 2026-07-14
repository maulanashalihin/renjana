package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// SchoolService handles CRUD for renjana_schools.
type SchoolService struct {
	querier *queries.Querier
}

func NewSchoolService(querier *queries.Querier) *SchoolService {
	return &SchoolService{querier: querier}
}

// SchoolSearchResult represents a school search result (lighter payload).
type SchoolSearchResult struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Level     string `json:"level"`
	Status    string `json:"status"`
	Kecamatan string `json:"kecamatan"`
}

// SchoolItem represents a full school record for admin CRUD.
type SchoolItem struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Level     string `json:"level"`
	Status    string `json:"status"`
	Kecamatan string `json:"kecamatan"`
	IsActive  bool   `json:"is_active"`
}

// SchoolInput for create/update.
type SchoolInput struct {
	Name      string `json:"name"`
	Level     string `json:"level"`
	Status    string `json:"status"`
	Kecamatan string `json:"kecamatan"`
}

// Search returns schools matching the query (for autocomplete).
func (s *SchoolService) Search(ctx context.Context, query string) ([]SchoolSearchResult, error) {
	if strings.TrimSpace(query) == "" {
		return nil, nil
	}

	rows, err := s.querier.SearchSchools(ctx, sql.NullString{
		String: query,
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	items := make([]SchoolSearchResult, 0, len(rows))
	for _, r := range rows {
		items = append(items, SchoolSearchResult{
			ID:        r.ID,
			Name:      r.Name,
			Level:     r.Level,
			Status:    r.Status,
			Kecamatan: r.Kecamatan,
		})
	}
	return items, nil
}

// List returns paginated schools for admin CRUD.
func (s *SchoolService) List(ctx context.Context, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	rows, err := s.querier.ListSchoolsPaginated(ctx, queries.ListSchoolsPaginatedParams{
		Limit:  int64(perPage),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]SchoolItem, 0, len(rows))
	for _, r := range rows {
		items = append(items, SchoolItem{
			ID:        r.ID,
			Name:      r.Name,
			Level:     r.Level,
			Status:    r.Status,
			Kecamatan: r.Kecamatan,
			IsActive:  r.IsActive == 1,
		})
	}

	total, err := s.querier.CountSchools(ctx)
	if err != nil {
		return nil, err
	}

	return BuildPagination(items, page, perPage, total), nil
}

// Get returns a single school by ID.
func (s *SchoolService) Get(ctx context.Context, id int64) (*SchoolItem, error) {
	r, err := s.querier.GetSchoolByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("sekolah tidak ditemukan")
		}
		return nil, err
	}

	return &SchoolItem{
		ID:        r.ID,
		Name:      r.Name,
		Level:     r.Level,
		Status:    r.Status,
		Kecamatan: r.Kecamatan,
		IsActive:  r.IsActive == 1,
	}, nil
}

// Create adds a new school.
func (s *SchoolService) Create(ctx context.Context, input SchoolInput) (*SchoolItem, error) {
	if strings.TrimSpace(input.Name) == "" {
		return nil, errors.New("nama sekolah wajib diisi")
	}
	if strings.TrimSpace(input.Level) == "" {
		return nil, errors.New("jenjang wajib diisi")
	}
	if strings.TrimSpace(input.Status) == "" {
		return nil, errors.New("status wajib diisi")
	}
	if strings.TrimSpace(input.Kecamatan) == "" {
		return nil, errors.New("kecamatan wajib diisi")
	}

	r, err := s.querier.CreateSchool(ctx, queries.CreateSchoolParams{
		Name:      strings.TrimSpace(input.Name),
		Level:     input.Level,
		Status:    input.Status,
		Kecamatan: strings.TrimSpace(input.Kecamatan),
	})
	if err != nil {
		return nil, err
	}

	return &SchoolItem{
		ID:        r.ID,
		Name:      r.Name,
		Level:     r.Level,
		Status:    r.Status,
		Kecamatan: r.Kecamatan,
		IsActive:  r.IsActive == 1,
	}, nil
}

// Update modifies an existing school.
func (s *SchoolService) Update(ctx context.Context, id int64, input SchoolInput) error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("nama sekolah wajib diisi")
	}
	if strings.TrimSpace(input.Level) == "" {
		return errors.New("jenjang wajib diisi")
	}
	if strings.TrimSpace(input.Status) == "" {
		return errors.New("status wajib diisi")
	}
	if strings.TrimSpace(input.Kecamatan) == "" {
		return errors.New("kecamatan wajib diisi")
	}

	return s.querier.UpdateSchool(ctx, queries.UpdateSchoolParams{
		Name:      strings.TrimSpace(input.Name),
		Level:     input.Level,
		Status:    input.Status,
		Kecamatan: strings.TrimSpace(input.Kecamatan),
		ID:        id,
	})
}

// Delete removes a school.
func (s *SchoolService) Delete(ctx context.Context, id int64) error {
	return s.querier.DeleteSchool(ctx, id)
}
