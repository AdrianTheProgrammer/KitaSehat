package handler

import (
	"KitaSehat_Backend/internal/features/users"
	"KitaSehat_Backend/internal/helper"
	"strconv"

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
	var input LoginRequest
	err := c.Bind(&input)
	if err != nil {
		c.Logger().Error("error binding login request body.", err.Error())
		return c.JSON(400, helper.ResponseFormat(400, "error binding login request body", nil, nil))
	}

	code, msg, token := uh.srv.Login(input.Email, input.Password)
	if code != 200 {
		return c.JSON(code, helper.ResponseFormat(code, msg, nil, nil))
	}

	return c.JSON(code, helper.ResponseFormat(code, msg, token, nil))
}

func (uh *UserHandler) Register(c echo.Context) error {
	var input RegisterRequest
	err := c.Bind(&input)
	if err != nil {
		c.Logger().Error("error binding register request body.", err.Error())
		return c.JSON(400, helper.ResponseFormat(400, "error binding register request body", nil, nil))
	}

	code, msg := uh.srv.Register(ToUserEntityRegister(input))
	return c.JSON(code, helper.ResponseFormat(code, msg, nil, nil))
}

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	currentPage, _ := strconv.Atoi(c.QueryParam("current_page"))

	code, msg, result, totalItems := uh.srv.GetAllUsers(currentPage)
	if code != 200 {
		return c.JSON(code, helper.ResponseFormat(code, msg, nil, nil))
	}

	return c.JSON(code, helper.ResponseFormat(code, msg, ToAllUsersResponse(result), helper.MetaResponse(currentPage, totalItems)))
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	code, msg, result := uh.srv.GetUser(id)
	if code != 200 {
		return c.JSON(code, helper.ResponseFormat(code, msg, nil, nil))
	}

	return c.JSON(code, helper.ResponseFormat(code, msg, ToUserResponse(result), nil))
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var input UpdateRequest
	err := c.Bind(&input)
	if err != nil {
		c.Logger().Error("error binding user update request body.", err.Error())
		return c.JSON(400, helper.ResponseFormat(400, "error binding user update request body", nil, nil))
	}

	code, msg := uh.srv.UpdateUser(c, id, ToUserEntityUpdate(input))
	return c.JSON(code, helper.ResponseFormat(code, msg, nil, nil))
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	code, msg := uh.srv.DeleteUser(id)
	return c.JSON(code, helper.ResponseFormat(code, msg, nil, nil))
}

func (uh *UserHandler) ChangePassword(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var input ChangePasswordRequest
	err := c.Bind(&input)
	if err != nil {
		c.Logger().Error("error binding change password request body.", err.Error())
		return c.JSON(400, helper.ResponseFormat(400, "error binding change password request body", nil, nil))
	}

	code, msg := uh.srv.ChangePassword(id, ToUserEntityChangePassword(input))
	return c.JSON(code, helper.ResponseFormat(code, msg, nil, nil))
}
