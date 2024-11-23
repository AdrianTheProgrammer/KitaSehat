package service

import "KitaSehat_Backend/internal/features/surveys"

type SurveyService struct {
	qry surveys.SQuery
}

func NewSurveyService(q surveys.SQuery) surveys.SService {
	return &SurveyService{
		qry: q,
	}
}

func (ss *SurveyService) AddSurvey(input surveys.Survey) error {
	return nil
}

func (sq *SurveyService) GetAllSurveys() ([]surveys.Survey, error) {
	return []surveys.Survey{}, nil
}

func (sq *SurveyService) GetSurvey(id int) (surveys.Survey, error) {
	return surveys.Survey{}, nil
}

func (sq *SurveyService) UpdateSurvey(input surveys.Survey) error {
	return nil
}

func (sq *SurveyService) DeleteSurvey(id int) error {
	return nil
}
