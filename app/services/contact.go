package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// ContactService handles CRUD for renjana_contacts.
type ContactService struct {
	querier *queries.Querier
}

func NewContactService(querier *queries.Querier) *ContactService {
	return &ContactService{querier: querier}
}

// ---------------------------------------------------------------------------
// DTOs — district_id = NULL berarti tingkat kabupaten (Koordinator)
// ---------------------------------------------------------------------------

type ContactItem struct {
	ID           int64     `json:"id"`
	DistrictID   int64     `json:"district_id"`
	DistrictName string    `json:"district_name"`
	Name         string    `json:"name"`
	Role         string    `json:"role"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

type ContactDetail struct {
	ID           int64     `json:"id"`
	DistrictID   int64     `json:"district_id"`
	DistrictName string    `json:"district_name"`
	Name         string    `json:"name"`
	Role         string    `json:"role"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateContactRequest struct {
	DistrictID int64  `json:"district_id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	IsActive   bool   `json:"is_active"`
}

type UpdateContactRequest struct {
	DistrictID int64  `json:"district_id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	IsActive   bool   `json:"is_active"`
}

var ErrContactNotFound = errors.New("contact not found")

// ListAll returns all contacts (for display pages).
func (s *ContactService) ListAll(ctx context.Context) ([]ContactItem, error) {
	rows, err := s.querier.ListContactsByDistrict(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]ContactItem, 0, len(rows))
	for _, r := range rows {
		did := int64(0)
		if r.DistrictID.Valid {
			did = r.DistrictID.Int64
		}
		dname := ""
		if r.DistrictName.Valid {
			dname = r.DistrictName.String
		}
		role := ""
		if r.Role != "" {
			role = r.Role
		}
		phone := ""
		if r.Phone.Valid {
			phone = r.Phone.String
		}
		email := ""
		if r.Email.Valid {
			email = r.Email.String
		}
		out = append(out, ContactItem{
			ID:           r.ID,
			DistrictID:   did,
			DistrictName: dname,
			Name:         r.Name,
			Role:         role,
			Phone:        phone,
			Email:        email,
			IsActive:     r.IsActive,
			CreatedAt:    r.CreatedAt,
		})
	}
	return out, nil
}

// List returns paginated contacts (for admin CRUD).
func (s *ContactService) List(ctx context.Context, search string, districtID int64, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	var searchArg interface{}
	var districtArg interface{}
	if strings.TrimSpace(search) == "" {
		searchArg = ""
	} else {
		searchArg = search
	}
	if districtID == 0 {
		districtArg = int64(0)
	} else {
		districtArg = districtID
	}

	rows, err := s.querier.ListContactsPaginated(ctx, queries.ListContactsPaginatedParams{
		Column1: searchArg,
		Column2: districtArg,
		Limit:   int64(perPage),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]ContactItem, 0, len(rows))
	for _, r := range rows {
		did := int64(0)
		if r.DistrictID.Valid {
			did = r.DistrictID.Int64
		}
		dname := ""
		if r.DistrictName.Valid {
			dname = r.DistrictName.String
		}
		role := ""
		if r.Role != "" {
			role = r.Role
		}
		phone := ""
		if r.Phone.Valid {
			phone = r.Phone.String
		}
		email := ""
		if r.Email.Valid {
			email = r.Email.String
		}
		items = append(items, ContactItem{
			ID:           r.ID,
			DistrictID:   did,
			DistrictName: dname,
			Name:         r.Name,
			Role:         role,
			Phone:        phone,
			Email:        email,
			IsActive:     r.IsActive,
			CreatedAt:    r.CreatedAt,
		})
	}

	total, err := s.querier.CountContactsFiltered(ctx, queries.CountContactsFilteredParams{
		Column1: searchArg,
		Column2: districtArg,
	})
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}

// Get returns a single contact.
func (s *ContactService) Get(ctx context.Context, id int64) (*ContactDetail, error) {
	r, err := s.querier.GetContactByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	role := ""
	if r.Role != "" {
		role = r.Role
	}
	phone := ""
	if r.Phone.Valid {
		phone = r.Phone.String
	}
	email := ""
	if r.Email.Valid {
		email = r.Email.String
	}
	did := int64(0)
	if r.DistrictID.Valid {
		did = r.DistrictID.Int64
	}
	dname := ""
	if r.DistrictName.Valid {
		dname = r.DistrictName.String
	}
	return &ContactDetail{
		ID:           r.ID,
		DistrictID:   did,
		DistrictName: dname,
		Name:         r.Name,
		Role:         role,
		Phone:        phone,
		Email:        email,
		IsActive:     r.IsActive,
		CreatedAt:    r.CreatedAt,
	}, nil
}

// Create inserts a new contact.
func (s *ContactService) Create(ctx context.Context, req CreateContactRequest) (*ContactDetail, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("nama wajib diisi")
	}
	// Koordinator (kabupaten) boleh tanpa district_id
	if req.DistrictID == 0 && req.Role != "Koordinator" {
		return nil, errors.New("kecamatan wajib dipilih")
	}
	if req.Role == "" {
		req.Role = "Fasilitator"
	}

	var districtArg sql.NullInt64
	if req.Role == "Koordinator" && req.DistrictID == 0 {
		districtArg = sql.NullInt64{Valid: false}
	} else {
		districtArg = sql.NullInt64{Int64: req.DistrictID, Valid: true}
	}

	id, err := s.querier.CreateContact(ctx, queries.CreateContactParams{
		DistrictID: districtArg,
		Name:       req.Name,
		Role:       req.Role,
		Phone:      sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Email:      sql.NullString{String: req.Email, Valid: req.Email != ""},
		IsActive:   req.IsActive,
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

// Update modifies an existing contact.
func (s *ContactService) Update(ctx context.Context, id int64, req UpdateContactRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("nama wajib diisi")
	}
	// Koordinator (kabupaten) boleh tanpa district_id
	if req.DistrictID == 0 && req.Role != "Koordinator" {
		return errors.New("kecamatan wajib dipilih")
	}
	if req.Role == "" {
		req.Role = "Fasilitator"
	}

	var districtArg sql.NullInt64
	if req.Role == "Koordinator" && req.DistrictID == 0 {
		districtArg = sql.NullInt64{Valid: false}
	} else {
		districtArg = sql.NullInt64{Int64: req.DistrictID, Valid: true}
	}

	rows, err := s.querier.UpdateContact(ctx, queries.UpdateContactParams{
		DistrictID: districtArg,
		Name:       req.Name,
		Role:       req.Role,
		Phone:      sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Email:      sql.NullString{String: req.Email, Valid: req.Email != ""},
		IsActive:   req.IsActive,
		ID:         id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrContactNotFound
	}
	return nil
}

// Delete removes a contact.
func (s *ContactService) Delete(ctx context.Context, id int64) error {
	rows, err := s.querier.DeleteContact(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrContactNotFound
	}
	return nil
}
