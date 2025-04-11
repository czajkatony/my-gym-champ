package models

import (
	"time"
)

// Workout represents a member's workout session
type Workout struct {
	ID        int64      `json:"id"`
	MemberID  int64      `json:"memberId"`
	Date      time.Time  `json:"date"`
	Notes     string     `json:"notes,omitempty"`
	Duration  int        `json:"duration"` // Duration in minutes
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Exercises []Exercise `json:"exercises,omitempty"`
}

// Exercise represents a specific exercise in a workout
type Exercise struct {
	ID               int64     `json:"id"`
	WorkoutID        int64     `json:"workoutId"`
	Name             string    `json:"name"`
	Sets             []Set     `json:"sets"`
	Notes            string    `json:"notes,omitempty"`
	IsPersonalRecord bool      `json:"isPersonalRecord"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// Set represents a set of an exercise
type Set struct {
	ID         int64     `json:"id"`
	ExerciseID int64     `json:"exerciseId"`
	Reps       int       `json:"reps"`
	Weight     float64   `json:"weight"`
	Unit       string    `json:"unit"` // "kg" or "lb"
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// WorkoutCreateRequest represents the data required to create a new workout
type WorkoutCreateRequest struct {
	MemberID  int64                   `json:"memberId" binding:"required"`
	Date      time.Time               `json:"date" binding:"required"`
	Notes     string                  `json:"notes"`
	Duration  int                     `json:"duration" binding:"required"` // Duration in minutes
	Exercises []ExerciseCreateRequest `json:"exercises"`
}

// ExerciseCreateRequest represents the data required to create a new exercise
type ExerciseCreateRequest struct {
	Name             string             `json:"name" binding:"required"`
	Sets             []SetCreateRequest `json:"sets" binding:"required"`
	Notes            string             `json:"notes"`
	IsPersonalRecord bool               `json:"isPersonalRecord"`
}

// SetCreateRequest represents the data required to create a new set
type SetCreateRequest struct {
	Reps   int     `json:"reps" binding:"required"`
	Weight float64 `json:"weight" binding:"required"`
	Unit   string  `json:"unit" binding:"required"` // "kg" or "lb"
}
