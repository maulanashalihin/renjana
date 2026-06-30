package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// ActivityService handles CRUD for renjana_activities.
type ActivityService struct {
	querier *queries.Querier
}

func NewActivityService(querier *queries.Querier) *ActivityService {
	return &ActivityService{querier: querier}
}

// ---------------------------------------------------------------------------
// DTOs
// ---------------------------------------------------------------------------

// ActivityListItem — one row in the CRUD table.
type ActivityListItem struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	TypeID       int64     `json:"type_id"`
	TypeName     string    `json:"type_name"`
	TypeColor    string    `json:"type_color"`
	TypeIcon     string    `json:"type_icon"`
	DistrictID   int64     `json:"district_id"`
	DistrictName string    `json:"district_name"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Date         time.Time `json:"date"`
	Time         string    `json:"time"`
	Status       string    `json:"status"`
}

// ActivityDetail — full record (for show/edit).
type ActivityDetail struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	TypeID       int64     `json:"type_id"`
	TypeName     string    `json:"type_name"`
	TypeColor    string    `json:"type_color"`
	TypeIcon     string    `json:"type_icon"`
	DistrictID   int64     `json:"district_id"`
	DistrictName string    `json:"district_name"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Date         time.Time `json:"date"`
	Time         string    `json:"time"`
	Status       string    `json:"status"`
}

// ActivityStats — banner stats for the list page.
type ActivityStats struct {
	Total     int64 `json:"total"`
	Upcoming  int64 `json:"upcoming"`
	Ongoing   int64 `json:"ongoing"`
	Completed int64 `json:"completed"`
}

// CreateActivityRequest — input for create.
type CreateActivityRequest struct {
	Title       string `json:"title"`
	TypeID      int64  `json:"type_id"`
	DistrictID  int64  `json:"district_id"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"` // YYYY-MM-DD
	Time        string `json:"time"`
	Status      string `json:"status"`
}

