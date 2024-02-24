package aggregate

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Projects  []string  `json:"projects"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterModel struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type UserViewModel struct {
	Username string `json:"username"`
}

func NewUser(reg UserRegisterModel) *User {
	hashedPassword, err := hashPassword(reg.Password)
	if err != nil {
		panic(err)
	}
	var emptyStringArray []string
	return &User{
		ID:        generateUniqueID(),
		Username:  reg.Username,
		Email:     reg.Email,
		Password:  hashedPassword,
		Projects:  emptyStringArray,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

func (u *User) UpdatePassword(password string) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	u.Password = hashedPassword
	u.UpdatedAt = time.Now()
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err

}

func generateUniqueID() string {
	return uuid.New().String()
}
