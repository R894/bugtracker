package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(username, email, password string) *User {
	return &User{
		ID:        generateUniqueID(),
		Username:  username,
		Email:     email,
		Password:  hashPassword(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

func (u *User) UpdatePassword(password string) {
	u.Password = hashPassword(password)
	u.UpdatedAt = time.Now()
}

func hashPassword(password string) string {
	// TODO: implement hashing
	return password
}

func generateUniqueID() string {
	return uuid.New().String()
}
