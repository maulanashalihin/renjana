package models

import (
	"database/sql"
	"time"
)

// UserRole represents the role of a user in the RENJANA system.
// The role hierarchy (highest → lowest):
//
//	super_admin → admin → koordinator → relawan
type UserRole string

const (
	// RoleRelawan — registered volunteer, can only see/edit own profile
	RoleRelawan UserRole = "relawan"
	// RoleKoordinator — district-level coordinator, scoped to single district
	RoleKoordinator UserRole = "koordinator"
	// RoleAdmin — full CRUD access across all districts
	RoleAdmin UserRole = "admin"
	// RoleSuperAdmin — future: system configuration
	RoleSuperAdmin UserRole = "super_admin"
	// RoleUser — backward-compatibility alias for RoleRelawan.
	// Deprecated: use RoleRelawan instead.
	RoleUser UserRole = RoleRelawan
)

// AllRoles returns the ordered list of valid roles.
func AllRoles() []UserRole {
	return []UserRole{RoleRelawan, RoleKoordinator, RoleAdmin, RoleSuperAdmin}
}

// IsValid checks if the role is one of the defined roles.
func (r UserRole) IsValid() bool {
	for _, valid := range AllRoles() {
		if r == valid {
			return true
		}
	}
	return false
}

// CanManageUsers returns true if the role can manage other users.
func (r UserRole) CanManageUsers() bool {
	return r == RoleAdmin || r == RoleSuperAdmin
}

// CanCRUDAll returns true if the role can CRUD across all districts.
func (r UserRole) CanCRUDAll() bool {
	return r == RoleAdmin || r == RoleSuperAdmin
}

// User represents a user account in the RENJANA system.
type User struct {
	ID            int64          `json:"id"`
	Email         string         `json:"email"`
	Name          string         `json:"name"`
	Avatar        string         `json:"avatar"`
	Password      sql.NullString `json:"-"` // Hashed password, never return in JSON (NULL for OAuth users)
	Role          UserRole       `json:"role"`
	GoogleID      sql.NullString `json:"-"` // OAuth provider ID (NULL for email/password users)
	EmailVerified bool           `json:"email_verified"`
	DistrictID    sql.NullInt64  `json:"district_id,omitempty"`  // For koordinator scope
	VolunteerID   sql.NullInt64  `json:"volunteer_id,omitempty"` // For relawan link to volunteer record
	IsActive      bool           `json:"is_active"`              // Soft-delete / deactivate
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// TableName returns the table name for User
func (User) TableName() string {
	return "users"
}
