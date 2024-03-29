package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	bugId := "123123"
	content := "hello"
	userId := "user"
	comment := NewComment(bugId, content, userId)

	assert.Equal(t, comment.BugId, bugId)
	assert.Equal(t, comment.Content, content)
	assert.NotEmpty(t, comment.ID)
	assert.NotEmpty(t, comment.BugId)
	assert.NotZero(t, comment.CreatedAt)
	assert.NotZero(t, comment.UpdatedAt)
}

func TestEditComment(t *testing.T) {
	bugId := "123123"
	content := "hello"
	userId := "user"
	comment := NewComment(bugId, content, userId)
	updatedContent := "test"
	comment.EditComment(updatedContent)

	assert.Equal(t, comment.BugId, bugId)
	assert.Equal(t, comment.Content, updatedContent)
	assert.NotEqual(t, comment.Content, content)
	assert.NotEmpty(t, comment.ID)
	assert.NotEmpty(t, comment.BugId)
	assert.NotZero(t, comment.CreatedAt)
	assert.NotZero(t, comment.UpdatedAt)
}
