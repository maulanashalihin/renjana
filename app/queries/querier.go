package queries

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/maulanashalihin/laju-go/app/models"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrSessionNotFound   = errors.New("session not found")
	ErrSessionExpired    = errors.New("session expired")
)

// Querier wraps the generated Queries with domain-level error handling
// and conversion to models.User where needed.
type Querier struct {
	*Queries
}

func NewQuerier(db DBTX) *Querier {
	return &Querier{
		Queries: New(db),
	}
}

// --- User helpers that convert queries.User -> models.User ---

// toModelUser converts a queries.User (full struct) to models.User
func toModelUser(qUser User) *models.User {
	return &models.User{
		ID:            qUser.ID,
		Email:         qUser.Email,
		Name:          qUser.Name,
		Password:      qUser.Password,
		Avatar:        nullStringToString(qUser.Avatar),
		Role:          models.UserRole(qUser.Role),
		GoogleID:      qUser.GoogleID,
		EmailVerified: qUser.EmailVerified,
		DistrictID:    qUser.DistrictID,
		VolunteerID:   qUser.VolunteerID,
		IsActive:      qUser.IsActive,
		CreatedAt:     qUser.CreatedAt,
		UpdatedAt:     qUser.UpdatedAt,
	}
}

// toModelUserFromRow converts sqlc-generated *Row types to models.User
func toModelUserFromIDRow(qUser GetUserByIDRow) *models.User {
	return &models.User{
		ID:            qUser.ID,
		Email:         qUser.Email,
		Name:          qUser.Name,
		Password:      qUser.Password,
		Avatar:        nullStringToString(qUser.Avatar),
		Role:          models.UserRole(qUser.Role),
		GoogleID:      qUser.GoogleID,
		EmailVerified: qUser.EmailVerified,
		DistrictID:    qUser.DistrictID,
		VolunteerID:   qUser.VolunteerID,
		IsActive:      qUser.IsActive,
		CreatedAt:     qUser.CreatedAt,
		UpdatedAt:     qUser.UpdatedAt,
	}
}

func toModelUserFromEmailRow(qUser GetUserByEmailRow) *models.User {
	return &models.User{
		ID:            qUser.ID,
		Email:         qUser.Email,
		Name:          qUser.Name,
		Password:      qUser.Password,
		Avatar:        nullStringToString(qUser.Avatar),
		Role:          models.UserRole(qUser.Role),
		GoogleID:      qUser.GoogleID,
		EmailVerified: qUser.EmailVerified,
		DistrictID:    qUser.DistrictID,
		VolunteerID:   qUser.VolunteerID,
		IsActive:      qUser.IsActive,
		CreatedAt:     qUser.CreatedAt,
		UpdatedAt:     qUser.UpdatedAt,
	}
}

func toModelUserFromGoogleRow(qUser GetUserByGoogleIDRow) *models.User {
	return &models.User{
		ID:            qUser.ID,
		Email:         qUser.Email,
		Name:          qUser.Name,
		Password:      qUser.Password,
		Avatar:        nullStringToString(qUser.Avatar),
		Role:          models.UserRole(qUser.Role),
		GoogleID:      qUser.GoogleID,
		EmailVerified: qUser.EmailVerified,
		DistrictID:    qUser.DistrictID,
		VolunteerID:   qUser.VolunteerID,
		IsActive:      qUser.IsActive,
		CreatedAt:     qUser.CreatedAt,
		UpdatedAt:     qUser.UpdatedAt,
	}
}

// toModelUserFromListRow converts ListUsersPaginatedRow to models.User
func toModelUserFromListRow(qUser ListUsersPaginatedRow) *models.User {
	return &models.User{
		ID:            qUser.ID,
		Email:         qUser.Email,
		Name:          qUser.Name,
		Password:      qUser.Password,
		Avatar:        nullStringToString(qUser.Avatar),
		Role:          models.UserRole(qUser.Role),
		GoogleID:      qUser.GoogleID,
		EmailVerified: qUser.EmailVerified,
		DistrictID:    qUser.DistrictID,
		VolunteerID:   qUser.VolunteerID,
		IsActive:      qUser.IsActive,
		CreatedAt:     qUser.CreatedAt,
		UpdatedAt:     qUser.UpdatedAt,
	}
}

