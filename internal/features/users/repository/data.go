package repository

import (
	saved "KitaSehat_Backend/internal/features/saved_surveys/repository"
	survey "KitaSehat_Backend/internal/features/surveys/repository"
	"KitaSehat_Backend/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	Avatar      string
	AccessLevel string              `gorm:"default:'user'"`
	Survey      []survey.Survey     `gorm:"foreignKey:UserID"`
	SavedSurvey []saved.SavedSurvey `gorm:"foreignKey:UserID"`
}

func ToUserEntity(input User) users.User {
	return users.User{
		UserID:      int(input.ID),
		Username:    input.Username,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Avatar:      input.Avatar,
		AccessLevel: input.AccessLevel,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		DeletedAt:   input.DeletedAt.Time,
	}
}

func ToAllUserEntity(input []User) []users.User {
	var result []users.User
	for _, val := range input {
		result = append(result, users.User{
			UserID:      int(val.ID),
			Username:    val.Username,
			Email:       val.Email,
			Password:    val.Password,
			PhoneNumber: val.PhoneNumber,
			Avatar:      val.Avatar,
			AccessLevel: val.AccessLevel,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			DeletedAt:   val.DeletedAt.Time,
		})
	}
	return result
}

func ToUserData(input users.User) User {
	return User{
		Username:    input.Username,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Avatar:      input.Avatar,
		AccessLevel: input.AccessLevel,
	}
}