// UpdateActivityRequest — input for update.
type UpdateActivityRequest struct {
	Title       string `json:"title"`
	TypeID      int64  `json:"type_id"`
	DistrictID  int64  `json:"district_id"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Status      string `json:"status"`
}

// ---------------------------------------------------------------------------
// Errors
// ---------------------------------------------------------------------------

var (
	ErrActivityNotFound = errors.New("activity not found")
)

// ---------------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------------

// List returns a paginated, filtered list of activities.
func (s *ActivityService) List(
	ctx context.Context,
	search string,
	typeID int64,
	status string,
	page, perPage int,
) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var searchArg, statusArg interface{}
	var typeArg interface{}
	if strings.TrimSpace(search) == "" {
		searchArg = ""
	} else {
		searchArg = search
	}
	if typeID == 0 {
		typeArg = int64(0)
	} else {
		typeArg = int64(typeID)
	}
	if strings.TrimSpace(status) == "" {
		statusArg = ""
	} else {
		statusArg = status
	}

	rows, err := s.querier.ListActivitiesPaginated(ctx, queries.ListActivitiesPaginatedParams{
		Column1: searchArg,
		Column2: typeArg,
		Column3: statusArg,
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]ActivityListItem, 0, len(rows))
	for _, r := range rows {
		desc := ""
		if r.Description.Valid {
			desc = r.Description.String
		}
		items = append(items, ActivityListItem{
			ID:           r.ID,
			Title:        r.Title,
			TypeID:       r.TypeID,
			TypeName:     r.TypeName,
			TypeColor:    r.TypeColor,
			TypeIcon:     r.TypeIcon,
			DistrictID:   r.DistrictID,
			DistrictName: r.DistrictName,
			Description:  desc,
			Location:     r.Location,
			Date:         r.Date,
			Time:         r.Time,
			Status:       r.Status,
		})
	}

	total, err := s.querier.CountActivitiesFiltered(ctx, queries.CountActivitiesFilteredParams{
		Column1: searchArg,
		Column2: typeArg,
		Column3: statusArg,
	})
	if err != nil {
		return nil, err
	}

	return BuildPagination(items, page, perPage, total), nil
}

// ListScoped returns paginated activities filtered by district_id.
// Pass districtID = 0 to get all (admin use).
// Otherwise only activities in the specified district are returned.
func (s *ActivityService) ListScoped(
	ctx context.Context,
	search string,
	typeID int64,
	status string,
	districtID int64,
	page, perPage int,
) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var searchArg, statusArg interface{}
	var typeArg interface{}
	if strings.TrimSpace(search) == "" {
		searchArg = ""
	} else {
		searchArg = search
	}
	if typeID == 0 {
		typeArg = int64(0)
	} else {
		typeArg = typeID
	}
	if strings.TrimSpace(status) == "" {
		statusArg = ""
	} else {
		statusArg = status
	}

	rows, err := s.querier.ListActivitiesPaginatedScoped(ctx, queries.ListActivitiesPaginatedScopedParams{
		Column1: searchArg,
		Column2: typeArg,
		Column3: statusArg,
		Column4: int64(districtID),
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]ActivityListItem, 0, len(rows))
	for _, r := range rows {
		desc := ""
		if r.Description.Valid {
			desc = r.Description.String
		}
		items = append(items, ActivityListItem{
			ID:           r.ID,
			Title:        r.Title,
			TypeID:       r.TypeID,
			TypeName:     r.TypeName,
			TypeColor:    r.TypeColor,
			TypeIcon:     r.TypeIcon,
			DistrictID:   r.DistrictID,
			DistrictName: r.DistrictName,
			Description:  desc,
			Location:     r.Location,
			Date:         r.Date,
			Time:         r.Time,
			Status:       r.Status,
		})
	}

	total, err := s.querier.CountActivitiesFilteredScoped(ctx, queries.CountActivitiesFilteredScopedParams{
		Column1: searchArg,
		Column2: typeArg,
		Column3: statusArg,
		Column4: int64(districtID),
	})
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}

// Get returns a single activity detail.
func (s *ActivityService) Get(ctx context.Context, id int64) (*ActivityDetail, error) {
	r, err := s.querier.GetActivityByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrActivityNotFound
		}
		return nil, err
	}

	desc := ""
	if r.Description.Valid {
		desc = r.Description.String
	}

	return &ActivityDetail{
		ID:           r.ID,
		Title:        r.Title,
		TypeID:       r.TypeID,
		TypeName:     r.TypeName,
		TypeColor:    r.TypeColor,
		TypeIcon:     r.TypeIcon,
		DistrictID:   r.DistrictID,
		DistrictName: r.DistrictName,
		Description:  desc,
		Location:     r.Location,
		Date:         r.Date,
		Time:         r.Time,
		Status:       r.Status,
	}, nil
}

// Create inserts a new activity.
func (s *ActivityService) Create(ctx context.Context, req CreateActivityRequest) (*ActivityDetail, error) {
	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("judul wajib diisi")
	}
	if req.TypeID == 0 {
		return nil, errors.New("jenis kegiatan wajib dipilih")
	}
	if req.DistrictID == 0 {
		return nil, errors.New("kecamatan wajib dipilih")
	}
	if strings.TrimSpace(req.Location) == "" {
		return nil, errors.New("lokasi wajib diisi")
	}
	if req.Date == "" {
		return nil, errors.New("tanggal wajib diisi")
	}
	if req.Time == "" {
		return nil, errors.New("waktu wajib diisi")
	}
	if req.Status == "" {
		req.Status = "akan_datang"
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("tanggal tidak valid")
	}

	id, err := s.querier.CreateActivity(ctx, queries.CreateActivityParams{
		Title:       req.Title,
		TypeID:      req.TypeID,
		DistrictID:  req.DistrictID,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Location:    req.Location,
		Date:        date,
		Time:        req.Time,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

// Update modifies an existing activity.
func (s *ActivityService) Update(ctx context.Context, id int64, req UpdateActivityRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("judul wajib diisi")
	}
	if req.TypeID == 0 {
		return errors.New("jenis kegiatan wajib dipilih")
	}
	if req.DistrictID == 0 {
		return errors.New("kecamatan wajib dipilih")
	}
	if strings.TrimSpace(req.Location) == "" {
		return errors.New("lokasi wajib diisi")
	}
	if req.Date == "" {
		return errors.New("tanggal wajib diisi")
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return errors.New("tanggal tidak valid")
	}

	rows, err := s.querier.UpdateActivity(ctx, queries.UpdateActivityParams{
		Title:       req.Title,
		TypeID:      req.TypeID,
		DistrictID:  req.DistrictID,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Location:    req.Location,
		Date:        date,
		Time:        req.Time,
		Status:      req.Status,
		ID:          id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrActivityNotFound
	}
	return nil
}

// Delete removes an activity.
func (s *ActivityService) Delete(ctx context.Context, id int64) error {
	rows, err := s.querier.DeleteActivity(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrActivityNotFound
	}
	return nil
}

// GetStats returns aggregate stats for the list page banner.
func (s *ActivityService) GetStats(ctx context.Context) (*ActivityStats, error) {
	r, err := s.querier.GetActivityStats(ctx)
	if err != nil {
		return nil, err
	}
	return &ActivityStats{
		Total:     r.Total,
		Upcoming:  nullFloatToInt(r.Upcoming),
		Ongoing:   nullFloatToInt(r.Ongoing),
		Completed: nullFloatToInt(r.Completed),
	}, nil
}
