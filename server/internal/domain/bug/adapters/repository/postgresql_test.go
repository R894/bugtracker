package repository

import (
	"context"
	"testing"

	"github.com/R894/bugtracker/internal/domain/bug/core/domain/aggregate"

	"github.com/stretchr/testify/assert"
)

func TestSaveBug(t *testing.T) {
	repo := NewInMemoryBugRepository()
	ctx := context.Background()

	req := aggregate.CreateBugRequest{
		Title:       "Bug",
		Description: "Description",
		ProjectId:   "123123",
	}
	bug := aggregate.NewBug(req)

	err := repo.SaveBug(ctx, bug)
	assert.Nil(t, err)

	fetchedBug, err := repo.GetBugByID(ctx, bug.ID)
	assert.Nil(t, err)
	assert.Equal(t, fetchedBug.Title, req.Title)
	assert.Equal(t, fetchedBug.Description, req.Description)
	assert.Equal(t, fetchedBug.ProjectId, req.ProjectId)
}
