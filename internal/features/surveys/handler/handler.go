package handler

import (
	"KitaSehat_Backend/internal/features/surveys"

	"github.com/labstack/echo/v4"
)

type SurveyHandler struct {
	srv surveys.SService
}

func NewServiceHandler(s surveys.SService) surveys.SHandler {
	return &SurveyHandler{
		srv: s,
	}
}

func (sh *SurveyHandler) AddSurvey(c echo.Context) error {
	return nil
}

func (sh *SurveyHandler) GetAllSurveys(c echo.Context) error {
	return nil
}

func (sh *SurveyHandler) GetSurvey(c echo.Context) error {
	return nil
}

func (sh *SurveyHandler) UpdateSurvey(c echo.Context) error {
	return nil
}

func (sh *SurveyHandler) DeleteSurvey(c echo.Context) error {
	return nil
}
