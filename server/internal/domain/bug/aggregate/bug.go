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
	Comments    *[]entity.Comment  `json:"comments"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func NewBug(title, description string) *Bug {
	return &Bug{
		ID:          generateUniqueID(),
		Title:       title,
		Description: description,
		Status:      entity.BugStatusOpen,
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
