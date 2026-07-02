package services

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/maulanashalihin/laju-go/app/queries"
)

// EducationService handles LMS operations for courses, quizzes, and certificates.
type EducationService struct {
	querier *queries.Querier
}

func NewEducationService(querier *queries.Querier) *EducationService {
	return &EducationService{querier: querier}
}

// CourseItem represents an education course for listing.
type CourseItem struct {
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

// CourseModuleItem represents a single module within a course.
type CourseModuleItem struct {
	ID         int64     `json:"id"`
	CourseID   int64     `json:"course_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	VideoURL   string    `json:"video_url"`
	OrderIndex int64     `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
}

// QuizQuestionItem is a quiz question (without correct_answer for frontend).
type QuizQuestionItem struct {
	ID         int64    `json:"id"`
	CourseID   int64    `json:"course_id"`
	Question   string   `json:"question"`
	Options    []string `json:"options"`
	OrderIndex int64    `json:"order_index"`
}

// QuizQuestionAdmin is a quiz question WITH correct_answer (backend use).
type QuizQuestionAdmin struct {
	ID            int64    `json:"id"`
	CourseID      int64    `json:"course_id"`
	Question      string   `json:"question"`
	Options       []string `json:"options"`
	CorrectOption int64    `json:"correct_option"`
	OrderIndex    int64    `json:"order_index"`
}

// CourseDetail bundles course info, modules, and user progress.
type CourseDetail struct {
	Course         CourseItem         `json:"course"`
	Modules        []CourseModuleItem `json:"modules"`
	Progress       *CourseProgress    `json:"progress,omitempty"`
	QuizCount      int                `json:"quiz_count"`
	CertificateID  int64              `json:"certificate_id,omitempty"`
	HasCertificate bool               `json:"has_certificate"`
}

// CourseProgress represents user progress in a course.
type CourseProgress struct {
	CompletedModules int64      `json:"completed_modules"`
	TotalModules     int64      `json:"total_modules"`
	Completed        bool       `json:"completed"`
	StartedAt        time.Time  `json:"started_at"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
}

// QuizAttemptResult is the result of a quiz attempt.
type QuizAttemptResult struct {
	ID             int64          `json:"id"`
	Score          int64          `json:"score"`
	TotalQuestions int64          `json:"total_questions"`
	Passed         bool           `json:"passed"`
	PassingScore   int64          `json:"passing_score"`
	Answers        []AnswerDetail `json:"answers"`
	StartedAt      time.Time      `json:"started_at"`
	CompletedAt    time.Time      `json:"completed_at"`
}

// AnswerDetail represents each question's answer in the result.
type AnswerDetail struct {
	QuestionID     int64 `json:"question_id"`
	SelectedOption int64 `json:"selected_option"`
	Correct        bool  `json:"correct"`
}

// CertificateItem represents a certificate.
type CertificateItem struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	CourseID        int64     `json:"course_id"`
	CertificateCode string    `json:"certificate_code"`
	Score           int64     `json:"score"`
	IssuedAt        time.Time `json:"issued_at"`
	UserName        string    `json:"user_name,omitempty"`
	UserEmail       string    `json:"user_email,omitempty"`
	CourseTitle     string    `json:"course_title,omitempty"`
	CourseCategory  string    `json:"course_category,omitempty"`
}

// ListCourses returns all published courses.
func (s *EducationService) ListCourses(ctx context.Context) ([]CourseItem, error) {
	rows, err := s.querier.ListEducationCourses(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]CourseItem, 0, len(rows))
	for _, r := range rows {
		age := ""
		if r.AgeGroup.Valid {
			age = r.AgeGroup.String
		}
		cover := ""
		if r.CoverImage.Valid {
			cover = r.CoverImage.String
		}
		dur := int64(0)
		if r.DurationMinutes.Valid {
			dur = r.DurationMinutes.Int64
		}
		items = append(items, CourseItem{
			ID:              r.ID,
			Title:           r.Title,
			Category:        r.Category,
			Body:            r.Body,
			AgeGroup:        age,
			DurationMinutes: dur,
			IsPublished:     r.IsPublished,
			CoverImage:      cover,
			PassingScore:    r.PassingScore,
			TotalModules:    r.TotalModules,
			IsCourse:        r.IsCourse,
			CreatedAt:       r.CreatedAt,
			UpdatedAt:       r.UpdatedAt,
		})
	}
	return items, nil
}

// GetCourseDetail returns a course with its modules and user progress.
func (s *EducationService) GetCourseDetail(ctx context.Context, courseID, userID int64) (*CourseDetail, error) {
	// Get course
	courseRow, err := s.querier.GetEducationCourseByID(ctx, courseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	// Get modules
	moduleRows, err := s.querier.ListCourseModules(ctx, courseID)
	if err != nil {
		return nil, err
	}

	// Get quiz count
	quizRows, err := s.querier.ListQuizQuestions(ctx, courseID)
	if err != nil {
		return nil, err
	}

	// Build course item
	age := ""
	if courseRow.AgeGroup.Valid {
		age = courseRow.AgeGroup.String
	}
	cover := ""
	if courseRow.CoverImage.Valid {
		cover = courseRow.CoverImage.String
	}
	dur := int64(0)
	if courseRow.DurationMinutes.Valid {
		dur = courseRow.DurationMinutes.Int64
	}

	course := CourseItem{
		ID:              courseRow.ID,
		Title:           courseRow.Title,
		Category:        courseRow.Category,
		Body:            courseRow.Body,
		AgeGroup:        age,
		DurationMinutes: dur,
		IsPublished:     courseRow.IsPublished,
		CoverImage:      cover,
		PassingScore:    courseRow.PassingScore,
		TotalModules:    courseRow.TotalModules,
		IsCourse:        courseRow.IsCourse,
		CreatedAt:       courseRow.CreatedAt,
		UpdatedAt:       courseRow.UpdatedAt,
	}

	// Build modules
	modules := make([]CourseModuleItem, 0, len(moduleRows))
	for _, m := range moduleRows {
		videoURL := ""
		if m.VideoUrl.Valid {
			videoURL = m.VideoUrl.String
		}
		modules = append(modules, CourseModuleItem{
			ID:         m.ID,
			CourseID:   m.CourseID,
			Title:      m.Title,
			Content:    m.Content,
			VideoURL:   videoURL,
			OrderIndex: m.OrderIndex,
			CreatedAt:  m.CreatedAt,
		})
	}

	// Build progress
	var progress *CourseProgress
	if userID > 0 {
		prog, err := s.querier.GetCourseProgress(ctx, queries.GetCourseProgressParams{
			UserID:   userID,
			CourseID: courseID,
		})
		if err == nil {
			var completedAt *time.Time
			if prog.CompletedAt.Valid {
				completedAt = &prog.CompletedAt.Time
			}
			progress = &CourseProgress{
				CompletedModules: prog.CompletedModules,
				TotalModules:     prog.TotalModules,
				Completed:        prog.Completed,
				StartedAt:        prog.StartedAt,
				CompletedAt:      completedAt,
			}
		}
	}

	// Check certificate
	hasCert := false
	var certID int64
	if userID > 0 {
		cert, err := s.querier.GetCertificateByUserAndCourse(ctx, queries.GetCertificateByUserAndCourseParams{
			UserID:   userID,
			CourseID: courseID,
		})
		if err == nil {
			hasCert = true
			certID = cert.ID
		}
	}

	return &CourseDetail{
		Course:         course,
		Modules:        modules,
		Progress:       progress,
		QuizCount:      len(quizRows),
		CertificateID:  certID,
		HasCertificate: hasCert,
	}, nil
}

// GetQuizQuestions returns quiz questions WITHOUT correct answers for frontend use.
func (s *EducationService) GetQuizQuestions(ctx context.Context, courseID int64) ([]QuizQuestionItem, error) {
	rows, err := s.querier.ListQuizQuestions(ctx, courseID)
	if err != nil {
		return nil, err
	}
	items := make([]QuizQuestionItem, 0, len(rows))
	for _, r := range rows {
		var opts []string
		if err := json.Unmarshal([]byte(r.Options), &opts); err != nil {
			opts = []string{}
		}
		items = append(items, QuizQuestionItem{
			ID:         r.ID,
			CourseID:   r.CourseID,
			Question:   r.Question,
			Options:    opts,
			OrderIndex: r.OrderIndex,
		})
	}
	return items, nil
}

// SubmitQuizAttempt evaluates quiz answers and returns the result.
func (s *EducationService) SubmitQuizAttempt(ctx context.Context, userID, courseID int64, selectedAnswers map[int64]int64) (*QuizAttemptResult, error) {
	// Get course to find passing score
	courseRow, err := s.querier.GetEducationCourseByID(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("course not found: %w", err)
	}

	// Get all questions with correct answers
	rows, err := s.querier.ListQuizQuestions(ctx, courseID)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("no quiz questions found for this course")
	}

	// Score the answers
	var correctCount int64
	details := make([]AnswerDetail, 0, len(rows))
	for _, q := range rows {
		selected, answered := selectedAnswers[q.ID]
		isCorrect := answered && selected == q.CorrectOption
		if isCorrect {
			correctCount++
		}
		details = append(details, AnswerDetail{
			QuestionID:     q.ID,
			SelectedOption: selected,
			Correct:        isCorrect,
		})
	}

	total := int64(len(rows))
	passingScore := courseRow.PassingScore
	if passingScore == 0 {
		passingScore = 70 // default
	}

	// Calculate percentage
	percentage := int64(math.Round(float64(correctCount) / float64(total) * 100))
	passed := percentage >= passingScore

	// Serialize answers
	answersJSON, _ := json.Marshal(details)

	now := time.Now()

	// Create attempt record
	err = s.querier.CreateQuizAttempt(ctx, queries.CreateQuizAttemptParams{
		UserID:         userID,
		CourseID:       courseID,
		Score:          percentage,
		TotalQuestions: total,
		Passed:         passed,
		Answers:        sql.NullString{String: string(answersJSON), Valid: true},
		StartedAt:      now,
		CompletedAt:    sql.NullTime{Time: now, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// Update progress
	if passed {
		_ = s.querier.UpsertCourseProgress(ctx, queries.UpsertCourseProgressParams{
			UserID:           userID,
			CourseID:         courseID,
			CompletedModules: courseRow.TotalModules,
			TotalModules:     courseRow.TotalModules,
			Completed:        true,
			StartedAt:        now,
			CompletedAt:      sql.NullTime{Time: now, Valid: true},
		})

		// Issue certificate if passed
		code, err := generateCertificateCode()
		if err != nil {
			return nil, fmt.Errorf("failed to generate certificate: %w", err)
		}
		err = s.querier.CreateCertificate(ctx, queries.CreateCertificateParams{
			UserID:          userID,
			CourseID:        courseID,
			CertificateCode: code,
			Score:           percentage,
			IssuedAt:        now,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create certificate: %w", err)
		}
	}

	// Get the created attempt to return its ID
	bestAttempt, err := s.querier.GetBestQuizAttempt(ctx, queries.GetBestQuizAttemptParams{
		UserID:   userID,
		CourseID: courseID,
	})
	var attemptID int64
	if err == nil {
		attemptID = bestAttempt.ID
	}

	return &QuizAttemptResult{
		ID:             attemptID,
		Score:          percentage,
		TotalQuestions: total,
		Passed:         passed,
		PassingScore:   passingScore,
		Answers:        details,
		StartedAt:      now,
		CompletedAt:    now,
	}, nil
}

// GetCertificate returns a certificate by user and course.
func (s *EducationService) GetCertificate(ctx context.Context, userID, courseID int64) (*CertificateItem, error) {
	cert, err := s.querier.GetCertificateByUserAndCourse(ctx, queries.GetCertificateByUserAndCourseParams{
		UserID:   userID,
		CourseID: courseID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	// Get user & course info
	course, err := s.querier.GetEducationCourseByID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	user, err := s.querier.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &CertificateItem{
		ID:              cert.ID,
		UserID:          cert.UserID,
		CourseID:        cert.CourseID,
		CertificateCode: cert.CertificateCode,
		Score:           cert.Score,
		IssuedAt:        cert.IssuedAt,
		UserName:        user.Name,
		UserEmail:       user.Email,
		CourseTitle:     course.Title,
		CourseCategory:  course.Category,
	}, nil
}

// GetCertificateByCode returns a certificate by its code (public access).
func (s *EducationService) GetCertificateByCode(ctx context.Context, code string) (*CertificateItem, error) {
	row, err := s.querier.GetCertificateByCode(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &CertificateItem{
		ID:              row.ID,
		UserID:          row.UserID,
		CourseID:        row.CourseID,
		CertificateCode: row.CertificateCode,
		Score:           row.Score,
		IssuedAt:        row.IssuedAt,
		UserName:        row.UserName,
		UserEmail:       row.UserEmail,
		CourseTitle:     row.CourseTitle,
		CourseCategory:  row.CourseCategory,
	}, nil
}

// ListUserCertificates returns all certificates for a user.
func (s *EducationService) ListUserCertificates(ctx context.Context, userID int64) ([]CertificateItem, error) {
	rows, err := s.querier.ListUserCertificates(ctx, userID)
	if err != nil {
		return nil, err
	}
	items := make([]CertificateItem, 0, len(rows))
	for _, r := range rows {
		items = append(items, CertificateItem{
			ID:              r.ID,
			UserID:          r.UserID,
			CourseID:        r.CourseID,
			CertificateCode: r.CertificateCode,
			Score:           r.Score,
			IssuedAt:        r.IssuedAt,
			UserName:        r.UserName,
			UserEmail:       r.UserEmail,
			CourseTitle:     r.CourseTitle,
			CourseCategory:  r.CourseCategory,
		})
	}
	return items, nil
}

// GetUserProgress returns progress for a specific course.
func (s *EducationService) GetUserProgress(ctx context.Context, userID, courseID int64) (*CourseProgress, error) {
	prog, err := s.querier.GetCourseProgress(ctx, queries.GetCourseProgressParams{
		UserID:   userID,
		CourseID: courseID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	var completedAt *time.Time
	if prog.CompletedAt.Valid {
		completedAt = &prog.CompletedAt.Time
	}
	return &CourseProgress{
		CompletedModules: prog.CompletedModules,
		TotalModules:     prog.TotalModules,
		Completed:        prog.Completed,
		StartedAt:        prog.StartedAt,
		CompletedAt:      completedAt,
	}, nil
}

// generateCertificateCode creates a unique certificate code.
func generateCertificateCode() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	code := hex.EncodeToString(b)
	// Format: RENJANA-XXXXXXXX-XXXXXXXX
	return fmt.Sprintf("RENJANA-%s-%s",
		strings.ToUpper(code[:8]),
		strings.ToUpper(code[8:16]),
	), nil
}
