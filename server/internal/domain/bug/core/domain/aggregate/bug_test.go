package aggregate

import (
	"testing"

	"github.com/R894/bugtracker/internal/domain/bug/core/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewBug(t *testing.T) {
	bugReq := CreateBugRequest{
		Title:       "test",
		Description: "test",
	}

	bug := NewBug(bugReq)

	assert.NotNil(t, bug)
	assert.NotEmpty(t, bug.Status)
	assert.Equal(t, bugReq.Title, bug.Title)
	assert.Equal(t, bugReq.Description, bug.Description)
	assert.NotZero(t, bug.CreatedAt)
	assert.NotZero(t, bug.UpdatedAt)
}

func TestUpdateDetails(t *testing.T) {
	bugReq := CreateBugRequest{
		Title:       "test",
		Description: "test",
	}

	bug := NewBug(bugReq)

	bug.UpdateDetails("test2", "test2")

	assert.NotEqual(t, bugReq.Title, bug.Title)
	assert.NotEqual(t, bugReq.Description, bug.Description)
	assert.NotZero(t, bug.CreatedAt)
	assert.NotZero(t, bug.UpdatedAt)
}

func TestAssignTo(t *testing.T) {
	bugReq := CreateBugRequest{
		Title:       "test",
		Description: "test",
	}

	bug := NewBug(bugReq)

	assignee := "John Doe"
	bug.AssignTo(assignee)

	assert.Equal(t, assignee, bug.Assignee)
	assert.NotZero(t, bug.UpdatedAt)
}

func TestChangeStatus(t *testing.T) {
	bugReq := CreateBugRequest{
		Title:       "test",
		Description: "test",
	}

	bug := NewBug(bugReq)

	status := entity.BugStatusClosed
	bug.ChangeStatus(status)

	assert.Equal(t, status, bug.Status)
	assert.NotZero(t, bug.UpdatedAt)
}

func TestUpdateDetailsWithSameValues(t *testing.T) {
	bugReq := CreateBugRequest{
		Title:       "test",
		Description: "test",
	}

	bug := NewBug(bugReq)

	bug.UpdateDetails("test", "test")

	assert.Equal(t, bugReq.Title, bug.Title)
	assert.Equal(t, bugReq.Description, bug.Description)
	assert.Equal(t, bug.CreatedAt, bug.UpdatedAt)
}

func TestUpdateDetailsWithEmptyValues(t *testing.T) {
	bugReq := CreateBugRequest{
		Title:       "test",
		Description: "test",
	}

	bug := NewBug(bugReq)

	bug.UpdateDetails("", "")

	assert.Equal(t, bugReq.Title, bug.Title)
	assert.Equal(t, bugReq.Description, bug.Description)
	assert.Equal(t, bug.CreatedAt, bug.UpdatedAt)
}
