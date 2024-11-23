package repository

import (
	"KitaSehat_Backend/internal/features/surveys"

	"gorm.io/gorm"
)

type SurveyQuery struct {
	db *gorm.DB
}

func NewSurveyQuery(DBConnect *gorm.DB) surveys.SQuery {
	return &SurveyQuery{
		db: DBConnect,
	}
}

func (sq *SurveyQuery) AddSurvey(input surveys.Survey) error {
	return nil
}

func (sq *SurveyQuery) GetAllSurveys() ([]surveys.Survey, error) {
	return []surveys.Survey{}, nil
}

func (sq *SurveyQuery) GetSurvey(id int) (surveys.Survey, error) {
	return surveys.Survey{}, nil
}

func (sq *SurveyQuery) UpdateSurvey(input surveys.Survey) error {
	return nil
}

func (sq *SurveyQuery) DeleteSurvey(id int) error {
	return nil
}
