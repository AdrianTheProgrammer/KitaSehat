package handler

import "KitaSehat_Backend/internal/features/users"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateRequest struct {
	Username    string `form:"username"`
	PhoneNumber string `form:"phone_number"`
	Avatar      string `form:"avatar"`
	AccessLevel string `form:"access_level"`
}

type ChangePasswordRequest struct {
	Password string `json:"password"`
}

func ToUserEntityRegister(input RegisterRequest) users.User {
	return users.User{
		Username:    input.Username,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
}

func ToUserEntityUpdate(input UpdateRequest) users.User {
	return users.User{
		Username:    input.Username,
		PhoneNumber: input.PhoneNumber,
		Avatar:      input.Avatar,
		AccessLevel: input.AccessLevel,
	}
}

func ToUserEntityChangePassword(input ChangePasswordRequest) users.User {
	return users.User{
		Password: input.Password,
	}
}
