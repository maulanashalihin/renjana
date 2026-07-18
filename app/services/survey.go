package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// ---------------------------------------------------------------------------
// Survey SKM (Survey Kepuasan Masyarakat)
// ---------------------------------------------------------------------------

type SKMQuestionLabel struct {
	Number  int         `json:"number"`
	Label   string      `json:"label"`
	Options []SKMOption `json:"options"`
}

type SKMOption struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

// Standard 9 SKM questions with 4-point scale
var SKMQuestions = []SKMQuestionLabel{
	{Number: 1, Label: "Bagaimana pendapat saudara tentang Kesesuaian Persyaratan Pelayanan dengan jenis pelayanannya", Options: []SKMOption{{1, "Tidak sesuai"}, {2, "Kurang sesuai"}, {3, "Sesuai"}, {4, "Sangat sesuai"}}},
	{Number: 2, Label: "Bagaimana pemahaman saudara tentang Kemudahan Prosedur Pelayanan di unit ini", Options: []SKMOption{{1, "Tidak mudah"}, {2, "Kurang mudah"}, {3, "Mudah"}, {4, "Sangat mudah"}}},
	{Number: 3, Label: "Bagaimana pendapat saudara tentang Kecepatan Pelayanan di unit ini", Options: []SKMOption{{1, "Tidak tepat waktu"}, {2, "Kadang-kadang tepat waktu"}, {3, "Banyak tepat waktu"}, {4, "Selalu tepat waktu"}}},
	{Number: 4, Label: "Bagaimana pendapat saudara tentang Kesesuaian Antara Biaya yang dibayarkan dengan biaya yang telah ditetapkan", Options: []SKMOption{{1, "Selalu tidak sesuai"}, {2, "Kadang-kadang sesuai"}, {3, "Banyak sesuainga"}, {4, "Selalu sesuai"}}},
	{Number: 5, Label: "Bagaimana pendapat saudara tentang Kesesuaian Hasil Pelayanan yang di berikan dan diterima dengan ketentuan yang telah di tetapkan", Options: []SKMOption{{1, "Tidak sesuai"}, {2, "Kurang sesuai"}, {3, "Sesuai"}, {4, "Sangat sesuai"}}},
	{Number: 6, Label: "Bagaimana pendapat saudara tentang Kemampuan Petugas dalam memberikan pelayanan", Options: []SKMOption{{1, "Tidak mampu"}, {2, "Kurang mampu"}, {3, "Mampu"}, {4, "Sangat mampu"}}},
	{Number: 7, Label: "Bagaimana pendapat saudara tentang Sikap Petugas dalam memberikan pelayanan", Options: []SKMOption{{1, "Tidak baik"}, {2, "Kurang baik"}, {3, "Baik"}, {4, "Sangat baik"}}},
	{Number: 8, Label: "Bagaimana pendapat saudara tentang penanganan pengaduan, saran & masukan terkait pelayanan yang diberikan", Options: []SKMOption{{1, "Tidak baik"}, {2, "Kurang baik"}, {3, "Baik"}, {4, "Sangat baik"}}},
	{Number: 9, Label: "Bagaimana pendapat saudara tentang sarana dan prasarana dalam pelayanan yang digunakan", Options: []SKMOption{{1, "Tidak sesuai"}, {2, "Kurang sesuai"}, {3, "Sesuai"}, {4, "Sangat sesuai"}}},
}

