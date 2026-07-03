package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// VolunteerService handles CRUD for renjana_volunteers.
type VolunteerService struct {
	querier *queries.Querier
}

func NewVolunteerService(querier *queries.Querier) *VolunteerService {
	return &VolunteerService{querier: querier}
}

// ---------------------------------------------------------------------------
// DTOs
// ---------------------------------------------------------------------------

// VolunteerListItem — one row in the CRUD table.
type VolunteerListItem struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"`
	Name              string    `json:"name"`
	School            string    `json:"school"`
	DistrictID        int64     `json:"district_id"`
	DistrictName      string    `json:"district_name"`
	Status            string    `json:"status"`
	Phone             string    `json:"phone"`
	AvatarURL         string    `json:"avatar_url"`
	ApplicationStatus string    `json:"application_status"`
	JoinedAt          time.Time `json:"joined_at"`
	IsActive          bool      `json:"is_active"`
	CertificateCount  int64     `json:"certificate_count"`
}

// VolunteerDetail — full record (for show/edit).
type VolunteerDetail struct {
	ID                int64          `json:"id"`
	UserID            int64          `json:"user_id"`
	Name              string         `json:"name"`
	School            string         `json:"school"`
	DistrictID        int64          `json:"district_id"`
	DistrictName      string         `json:"district_name"`
	Phone             string         `json:"phone"`
	Status            string         `json:"status"`
	AvatarURL         string         `json:"avatar_url"`
	JoinedAt          time.Time      `json:"joined_at"`
	IsActive          bool           `json:"is_active"`
	ApplicationStatus string         `json:"application_status"`
	ReviewerID        sql.NullInt64  `json:"-"`
	ReviewedAt        sql.NullTime   `json:"-"`
	RejectionReason   sql.NullString `json:"-"`
}

// VolunteerStats — banner stats for the list page.
type VolunteerStats struct {
	Total    int64 `json:"total"`
	Active   int64 `json:"active"`
	Inactive int64 `json:"inactive"`
	Pending  int64 `json:"pending"`
	Rejected int64 `json:"rejected"`
	Schools  int64 `json:"schools"`
}

// CreateVolunteerRequest — input for create.
type CreateVolunteerRequest struct {
	Name              string `json:"name"`
	School            string `json:"school"`
	DistrictID        int64  `json:"district_id"`
	Phone             string `json:"phone"`
	Status            string `json:"status"`
	JoinedAt          string `json:"joined_at"` // YYYY-MM-DD
	ApplicationStatus string `json:"application_status"`
}

// UpdateVolunteerRequest — input for update.
type UpdateVolunteerRequest struct {
	Name       string `json:"name"`
	School     string `json:"school"`
	DistrictID int64  `json:"district_id"`
	Phone      string `json:"phone"`
	Status     string `json:"status"`
	JoinedAt   string `json:"joined_at"`
}

// ---------------------------------------------------------------------------
// Errors
// ---------------------------------------------------------------------------

var (
	ErrVolunteerNotFound = errors.New("volunteer not found")
	ErrInvalidInput      = errors.New("invalid input")
)

// ---------------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------------

