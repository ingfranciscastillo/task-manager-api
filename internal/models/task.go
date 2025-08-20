package models

import (
	"time"
)

// TaskStatus representa el estado de una tarea
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusCompleted TaskStatus = "completed"
)

// Task representa una tarea en el sistema
type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status" gorm:"default:'pending'"`
	UserID      uint       `json:"user_id" gorm:"not null;index"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relación con usuario
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TaskSummary representa el resumen de tareas de un usuario
type TaskSummary struct {
	TotalTasks     int64 `json:"total_tasks"`
	CompletedTasks int64 `json:"completed_tasks"`
	PendingTasks   int64 `json:"pending_tasks"`
}
