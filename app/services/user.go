package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
)

type UserService struct {
	querier *queries.Querier
}

func NewUserService(querier *queries.Querier) *UserService {
	return &UserService{
		querier: querier,
	}
}

// GetProfile retrieves a user's profile directly from DB.
func (s *UserService) GetProfile(userID int64) (*models.UserResponse, error) {
	user, err := s.querier.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

// GetProfileByEmail retrieves a user's profile by email
func (s *UserService) GetProfileByEmail(email string) (*models.User, error) {
	return s.querier.GetUserByEmail(context.Background(), email)
}

// UpdatePassword updates a user's password
func (s *UserService) UpdatePassword(userID int64, hashedPassword string) error {
	return s.querier.UpdateUserPassword(context.Background(), userID, hashedPassword)
}

// UpdateAvatar updates a user's avatar URL
func (s *UserService) UpdateAvatar(userID int64, avatarURL string) error {
	return s.querier.UpdateUserAvatar(context.Background(), userID, avatarURL)
}

// UpdateProfile updates a user's profile
func (s *UserService) UpdateProfile(userID int64, req models.UpdateProfileRequest) (*models.UserResponse, error) {
	user, err := s.querier.GetUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	// Update user fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	if err := s.querier.UpdateUser(context.Background(), user); err != nil {
		return nil, err
	}

	// Sync ke volunteer record kalau user punya volunteer account
	vol, volErr := s.querier.GetVolunteerByUserID(context.Background(), userID)
	if volErr == nil {
		// Update nama & avatar
		_ = s.querier.UpdateVolunteerProfile(context.Background(), vol.ID, req.Name, user.Avatar)

		// Update school, phone, district_id kalau dikirim
		if req.School != "" || req.Phone != "" || req.DistrictID != 0 {
			phone := sql.NullString{}
			if req.Phone != "" {
				phone = sql.NullString{String: req.Phone, Valid: true}
			}
			school := vol.School
			if req.School != "" {
				school = req.School
			}
			districtID := vol.DistrictID
			if req.DistrictID != 0 {
				districtID = req.DistrictID
			}
			_, _ = s.querier.UpdateVolunteer(context.Background(), queries.UpdateVolunteerParams{
				Name:       vol.Name,
				School:     school,
				DistrictID: districtID,
				Phone:      phone,
				Status:     vol.Status,
				JoinedAt:   vol.JoinedAt,
				ID:         vol.ID,
			})
		}
	}

	response := user.ToResponse()
	return &response, nil
}

// ChangePassword changes a user's password
func (s *UserService) ChangePassword(userID int64, oldPassword, newPassword string) error {
	user, err := s.querier.GetUserByID(context.Background(), userID)
	if err != nil {
		return err
	}

	// Verify old password - user must have a password
	if !user.Password.Valid {
		return errors.New("invalid current password")
	}

	if !checkPassword(user.Password.String, oldPassword) {
		return errors.New("invalid current password")
	}

	// Hash new password
	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.querier.UpdateUserPassword(context.Background(), userID, hashedPassword)
}

// DeleteAccount deletes a user's account
func (s *UserService) DeleteAccount(userID int64) error {
	return s.querier.DeleteUser(context.Background(), userID)
}

// IsAdmin checks if a user is an admin (direct DB query).
func (s *UserService) IsAdmin(userID int64) (bool, error) {
	user, err := s.querier.GetUserByID(context.Background(), userID)
	if err != nil {
		return false, err
	}

	return user.Role == models.RoleAdmin, nil
}
