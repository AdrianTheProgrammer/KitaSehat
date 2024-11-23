package service

import savedsurveys "KitaSehat_Backend/internal/features/saved_surveys"

type SavedSurveyService struct {
	qry savedsurveys.SSQuery
}

func NewSavedSurveyService(q savedsurveys.SSQuery) savedsurveys.SSService {
	return &SavedSurveyService{
		qry: q,
	}
}

func (sss *SavedSurveyService) SaveSurvey(input savedsurveys.SavedSurvey) error {
	return nil
}

func (sss *SavedSurveyService) GetAllSavedSurveys(userID int) ([]savedsurveys.SavedSurvey, error) {
	return []savedsurveys.SavedSurvey{}, nil
}

func (sss *SavedSurveyService) GetSavedSurvey(id int) (savedsurveys.SavedSurvey, error) {
	return savedsurveys.SavedSurvey{}, nil
}

func (sss *SavedSurveyService) DeleteSavedSurvey(id int) error {
	return nil
}
