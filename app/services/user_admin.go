package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
)

// UserAdminService handles admin-only user management operations.
type UserAdminService struct {
	querier *queries.Querier
}

func NewUserAdminService(querier *queries.Querier) *UserAdminService {
	return &UserAdminService{querier: querier}
}

// UserFilter holds optional filters for listing users.
type UserFilter struct {
	Role   models.UserRole
	Search string
}

// ListUsers returns paginated users, optionally filtered by role and search term.
func (s *UserAdminService) ListUsers(ctx context.Context, filter UserFilter, page, perPage int) (*PaginationResult, error) {
	page, perPage, offset := NormalizePagination(page, perPage)

	// For now, use simple list-all (filtering would require more queries)
	users, err := s.querier.ListUsersPaginated(ctx, int64(perPage), int64(offset))
	if err != nil {
		return nil, err
	}

	total, err := s.querier.CountUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Filter in memory (simple search by name/email)
	if filter.Search != "" {
		filtered := make([]*models.User, 0)
		for _, u := range users {
			if contains(u.Name, filter.Search) || contains(u.Email, filter.Search) {
				filtered = append(filtered, u)
			}
		}
		users = filtered
	}

	if filter.Role != "" {
		filtered := make([]*models.User, 0)
		for _, u := range users {
			if u.Role == filter.Role {
				filtered = append(filtered, u)
			}
		}
		users = filtered
	}

	return BuildPagination(users, page, perPage, total), nil
}

func contains(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

// CreateUser creates a new user (admin operation).
func (s *UserAdminService) CreateUser(ctx context.Context, name, email, password string, role models.UserRole, districtID, volunteerID int64) (*models.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email, and password are required")
	}

	if !role.IsValid() {
		return nil, errors.New("invalid role")
	}

	// Hash password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Name:     name,
		Password: sql.NullString{String: hashedPassword, Valid: true},
		Role:     role,
	}
	if districtID > 0 {
		user.DistrictID = sql.NullInt64{Int64: districtID, Valid: true}
	}
	if volunteerID > 0 {
		user.VolunteerID = sql.NullInt64{Int64: volunteerID, Valid: true}
	}
	user.IsActive = true

	err = s.querier.CreateUserFull(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserRole updates a user's role, district_id, and volunteer_id.
func (s *UserAdminService) UpdateUserRole(ctx context.Context, userID int64, role models.UserRole, districtID, volunteerID int64) error {
	if !role.IsValid() {
		return errors.New("invalid role")
	}

	var did, vid sql.NullInt64
	if districtID > 0 {
		did = sql.NullInt64{Int64: districtID, Valid: true}
	}
	if volunteerID > 0 {
		vid = sql.NullInt64{Int64: volunteerID, Valid: true}
	}

	return s.querier.UpdateUserRole(ctx, userID, role, did, vid)
}

// SetActive toggles a user's is_active status.
func (s *UserAdminService) SetActive(ctx context.Context, userID int64, isActive bool) error {
	return s.querier.SetUserActive(ctx, userID, isActive)
}

// DeleteUser deletes a user (admin operation).
func (s *UserAdminService) DeleteUser(ctx context.Context, userID int64) error {
	return s.querier.DeleteUser(ctx, userID)
}

// GetUser returns a single user by ID.
func (s *UserAdminService) GetUser(ctx context.Context, userID int64) (*models.User, error) {
	return s.querier.GetUserByID(ctx, userID)
}

// CountByRole returns user count by role.
func (s *UserAdminService) CountByRole(ctx context.Context, role models.UserRole) (int64, error) {
	return s.querier.CountUsersByRole(ctx, string(role))
}

// PromoteToAdmin looks up a user by email and promotes them to admin.
// Returns an error if the user is not found or is already an admin.
func (s *UserAdminService) PromoteToAdmin(ctx context.Context, email string) error {
	user, err := s.querier.GetUserByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}
	if user.Role == models.RoleAdmin {
		return errors.New("user is already an admin")
	}
	if !user.IsActive {
		return errors.New("user is not active")
	}
	return s.querier.UpdateUserRole(ctx, user.ID, models.RoleAdmin, user.DistrictID, user.VolunteerID)
}
