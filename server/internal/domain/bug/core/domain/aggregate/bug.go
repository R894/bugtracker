package aggregate

import (
	"time"

	"github.com/R894/bugtracker/internal/domain/bug/core/domain/entity"

	"github.com/google/uuid"
)

type Bug struct {
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      entity.BugStatus   `json:"status"`
	Priority    entity.BugPriority `json:"priority"`
	Assignee    string             `json:"asignee"`
	ProjectId   string             `json:"projectId"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

type CreateBugRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectId   string `json:"projectId"`
}

type UpdateBugRequest struct {
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      entity.BugStatus   `json:"status"`
	Priority    entity.BugPriority `json:"priority"`
	Assignee    string             `json:"asignee"`
	UpdatedAt   time.Time          `json:"-"`
}

type UpdateBugPriority struct {
	Priority entity.BugPriority `json:"priority"`
}

func NewBug(req CreateBugRequest) *Bug {
	return &Bug{
		ID:          generateUniqueID(),
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.BugStatusOpen,
		ProjectId:   req.ProjectId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (b *Bug) UpdateDetails(title, description string) {
	if title == "" && description == "" {
		return // No need to update if both title and description are empty
	}
	if title == b.Title && description == b.Description {
		return // No need to update if both title and description are the same
	}
	if title != "" {
		b.Title = title
	}
	if description != "" {
		b.Description = description
	}
	b.UpdatedAt = time.Now()
}

func (b *Bug) ChangeStatus(status entity.BugStatus) {
	b.Status = status
	b.UpdatedAt = time.Now()
}

func (b *Bug) AssignTo(assignee string) {
	b.Assignee = assignee
	b.UpdatedAt = time.Now()
}

func (b *Bug) UpdatePriority(priority entity.BugPriority) {
	b.Priority = priority
	b.UpdatedAt = time.Now()
}

func generateUniqueID() string {
	return uuid.New().String()
}
