package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// ---------------------------------------------------------------------------
// Complaint (Pengaduan)
// ---------------------------------------------------------------------------

type ComplaintItem struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone,omitempty"`
	Category    string     `json:"category"`
	Message     string     `json:"message"`
	Status      string     `json:"status"`
	Response    string     `json:"response,omitempty"`
	RespondedBy *int64     `json:"responded_by,omitempty"`
	RespondedAt *time.Time `json:"responded_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

type ComplaintStats struct {
	Total     int64 `json:"total"`
	Pending   int64 `json:"pending"`
	Processed int64 `json:"processed"`
	Resolved  int64 `json:"resolved"`
}

type ComplaintService struct {
	querier *queries.Querier
}

func NewComplaintService(querier *queries.Querier) *ComplaintService {
	return &ComplaintService{querier: querier}
}

func (s *ComplaintService) List(ctx context.Context, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	rows, err := s.querier.ListComplaintsPaginated(ctx, queries.ListComplaintsPaginatedParams{
		Limit:  int64(perPage),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]ComplaintItem, 0, len(rows))
	for _, r := range rows {
		items = append(items, *complaintFromRow(&r))
	}

	total, err := s.querier.CountComplaints(ctx)
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}

func (s *ComplaintService) Get(ctx context.Context, id int64) (*ComplaintItem, error) {
	r, err := s.querier.GetComplaintByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return complaintFromRow(&r), nil
}

func (s *ComplaintService) Create(ctx context.Context, name, email, phone, category, message string) (*ComplaintItem, error) {
	r, err := s.querier.CreateComplaint(ctx, queries.CreateComplaintParams{
		Name:     name,
		Email:    email,
		Phone:    sql.NullString{String: phone, Valid: phone != ""},
		Category: category,
		Message:  message,
	})
	if err != nil {
		return nil, err
	}
	return complaintFromRow(&r), nil
}

func (s *ComplaintService) UpdateStatus(ctx context.Context, id int64, status, response string, respondedBy int64) (*ComplaintItem, error) {
	r, err := s.querier.UpdateComplaintStatus(ctx, queries.UpdateComplaintStatusParams{
		ID:          id,
		Status:      status,
		Response:    sql.NullString{String: response, Valid: response != ""},
		RespondedBy: sql.NullInt64{Int64: respondedBy, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return complaintFromRow(&r), nil
}

func (s *ComplaintService) GetStats(ctx context.Context) (*ComplaintStats, error) {
	r, err := s.querier.GetComplaintStats(ctx)
	if err != nil {
		return nil, err
	}
	return &ComplaintStats{
		Total:     r.Total,
		Pending:   int64(r.Pending.Float64),
		Processed: int64(r.Processed.Float64),
		Resolved:  int64(r.Resolved.Float64),
	}, nil
}

func (s *ComplaintService) Delete(ctx context.Context, id int64) error {
	return s.querier.DeleteComplaint(ctx, id)
}

func complaintFromRow(r *queries.RenjanaComplaint) *ComplaintItem {
	phone := ""
	if r.Phone.Valid {
		phone = r.Phone.String
	}
	response := ""
	if r.Response.Valid {
		response = r.Response.String
	}
	var respondedBy *int64
	if r.RespondedBy.Valid {
		respondedBy = &r.RespondedBy.Int64
	}
	var respondedAt *time.Time
	if r.RespondedAt.Valid {
		respondedAt = &r.RespondedAt.Time
	}
	return &ComplaintItem{
		ID:          r.ID,
		Name:        r.Name,
		Email:       r.Email,
		Phone:       phone,
		Category:    r.Category,
		Message:     r.Message,
		Status:      r.Status,
		Response:    response,
		RespondedBy: respondedBy,
		RespondedAt: respondedAt,
		CreatedAt:   r.CreatedAt,
	}
}
