package aggregate

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

type UserLoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserViewModel struct {
	Username string `json:"username"`
}

func NewUser(reg UserRegisterModel) *User {
	return &User{
		ID:        generateUniqueID(),
		Username:  reg.Username,
		Email:     reg.Email,
		Password:  hashPassword(reg.Password),
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