func toModelUserFromListByRoleRow(qUser ListUsersByRoleRow) *models.User {
	return &models.User{
		ID:            qUser.ID,
		Email:         qUser.Email,
		Name:          qUser.Name,
		Password:      qUser.Password,
		Avatar:        nullStringToString(qUser.Avatar),
		Role:          models.UserRole(qUser.Role),
		GoogleID:      qUser.GoogleID,
		EmailVerified: qUser.EmailVerified,
		DistrictID:    qUser.DistrictID,
		VolunteerID:   qUser.VolunteerID,
		IsActive:      qUser.IsActive,
		CreatedAt:     qUser.CreatedAt,
		UpdatedAt:     qUser.UpdatedAt,
	}
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// --- User operations ---

func (q *Querier) CreateUser(ctx context.Context, user *models.User) error {
	id, err := q.Queries.CreateUser(ctx, CreateUserParams{
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		Role:      string(user.Role),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		if isDuplicateEmail(err) {
			return ErrUserAlreadyExists
		}
		return err
	}
	user.ID = id
	return nil
}

func (q *Querier) CreateUserWithGoogleID(ctx context.Context, user *models.User) error {
	id, err := q.Queries.CreateUserWithGoogleID(ctx, CreateUserWithGoogleIDParams{
		Email:         user.Email,
		Name:          user.Name,
		GoogleID:      user.GoogleID,
		Avatar:        sql.NullString{String: user.Avatar, Valid: user.Avatar != ""},
		EmailVerified: user.EmailVerified,
		Role:          string(user.Role),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})
	if err != nil {
		if isDuplicateEmail(err) {
			return ErrUserAlreadyExists
		}
		return err
	}
	user.ID = id
	return nil
}

// CreateUserFull creates a user with full RBAC details (admin-only operation)
func (q *Querier) CreateUserFull(ctx context.Context, user *models.User) error {
	isActive := user.IsActive
	if !isActive && user.CreatedAt.IsZero() {
		isActive = true
	}
	id, err := q.Queries.CreateUserFull(ctx, CreateUserFullParams{
		Email:       user.Email,
		Name:        user.Name,
		Password:    user.Password,
		Role:        string(user.Role),
		DistrictID:  user.DistrictID,
		VolunteerID: user.VolunteerID,
		IsActive:    isActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		if isDuplicateEmail(err) {
			return ErrUserAlreadyExists
		}
		return err
	}
	user.ID = id
	user.IsActive = isActive
	return nil
}

func (q *Querier) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	qUser, err := q.Queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return toModelUserFromIDRow(qUser), nil
}

func (q *Querier) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	qUser, err := q.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return toModelUserFromEmailRow(qUser), nil
}

func (q *Querier) GetUserByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
	qUser, err := q.Queries.GetUserByGoogleID(ctx, sql.NullString{String: googleID, Valid: googleID != ""})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return toModelUserFromGoogleRow(qUser), nil
}

