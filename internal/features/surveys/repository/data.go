package repository

import "gorm.io/gorm"

type Survey struct {
	gorm.Model
	SurveyTitle string
	UserID      int
	Question    []Question `gorm:"foreignKey:SurveyID"`
}

type Question struct {
	gorm.Model
	SurveyID   int
	Order      int
	Question   string
	AnswerType string
	Answer     []Answer `gorm:"foreignKey:QuestionID"`
}

type Answer struct {
	gorm.Model
	QuestionID int
	Answer     string
}