// List returns a paginated, filtered list of volunteers.
func (s *VolunteerService) List(
	ctx context.Context,
	search string,
	districtID int64,
	status, applicationStatus string,
	page, perPage int,
) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var searchArg, statusArg, appStatusArg interface{}
	var districtArg interface{}
	if strings.TrimSpace(search) == "" {
		searchArg = ""
	} else {
		searchArg = search
	}
	if districtID == 0 {
		districtArg = int64(0)
	} else {
		districtArg = int64(districtID)
	}
	if strings.TrimSpace(status) == "" {
		statusArg = ""
	} else {
		statusArg = status
	}
	if strings.TrimSpace(applicationStatus) == "" {
		appStatusArg = ""
	} else {
		appStatusArg = applicationStatus
	}

	rows, err := s.querier.ListVolunteersPaginated(ctx, queries.ListVolunteersPaginatedParams{
		Column1: searchArg,
		Column2: districtArg,
		Column3: statusArg,
		Column4: appStatusArg,
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]VolunteerListItem, 0, len(rows))
	for _, r := range rows {
		phone := ""
		if r.Phone.Valid {
			phone = r.Phone.String
		}
		avatar := ""
		if r.AvatarUrl.Valid {
			avatar = r.AvatarUrl.String
		}

		// Count certificates for this volunteer (shared ID: volunteer.id = user.id)
		var certCount int64
		cnt, err := s.querier.CountCertificatesByUser(ctx, r.ID)
		if err == nil {
			certCount = cnt
		}

		items = append(items, VolunteerListItem{
			ID:                r.ID,
			UserID:            r.ID,
			Name:              r.Name,
			School:            r.School,
			DistrictID:        r.DistrictID,
			DistrictName:      r.DistrictName,
			Status:            r.Status,
			Phone:             phone,
			AvatarURL:         avatar,
			ApplicationStatus: r.ApplicationStatus,
			JoinedAt:          r.JoinedAt,
			IsActive:          r.IsActive,
			CertificateCount:  certCount,
		})
	}

	total, err := s.querier.CountVolunteersFiltered(ctx, queries.CountVolunteersFilteredParams{
		Column1: searchArg,
		Column2: districtArg,
		Column3: statusArg,
		Column4: appStatusArg,
	})
	if err != nil {
		return nil, err
	}

	return BuildPagination(items, page, perPage, total), nil
}

// Get returns a single volunteer detail.
func (s *VolunteerService) Get(ctx context.Context, id int64) (*VolunteerDetail, error) {
	r, err := s.querier.GetVolunteerByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrVolunteerNotFound
		}
		return nil, err
	}

	phone := ""
	if r.Phone.Valid {
		phone = r.Phone.String
	}
	avatar := ""
	if r.AvatarUrl.Valid {
		avatar = r.AvatarUrl.String
	}
	districtName := ""
	if r.DistrictName.Valid {
		districtName = r.DistrictName.String
	}

	return &VolunteerDetail{
		ID:                r.ID,
		UserID:            r.ID,
		Name:              r.Name,
		School:            r.School,
		DistrictID:        r.DistrictID,
		DistrictName:      districtName,
		Phone:             phone,
		Status:            r.Status,
		AvatarURL:         avatar,
		JoinedAt:          r.JoinedAt,
		IsActive:          r.IsActive,
		ApplicationStatus: r.ApplicationStatus,
		ReviewerID:        r.ReviewerID,
		ReviewedAt:        r.ReviewedAt,
		RejectionReason:   r.RejectionReason,
	}, nil
}

