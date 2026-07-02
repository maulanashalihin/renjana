package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// ---------------------------------------------------------------------------
// Survey (Survey Pelayanan Publik)
// ---------------------------------------------------------------------------

type SurveyItem struct {
	ID              int64     `json:"id"`
	RespondentName  string    `json:"respondent_name,omitempty"`
	RespondentEmail string    `json:"respondent_email,omitempty"`
	ServiceType     string    `json:"service_type"`
	Rating          int64     `json:"rating"`
	Feedback        string    `json:"feedback,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

type SurveyStats struct {
	Total         int64   `json:"total"`
	AverageRating float64 `json:"average_rating"`
	Rating5       int64   `json:"rating_5"`
	Rating4       int64   `json:"rating_4"`
	Rating3       int64   `json:"rating_3"`
	Rating2       int64   `json:"rating_2"`
	Rating1       int64   `json:"rating_1"`
}

type SurveyServiceStatsByService struct {
	ServiceType   string  `json:"service_type"`
	Total         int64   `json:"total"`
	AverageRating float64 `json:"average_rating"`
}

type SurveyService struct {
	querier *queries.Querier
}

func NewSurveyService(querier *queries.Querier) *SurveyService {
	return &SurveyService{querier: querier}
}

func (s *SurveyService) List(ctx context.Context, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	rows, err := s.querier.ListSurveysPaginated(ctx, queries.ListSurveysPaginatedParams{
		Limit:  int64(perPage),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]SurveyItem, 0, len(rows))
	for _, r := range rows {
		items = append(items, *surveyFromRow(&r))
	}

	total, err := s.querier.CountSurveys(ctx)
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}

func (s *SurveyService) Create(ctx context.Context, respondentName, respondentEmail, serviceType string, rating int64, feedback string) (*SurveyItem, error) {
	r, err := s.querier.CreateSurvey(ctx, queries.CreateSurveyParams{
		RespondentName:  sql.NullString{String: respondentName, Valid: respondentName != ""},
		RespondentEmail: sql.NullString{String: respondentEmail, Valid: respondentEmail != ""},
		ServiceType:     serviceType,
		Rating:          rating,
		Feedback:        sql.NullString{String: feedback, Valid: feedback != ""},
	})
	if err != nil {
		return nil, err
	}
	return surveyFromRow(&r), nil
}

func (s *SurveyService) GetStats(ctx context.Context) (*SurveyStats, error) {
	r, err := s.querier.GetSurveyStats(ctx)
	if err != nil {
		return nil, err
	}
	return &SurveyStats{
		Total:         r.Total,
		AverageRating: r.AverageRating,
		Rating5:       int64(r.Rating5.Float64),
		Rating4:       int64(r.Rating4.Float64),
		Rating3:       int64(r.Rating3.Float64),
		Rating2:       int64(r.Rating2.Float64),
		Rating1:       int64(r.Rating1.Float64),
	}, nil
}

func (s *SurveyService) GetStatsByService(ctx context.Context) ([]SurveyServiceStatsByService, error) {
	rows, err := s.querier.GetSurveyStatsByService(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]SurveyServiceStatsByService, 0, len(rows))
	for _, r := range rows {
		items = append(items, SurveyServiceStatsByService{
			ServiceType:   r.ServiceType,
			Total:         r.Total,
			AverageRating: r.AverageRating,
		})
	}
	return items, nil
}

func surveyFromRow(r *queries.RenjanaSurvey) *SurveyItem {
	name := ""
	if r.RespondentName.Valid {
		name = r.RespondentName.String
	}
	email := ""
	if r.RespondentEmail.Valid {
		email = r.RespondentEmail.String
	}
	feedback := ""
	if r.Feedback.Valid {
		feedback = r.Feedback.String
	}
	return &SurveyItem{
		ID:              r.ID,
		RespondentName:  name,
		RespondentEmail: email,
		ServiceType:     r.ServiceType,
		Rating:          r.Rating,
		Feedback:        feedback,
		CreatedAt:       r.CreatedAt,
	}
}
