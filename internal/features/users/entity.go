package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserID      int
	Username    string
	Email       string
	PhoneNumber string
	AvatarLink  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	JWT         string
}

type UHandler interface {
	Login(echo.Context) error
	Register(echo.Context) error
	GetAllUsers(echo.Context) error
	GetUser(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
	ChangePassword(echo.Context) error
}

type UService interface {
	Login(string, string) (User, error)
	Register(User) error
	GetAllUsers() ([]User, error)
	GetUser(int) (User, error)
	UpdateUser(User) error
	DeleteUser(int) error
	ChangePassword(int, string) error
}

type UQuery interface {
	Login(string) (User, error)
	Register(User) error
	GetAllUsers() ([]User, error)
	GetUser(int) (User, error)
	UpdateUser(User) error
	DeleteUser(int) error
	ChangePassword(int, string) error
}

type LoginValidate struct {
}

type RegisterValidate struct {
}

type UpdateUserValidate struct {
}

type ChangePasswordValidate struct {
}
