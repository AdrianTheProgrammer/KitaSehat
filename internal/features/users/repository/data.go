package repository

import (
	saved "KitaSehat_Backend/internal/features/saved_surveys/repository"
	survey "KitaSehat_Backend/internal/features/surveys/repository"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	AvatarLink  string
	Survey      []survey.Survey     `gorm:"foreignKey:UserID"`
	SavedSurvey []saved.SavedSurvey `gorm:"foreignKey:UserID"`
}