type SurveySKMItem struct {
	ID         int64     `json:"id"`
	Age        int64     `json:"age"`
	Gender     string    `json:"gender"`
	Education  string    `json:"education"`
	Occupation string    `json:"occupation"`
	Year       int64     `json:"year"`
	Q1         int64     `json:"q1"`
	Q2         int64     `json:"q2"`
	Q3         int64     `json:"q3"`
	Q4         int64     `json:"q4"`
	Q5         int64     `json:"q5"`
	Q6         int64     `json:"q6"`
	Q7         int64     `json:"q7"`
	Q8         int64     `json:"q8"`
	Q9         int64     `json:"q9"`
	Feedback   string    `json:"feedback,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type SurveySKMStats struct {
	Total    int64   `json:"total"`
	SKMScore float64 `json:"skm_score"`
	AvgQ1    float64 `json:"avg_q1"`
	AvgQ2    float64 `json:"avg_q2"`
	AvgQ3    float64 `json:"avg_q3"`
	AvgQ4    float64 `json:"avg_q4"`
	AvgQ5    float64 `json:"avg_q5"`
	AvgQ6    float64 `json:"avg_q6"`
	AvgQ7    float64 `json:"avg_q7"`
	AvgQ8    float64 `json:"avg_q8"`
	AvgQ9    float64 `json:"avg_q9"`
}

type SurveySKMGroupStat struct {
	Label    string  `json:"label"`
	Count    int64   `json:"count"`
	AvgScore float64 `json:"avg_score"`
}

type SurveySKMService struct {
	querier *queries.Querier
}

func NewSurveySKMService(querier *queries.Querier) *SurveySKMService {
	return &SurveySKMService{querier: querier}
}

func (s *SurveySKMService) Create(ctx context.Context, age int64, gender, education, occupation string, year int64, q1, q2, q3, q4, q5, q6, q7, q8, q9 int64, feedback string) (*SurveySKMItem, error) {
	r, err := s.querier.CreateSurveySKM(ctx, queries.CreateSurveySKMParams{
		Age:        age,
		Gender:     gender,
		Education:  education,
		Occupation: occupation,
		Year:       year,
		Q1:         q1,
		Q2:         q2,
		Q3:         q3,
		Q4:         q4,
		Q5:         q5,
		Q6:         q6,
		Q7:         q7,
		Q8:         q8,
		Q9:         q9,
		Feedback:   sql.NullString{String: feedback, Valid: feedback != ""},
	})
	if err != nil {
		return nil, err
	}
	return skmFromRow(&r), nil
}

func (s *SurveySKMService) List(ctx context.Context, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	rows, err := s.querier.ListSurveySKMPaginated(ctx, queries.ListSurveySKMPaginatedParams{
		Limit:  int64(perPage),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, err
	}

	items := make([]SurveySKMItem, 0, len(rows))
	for _, r := range rows {
		items = append(items, *skmFromRow(&r))
	}

	total, err := s.querier.CountSurveySKM(ctx)
	if err != nil {
		return nil, err
	}
	return BuildPagination(items, page, perPage, total), nil
}

func (s *SurveySKMService) GetStats(ctx context.Context) (*SurveySKMStats, error) {
	r, err := s.querier.GetSurveySKMStats(ctx)
	if err != nil {
		return nil, err
	}
	return &SurveySKMStats{
		Total:    r.Total,
		SKMScore: r.SkmScore,
		AvgQ1:    r.AvgQ1,
		AvgQ2:    r.AvgQ2,
		AvgQ3:    r.AvgQ3,
		AvgQ4:    r.AvgQ4,
		AvgQ5:    r.AvgQ5,
		AvgQ6:    r.AvgQ6,
		AvgQ7:    r.AvgQ7,
		AvgQ8:    r.AvgQ8,
		AvgQ9:    r.AvgQ9,
	}, nil
}

func (s *SurveySKMService) GetByGender(ctx context.Context) ([]SurveySKMGroupStat, error) {
	rows, err := s.querier.GetSurveySKMByGender(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]SurveySKMGroupStat, 0, len(rows))
	for _, r := range rows {
		items = append(items, SurveySKMGroupStat{Label: r.Gender, Count: r.Count, AvgScore: r.AvgScore})
	}
	return items, nil
}

func (s *SurveySKMService) GetByEducation(ctx context.Context) ([]SurveySKMGroupStat, error) {
	rows, err := s.querier.GetSurveySKMByEducation(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]SurveySKMGroupStat, 0, len(rows))
	for _, r := range rows {
		items = append(items, SurveySKMGroupStat{Label: r.Education, Count: r.Count, AvgScore: r.AvgScore})
	}
	return items, nil
}

func (s *SurveySKMService) GetByOccupation(ctx context.Context) ([]SurveySKMGroupStat, error) {
	rows, err := s.querier.GetSurveySKMByOccupation(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]SurveySKMGroupStat, 0, len(rows))
	for _, r := range rows {
		items = append(items, SurveySKMGroupStat{Label: r.Occupation, Count: r.Count, AvgScore: r.AvgScore})
	}
	return items, nil
}

func skmFromRow(r *queries.RenjanaSurveySkm) *SurveySKMItem {
	feedback := ""
	if r.Feedback.Valid {
		feedback = r.Feedback.String
	}
	return &SurveySKMItem{
		ID:         r.ID,
		Age:        r.Age,
		Gender:     r.Gender,
		Education:  r.Education,
		Occupation: r.Occupation,
		Year:       r.Year,
		Q1:         r.Q1,
		Q2:         r.Q2,
		Q3:         r.Q3,
		Q4:         r.Q4,
		Q5:         r.Q5,
		Q6:         r.Q6,
		Q7:         r.Q7,
		Q8:         r.Q8,
		Q9:         r.Q9,
		Feedback:   feedback,
		CreatedAt:  r.CreatedAt,
	}
}
