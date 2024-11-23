package repository

import (
	savedsurveys "KitaSehat_Backend/internal/features/saved_surveys"

	"gorm.io/gorm"
)

type SavedSurveyQuery struct {
	db *gorm.DB
}

func NewSavedSurveyQuery(DBConnect *gorm.DB) savedsurveys.SSQuery {
	return &SavedSurveyQuery{
		db: DBConnect,
	}
}

func (ssq *SavedSurveyQuery) SaveSurvey(input savedsurveys.SavedSurvey) error {
	return nil
}

func (ssq *SavedSurveyQuery) GetAllSavedSurveys(userID int) ([]savedsurveys.SavedSurvey, error) {
	return []savedsurveys.SavedSurvey{}, nil
}

func (ssq *SavedSurveyQuery) GetSavedSurvey(id int) (savedsurveys.SavedSurvey, error) {
	return savedsurveys.SavedSurvey{}, nil
}

func (ssq *SavedSurveyQuery) DeleteSavedSurvey(id int) error {
	return nil
}
