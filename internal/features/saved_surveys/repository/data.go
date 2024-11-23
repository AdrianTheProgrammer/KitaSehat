package repository

import "gorm.io/gorm"

type SavedSurvey struct {
	gorm.Model
	UserID        int
	SurveyTitle   string
	AIConlusion   string
	SavedQuestion []SavedQuestion `gorm:"foreignKey:SavedSurveyID"`
}

type SavedQuestion struct {
	gorm.Model
	SavedSurveyID int
	Question      string
	Answer        string
}