func (q *Querier) UpdateUser(ctx context.Context, user *models.User) error {
	rows, err := q.Queries.UpdateUser(ctx, UpdateUserParams{
		Name:          user.Name,
		Avatar:        sql.NullString{String: user.Avatar, Valid: user.Avatar != ""},
		EmailVerified: user.EmailVerified,
		UpdatedAt:     time.Now(),
		ID:            user.ID,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (q *Querier) UpdateUserPassword(ctx context.Context, id int64, hashedPassword string) error {
	rows, err := q.Queries.UpdateUserPassword(ctx, UpdateUserPasswordParams{
		Password:  sql.NullString{String: hashedPassword, Valid: true},
		UpdatedAt: time.Now(),
		ID:        id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (q *Querier) UpdateUserAvatar(ctx context.Context, id int64, avatarURL string) error {
	rows, err := q.Queries.UpdateUserAvatar(ctx, UpdateUserAvatarParams{
		Avatar:    sql.NullString{String: avatarURL, Valid: avatarURL != ""},
		UpdatedAt: time.Now(),
		ID:        id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

// UpdateUserRole updates the role, district_id, and volunteer_id of a user
func (q *Querier) UpdateUserRole(ctx context.Context, id int64, role models.UserRole, districtID, volunteerID sql.NullInt64) error {
	rows, err := q.Queries.UpdateUserRole(ctx, UpdateUserRoleParams{
		Role:        string(role),
		DistrictID:  districtID,
		VolunteerID: volunteerID,
		UpdatedAt:   time.Now(),
		ID:          id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

// SetUserActive toggles is_active status
func (q *Querier) SetUserActive(ctx context.Context, id int64, isActive bool) error {
	rows, err := q.Queries.SetUserActive(ctx, SetUserActiveParams{
		IsActive:  isActive,
		UpdatedAt: time.Now(),
		ID:        id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (q *Querier) DeleteUser(ctx context.Context, id int64) error {
	rows, err := q.Queries.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (q *Querier) SetUserRoleAdmin(ctx context.Context, id int64) error {
	rows, err := q.Queries.SetUserRoleAdmin(ctx, SetUserRoleAdminParams{
		Role:      string(models.RoleAdmin),
		UpdatedAt: time.Now(),
		ID:        id,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

// ListUsersPaginated returns paginated users
func (q *Querier) ListUsersPaginated(ctx context.Context, limit, offset int64) ([]*models.User, error) {
	rows, err := q.Queries.ListUsersPaginated(ctx, ListUsersPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0, len(rows))
	for _, r := range rows {
		users = append(users, toModelUserFromListRow(r))
	}
	return users, nil
}

// ListUsersByRolePaginated returns paginated users filtered by role
func (q *Querier) ListUsersByRolePaginated(ctx context.Context, role models.UserRole, limit, offset int64) ([]*models.User, error) {
	rows, err := q.Queries.ListUsersByRole(ctx, ListUsersByRoleParams{
		Role:   string(role),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0, len(rows))
	for _, r := range rows {
		users = append(users, toModelUserFromListByRoleRow(r))
	}
	return users, nil
}

// --- Session operations ---

func (q *Querier) CreateSession(ctx context.Context, session *Session) error {
	return q.Queries.CreateSession(ctx, CreateSessionParams{
		ID:        session.ID,
		UserID:    session.UserID,
		Data:      session.Data,
		ExpiresAt: session.ExpiresAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (q *Querier) GetSessionByID(ctx context.Context, id string) (*Session, error) {
	qSession, err := q.Queries.GetSessionByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}

	return &qSession, nil
}

func (q *Querier) GetSessionsByUserID(ctx context.Context, userID int64) ([]*Session, error) {
	qSessions, err := q.Queries.GetSessionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var sessions []*Session
	for _, s := range qSessions {
		if s.ExpiresAt.After(time.Now()) {
			sessions = append(sessions, &s)
		}
	}
	return sessions, nil
}

func (q *Querier) UpdateSession(ctx context.Context, session *Session) error {
	rows, err := q.Queries.UpdateSession(ctx, UpdateSessionParams{
		Data:      session.Data,
		ExpiresAt: session.ExpiresAt,
		UpdatedAt: time.Now(),
		ID:        session.ID,
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrSessionNotFound
	}
	return nil
}

func (q *Querier) DeleteSession(ctx context.Context, id string) error {
	rows, err := q.Queries.DeleteSession(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrSessionNotFound
	}
	return nil
}

func (q *Querier) DeleteSessionsByUserID(ctx context.Context, userID int64) error {
	return q.Queries.DeleteSessionsByUserID(ctx, userID)
}

func (q *Querier) DeleteExpiredSessions(ctx context.Context) error {
	return q.Queries.DeleteExpiredSessions(ctx, time.Now())
}

func (q *Querier) DecodeSessionData(data string) (*models.SessionData, error) {
	var sessionData models.SessionData
	if err := sessionDataFromJSON(data, &sessionData); err != nil {
		return nil, err
	}
	return &sessionData, nil
}

func (q *Querier) EncodeSessionData(data *models.SessionData) (string, error) {
	return sessionDataToJSON(data)
}

// isDuplicateEmail checks if the error is a duplicate email error
func isDuplicateEmail(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "UNIQUE constraint failed: users.email") ||
		strings.Contains(msg, "UNIQUE constraint failed: users.google_id")
}

// VolunteerByUserIDRow is a manually-constructed row type to bypass sqlc generation
// (workaround for sqlc stripping `?` on this specific query).
type VolunteerByUserIDRow struct {
	ID                int64
	Name              string
	School            string
	DistrictID        int64
	Phone             sql.NullString
	Status            string
	AvatarUrl         sql.NullString
	JoinedAt          time.Time
	IsActive          bool
	ApplicationStatus string
	ReviewerID        sql.NullInt64
	ReviewedAt        sql.NullTime
	RejectionReason   sql.NullString
	DistrictName      sql.NullString
}

// GetVolunteerByUserID returns the volunteer linked to a user.
// Looks up via users.volunteer_id — the user record stores the volunteer ID
// (set during onboarding via CreateVolunteerForUserDirect).
func (q *Querier) GetVolunteerByUserID(ctx context.Context, userID int64) (VolunteerByUserIDRow, error) {
	const query = `SELECT v.id, v.name, v.school, v.district_id, v.phone, v.status, v.avatar_url, v.joined_at, v.is_active, v.application_status, v.reviewer_id, v.reviewed_at, v.rejection_reason, d.name AS district_name
FROM renjana_volunteers v
LEFT JOIN renjana_districts d ON d.id = v.district_id
WHERE v.id = (SELECT volunteer_id FROM users WHERE id = ?)`
	var row VolunteerByUserIDRow
	err := q.db.QueryRowContext(ctx, query, userID).Scan(
		&row.ID, &row.Name, &row.School, &row.DistrictID, &row.Phone, &row.Status,
		&row.AvatarUrl, &row.JoinedAt, &row.IsActive, &row.ApplicationStatus,
		&row.ReviewerID, &row.ReviewedAt, &row.RejectionReason,
		&row.DistrictName,
	)
	return row, err
}

// CreateVolunteerForUserDirect creates a volunteer linked to a user.
// The volunteer gets its own auto-increment ID; the user record is updated
// with volunteer_id for reverse lookup. Returns the new volunteer ID.
func (q *Querier) CreateVolunteerForUserDirect(ctx context.Context, userID int64, name, school string, districtID int64, phone string, joinedAt time.Time, avatarURL string) (int64, error) {
	const ins = `INSERT INTO renjana_volunteers (
    name, school, district_id, phone, status, avatar_url, joined_at,
    is_active, application_status
)
VALUES (?, ?, ?, ?, 'aktif', ?, ?, 1, 'approved')
RETURNING id`
	var volunteerID int64
	var avatarArg interface{}
	if avatarURL != "" {
		avatarArg = avatarURL
	}
	err := q.db.QueryRowContext(ctx, ins, name, school, districtID, phone, avatarArg, joinedAt).Scan(&volunteerID)
	if err != nil {
		return 0, err
	}
	// Set volunteer_id on the user record (for reverse lookup)
	_, _ = q.db.ExecContext(ctx, `UPDATE users SET volunteer_id = ? WHERE id = ?`, volunteerID, userID)
	return volunteerID, nil
}

// UpdateVolunteerProfile updates a volunteer's name and avatar_url pulled from their linked user account.
func (q *Querier) UpdateVolunteerProfile(ctx context.Context, volunteerID int64, name, avatarURL string) error {
	const query = `UPDATE renjana_volunteers SET name = ?, avatar_url = ? WHERE id = ?`
	_, err := q.db.ExecContext(ctx, query, name, avatarURL, volunteerID)
	return err
}
