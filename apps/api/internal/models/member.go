package models

import (
	"time"
)

// Member represents a gym member
type Member struct {
	ID              int64      `json:"id"`
	Email           string     `json:"email"`
	PasswordHash    string     `json:"-"` // Not exposed in JSON responses
	FirstName       string     `json:"firstName"`
	LastName        string     `json:"lastName"`
	MembershipLevel string     `json:"membershipLevel"`
	GymID           int64      `json:"gymId"`
	JoinDate        time.Time  `json:"joinDate"`
	LastCheckIn     *time.Time `json:"lastCheckIn,omitempty"`
	ProfileImage    *string    `json:"profileImage,omitempty"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}

// MemberSignupRequest represents the data required to sign up a new member
type MemberSignupRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=8"`
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	GymID        int64  `json:"gymId" binding:"required"`
}

// MemberLoginRequest represents the data required for member login
type MemberLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// MemberProfileResponse represents the member data sent in responses
type MemberProfileResponse struct {
	ID              int64      `json:"id"`
	Email           string     `json:"email"`
	FirstName       string     `json:"firstName"`
	LastName        string     `json:"lastName"`
	MembershipLevel string     `json:"membershipLevel"`
	GymID           int64      `json:"gymId"`
	JoinDate        time.Time  `json:"joinDate"`
	LastCheckIn     *time.Time `json:"lastCheckIn,omitempty"`
	ProfileImage    *string    `json:"profileImage,omitempty"`
} 