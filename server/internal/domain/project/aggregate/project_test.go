package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	name := "name"
	desc := "description"
	ownerId := "ownerid"
	proj := NewProject(name, desc, ownerId)

	assert.Equal(t, proj.Name, name)
	assert.Equal(t, proj.Description, desc)
	assert.Equal(t, proj.OwnerId, ownerId)
	assert.NotNil(t, proj.ID)
	assert.NotNil(t, proj.CreatedAt)
	assert.NotNil(t, proj.UpdatedAt)
}
