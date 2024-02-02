package aggregate

import "github.com/google/uuid"

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     string `json:"ownerId"`
}

func NewProject(name, description, ownerId string) *Project {
	return &Project{
		ID:          generateUniqueID(),
		Name:        name,
		Description: description,
		OwnerId:     ownerId,
	}
}

func generateUniqueID() string {
	return uuid.New().String()
}
