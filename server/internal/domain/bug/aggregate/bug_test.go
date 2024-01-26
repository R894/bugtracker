package aggregate

import (
	"testing"

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
