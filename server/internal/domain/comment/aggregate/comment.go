package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        string    `json:"id"`
	BugId     string    `json:"bugId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewComment(bugId, content string) Comment {
	return Comment{
		ID:        generateUniqueID(),
		BugId:     bugId,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Comment) EditComment(content string) {
	c.Content = content
	c.UpdatedAt = time.Now()
}

func generateUniqueID() string {
	return uuid.New().String()
}
