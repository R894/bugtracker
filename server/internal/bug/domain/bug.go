package domain

import (
	"time"

	"github.com/google/uuid"
)

type BugStatus string

const (
	BugStatusOpen    BugStatus = "Open"
	BugStatusClosed  BugStatus = "Closed"
	BugStatusPending BugStatus = "Pending"
)

type BugPriority string

const (
	BugPriorityLow    BugPriority = "Low"
	BugPriorityMedium BugPriority = "Medium"
	BugPriorityHigh   BugPriority = "High"
)

type Bug struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      BugStatus   `json:"status"`
	Priority    BugPriority `json:"priority"`
	Assignee    string      `json:"asignee"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func NewBug(title, description string) *Bug {
	return &Bug{
		ID:          generateUniqueID(),
		Title:       title,
		Description: description,
		Status:      BugStatusOpen,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (b *Bug) UpdateDetails(title, description string) {
	b.Title = title
	b.Description = description
	b.UpdatedAt = time.Now()
}

func (b *Bug) ChangeStatus(status BugStatus) {
	b.Status = status
	b.UpdatedAt = time.Now()
}

func (b *Bug) AssignTo(assignee string) {
	b.Assignee = assignee
	b.UpdatedAt = time.Now()
}

func generateUniqueID() string {
	return uuid.New().String()
}
