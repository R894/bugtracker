package repository

import (
	config "bugtracker"
	"bugtracker/internal/database"
	"bugtracker/internal/domain/bug/core/domain/aggregate"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveBug(t *testing.T) {
	config.FetchEnvVars()
	db, err := database.NewPostgresTestDB()
	if err != nil {
		log.Fatal("Error setting up test database: ", err)
	}
	repo := NewPostgresBugRepository(db.Db)
	ctx := context.Background()

	req := aggregate.CreateBugRequest{
		Title:       "Bug",
		Description: "Description",
		ProjectId:   "123123",
	}
	bug := aggregate.NewBug(req)

	err = repo.SaveBug(ctx, bug)
	assert.Nil(t, err)

	fetchedBug, err := repo.GetBugByID(ctx, bug.ID)
	assert.Nil(t, err)
	assert.Equal(t, fetchedBug.Title, req.Title)
	assert.Equal(t, fetchedBug.Description, req.Description)
	assert.Equal(t, fetchedBug.ProjectId, req.ProjectId)
}
