package domain

import "time"

type Task struct {
	Id          uint64
	UserId      uint64
	Title       string
	Description *string
	Status      TaskStatus
	Date        *time.Time
}

type TaskStatus string

const (
	Draft     TaskStatus = "DRAFT"
	Assigned  TaskStatus = "ASSIGNED"
	Completed TaskStatus = "COMPLETED"
)
