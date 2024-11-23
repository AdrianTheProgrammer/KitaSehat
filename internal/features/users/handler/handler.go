package handler

import (
	"KitaSehat_Backend/internal/features/users"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.UService
}

func NewUserHandler(s users.UService) users.UHandler {
	return &UserHandler{
		srv: s,
	}
}

func (uh *UserHandler) Login(c echo.Context) error {
	return nil
}

func (uh *UserHandler) Register(c echo.Context) error {
	return nil
}

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	return nil
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	return nil
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	return nil
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	return nil
}

func (uh *UserHandler) ChangePassword(c echo.Context) error {
	return nil
}
