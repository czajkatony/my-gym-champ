package models

import (
	"time"
)

// Gym represents a gym entity
type Gym struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	ZipCode     string    `json:"zipCode"`
	Country     string    `json:"country"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Website     string    `json:"website,omitempty"`
	Description string    `json:"description"`
	LogoImage   string    `json:"logoImage,omitempty"`
	OpenHours   string    `json:"openHours"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GymCreateRequest represents the data required to create a new gym
type GymCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required"`
	ZipCode     string `json:"zipCode" binding:"required"`
	Country     string `json:"country" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Website     string `json:"website"`
	Description string `json:"description"`
	OpenHours   string `json:"openHours" binding:"required"`
}

// GymResponse represents the gym data sent in responses
type GymResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	ZipCode     string  `json:"zipCode"`
	Country     string  `json:"country"`
	PhoneNumber string  `json:"phoneNumber"`
	Email       string  `json:"email"`
	Website     string  `json:"website,omitempty"`
	Description string  `json:"description"`
	LogoImage   string  `json:"logoImage,omitempty"`
	OpenHours   string  `json:"openHours"`
	MemberCount int     `json:"memberCount"`
	Rating      float64 `json:"rating,omitempty"`
} 