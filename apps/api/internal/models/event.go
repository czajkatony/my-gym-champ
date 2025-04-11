package models

import (
	"time"
)

// Event represents a gym event
type Event struct {
	ID          int64     `json:"id"`
	GymID       int64     `json:"gymId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Location    string    `json:"location"`
	ImageURL    string    `json:"imageUrl,omitempty"`
	MaxCapacity int       `json:"maxCapacity,omitempty"`
	IsPaid      bool      `json:"isPaid"`
	Price       float64   `json:"price,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedBy   int64     `json:"createdBy"` // ID of the gym owner/admin who created it
}

// EventCreateRequest represents the data required to create a new event
type EventCreateRequest struct {
	GymID       int64     `json:"gymId" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartTime   time.Time `json:"startTime" binding:"required"`
	EndTime     time.Time `json:"endTime" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	ImageURL    string    `json:"imageUrl"`
	MaxCapacity int       `json:"maxCapacity"`
	IsPaid      bool      `json:"isPaid"`
	Price       float64   `json:"price"`
}

// EventResponse represents the event data sent in responses
type EventResponse struct {
	ID             int64     `json:"id"`
	GymID          int64     `json:"gymId"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	Location       string    `json:"location"`
	ImageURL       string    `json:"imageUrl,omitempty"`
	MaxCapacity    int       `json:"maxCapacity,omitempty"`
	IsPaid         bool      `json:"isPaid"`
	Price          float64   `json:"price,omitempty"`
	AttendeesCount int       `json:"attendeesCount"`
}

// EventAttendee represents a member attending an event
type EventAttendee struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"eventId"`
	MemberID  int64     `json:"memberId"`
	Status    string    `json:"status"` // "registered", "attended", "cancelled"
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
