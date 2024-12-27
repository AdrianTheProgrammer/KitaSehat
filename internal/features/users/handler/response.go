package handler

import (
	"KitaSehat_Backend/internal/features/users"
	"time"
)

type UserResponse struct {
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func ToAllUsersResponse(input []users.User) []UserResponse {
	var result []UserResponse

	for _, val := range input {
		result = append(result, UserResponse{
			UserID:      val.UserID,
			Username:    val.Username,
			Email:       val.Email,
			PhoneNumber: val.PhoneNumber,
			Avatar:      val.Avatar,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			DeletedAt:   val.DeletedAt,
		})
	}

	return result
}

func ToUserResponse(input users.User) UserResponse {
	return UserResponse{
		UserID:      input.UserID,
		Username:    input.Username,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Avatar:      input.Avatar,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		DeletedAt:   input.DeletedAt,
	}
}
