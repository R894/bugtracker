package aggregate

import (
	"bugtracker/internal/domain/bug/entity"
	"time"

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
	b.Title = title
	b.Description = description
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

func generateUniqueID() string {
	return uuid.New().String()
}