// Create inserts a new volunteer. Validates required fields.
func (s *VolunteerService) Create(ctx context.Context, req CreateVolunteerRequest) (*VolunteerDetail, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("nama wajib diisi")
	}
	if strings.TrimSpace(req.School) == "" {
		return nil, errors.New("sekolah wajib diisi")
	}
	if req.DistrictID == 0 {
		return nil, errors.New("kecamatan wajib dipilih")
	}
	if req.Status == "" {
		req.Status = "aktif"
	}
	if req.ApplicationStatus == "" {
		req.ApplicationStatus = "approved"
	}
	if req.JoinedAt == "" {
		req.JoinedAt = time.Now().Format("2006-01-02")
	}

	joinedAt, err := time.Parse("2006-01-02", req.JoinedAt)
	if err != nil {
		return nil, errors.New("tanggal gabung tidak valid")
	}

	id, err := s.querier.CreateVolunteer(ctx, queries.CreateVolunteerParams{
		Name:              req.Name,
		School:            req.School,
		DistrictID:        req.DistrictID,
		Phone:             sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Status:            req.Status,
		JoinedAt:          joinedAt,
		ApplicationStatus: req.ApplicationStatus,
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

// Update modifies an existing volunteer.
func (s *VolunteerService) Update(ctx context.Context, id int64, req UpdateVolunteerRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("nama wajib diisi")
	}
	if strings.TrimSpace(req.School) == "" {
		return errors.New("sekolah wajib diisi")
	}
	if req.DistrictID == 0 {
		return errors.New("kecamatan wajib dipilih")
	}
	if req.Status == "" {
		req.Status = "aktif"
	}
	joinedAt, err := time.Parse("2006-01-02", req.JoinedAt)
	if err != nil {
		return errors.New("tanggal gabung tidak valid")
	}

	rows, err := s.querier.UpdateVolunteer(ctx, queries.UpdateVolunteerParams{
		Name:       req.Name,
		School:     req.School,
		DistrictID: req.DistrictID,
		Phone:      sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Status:     req.Status,
		JoinedAt:   joinedAt,
		ID:         id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrVolunteerNotFound
	}
	return nil
}

// Delete removes a volunteer (hard delete).
func (s *VolunteerService) Delete(ctx context.Context, id int64) error {
	rows, err := s.querier.DeleteVolunteer(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrVolunteerNotFound
	}
	return nil
}

// ApproveApplication accepts a pending volunteer.
func (s *VolunteerService) ApproveApplication(ctx context.Context, id, reviewerID int64) error {
	rows, err := s.querier.ApproveVolunteerApplication(ctx, queries.ApproveVolunteerApplicationParams{
		ReviewerID: sql.NullInt64{Int64: reviewerID, Valid: true},
		ID:         id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrVolunteerNotFound
	}
	return nil
}

// RejectApplication rejects a pending volunteer with reason.
func (s *VolunteerService) RejectApplication(ctx context.Context, id, reviewerID int64, reason string) error {
	rows, err := s.querier.RejectVolunteerApplication(ctx, queries.RejectVolunteerApplicationParams{
		ReviewerID:      sql.NullInt64{Int64: reviewerID, Valid: true},
		RejectionReason: sql.NullString{String: reason, Valid: reason != ""},
		ID:              id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrVolunteerNotFound
	}
	return nil
}

// GetPendingApplications returns pending applications.
func (s *VolunteerService) GetPendingApplications(ctx context.Context, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	rows, err := s.querier.ListPendingApplications(ctx, queries.ListPendingApplicationsParams{
		Limit:  int64(perPage),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]VolunteerListItem, 0, len(rows))
	for _, r := range rows {
		phone := ""
		if r.Phone.Valid {
			phone = r.Phone.String
		}
		avatar := ""
		if r.AvatarUrl.Valid {
			avatar = r.AvatarUrl.String
		}
		items = append(items, VolunteerListItem{
			ID:                r.ID,
			Name:              r.Name,
			School:            r.School,
			DistrictID:        r.DistrictID,
			DistrictName:      r.DistrictName,
			Status:            "nonaktif",
			Phone:             phone,
			AvatarURL:         avatar,
			ApplicationStatus: r.ApplicationStatus,
			JoinedAt:          r.JoinedAt,
			IsActive:          false,
		})
	}

	total, err := s.querier.CountPendingApplications(ctx)
	if err != nil {
		return nil, err
	}

	return BuildPagination(items, page, perPage, total), nil
}

// GetStats returns aggregate stats for the dashboard banner.
func (s *VolunteerService) GetStats(ctx context.Context) (*VolunteerStats, error) {
	r, err := s.querier.GetVolunteerStats(ctx)
	if err != nil {
		return nil, err
	}
	return &VolunteerStats{
		Total:    r.Total,
		Active:   nullFloatToInt(r.Active),
		Inactive: nullFloatToInt(r.Inactive),
		Pending:  nullFloatToInt(r.Pending),
		Rejected: nullFloatToInt(r.Rejected),
		Schools:  r.Schools,
	}, nil
}

// GetStatsByDistrict returns stats for a single district (for koordinator view).
func (s *VolunteerService) GetStatsByDistrict(ctx context.Context, districtID int64) (*VolunteerStats, error) {
	r, err := s.querier.GetVolunteerStatsByDistrict(ctx, districtID)
	if err != nil {
		return nil, err
	}
	return &VolunteerStats{
		Total:    r.Total,
		Active:   nullFloatToInt(r.Active),
		Inactive: nullFloatToInt(r.Inactive),
		Pending:  nullFloatToInt(r.Pending),
		Rejected: nullFloatToInt(r.Rejected),
		Schools:  r.Schools,
	}, nil
}

// GetPendingApplicationsByDistrict returns pending volunteer applications for a specific district.
func (s *VolunteerService) GetPendingApplicationsByDistrict(ctx context.Context, districtID int64, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	rows, err := s.querier.ListPendingApplicationsByDistrict(ctx, queries.ListPendingApplicationsByDistrictParams{
		DistrictID: districtID,
		Limit:      int64(perPage),
		Offset:     int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]VolunteerListItem, 0, len(rows))
	for _, r := range rows {
		phone := ""
		if r.Phone.Valid {
			phone = r.Phone.String
		}
		avatar := ""
		if r.AvatarUrl.Valid {
			avatar = r.AvatarUrl.String
		}
		items = append(items, VolunteerListItem{
			ID:                r.ID,
			Name:              r.Name,
			School:            r.School,
			DistrictID:        r.DistrictID,
			DistrictName:      r.DistrictName,
			Status:            "nonaktif",
			Phone:             phone,
			AvatarURL:         avatar,
			ApplicationStatus: r.ApplicationStatus,
			JoinedAt:          r.JoinedAt,
			IsActive:          false,
		})
	}

	total, err := s.querier.CountPendingApplicationsByDistrict(ctx, districtID)
	if err != nil {
		return nil, err
	}

	return BuildPagination(items, page, perPage, total), nil
}

func nullFloatToInt(n sql.NullFloat64) int64 {
	if !n.Valid {
		return 0
	}
	return int64(n.Float64)
}

// ---------------------------------------------------------------------------
// User ↔ Volunteer linkage (1:1 onboarding flow)
// ---------------------------------------------------------------------------

// OnboardingRequest — input from onboarding form.
type OnboardingRequest struct {
	School     string `json:"school"`
	DistrictID int64  `json:"district_id"`
	Phone      string `json:"phone"`
	AvatarURL  string `json:"avatar_url"`
}

// GetByUserID returns the volunteer record linked to a user.
// Returns (nil, nil) if no record exists — that's not an error.
func (s *VolunteerService) GetByUserID(ctx context.Context, userID int64) (*VolunteerDetail, error) {
	r, err := s.querier.GetVolunteerByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	phone := ""
	if r.Phone.Valid {
		phone = r.Phone.String
	}
	avatar := ""
	if r.AvatarUrl.Valid {
		avatar = r.AvatarUrl.String
	}
	districtName := ""
	if r.DistrictName.Valid {
		districtName = r.DistrictName.String
	}

	return &VolunteerDetail{
		ID:                r.ID,
		UserID:            r.ID,
		Name:              r.Name,
		School:            r.School,
		DistrictID:        r.DistrictID,
		DistrictName:      districtName,
		Phone:             phone,
		Status:            r.Status,
		AvatarURL:         avatar,
		JoinedAt:          r.JoinedAt,
		IsActive:          r.IsActive,
		ApplicationStatus: r.ApplicationStatus,
		ReviewerID:        r.ReviewerID,
		ReviewedAt:        r.ReviewedAt,
		RejectionReason:   r.RejectionReason,
	}, nil
}

// CreateForUser creates a volunteer record linked to a user (1:1).
// Called during onboarding after a user registers.
// Sets status='aktif' and application_status='approved' — user is auto-approved.
func (s *VolunteerService) CreateForUser(ctx context.Context, userID int64, userName, userAvatar string, req OnboardingRequest) (*VolunteerDetail, error) {
	if strings.TrimSpace(req.School) == "" {
		return nil, errors.New("sekolah wajib diisi")
	}
	if req.DistrictID == 0 {
		return nil, errors.New("kecamatan wajib dipilih")
	}

	id, err := s.querier.CreateVolunteerForUserDirect(ctx, userID, userName, req.School, req.DistrictID, req.Phone, time.Now(), userAvatar)
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}
