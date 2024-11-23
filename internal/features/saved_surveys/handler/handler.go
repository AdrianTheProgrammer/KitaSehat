package handler

import (
	savedsurveys "KitaSehat_Backend/internal/features/saved_surveys"

	"github.com/labstack/echo/v4"
)

type SavedSurveyHandler struct {
	srv savedsurveys.SSService
}

func NewSavedSurveyHandler(s savedsurveys.SSService) savedsurveys.SSHandler {
	return &SavedSurveyHandler{
		srv: s,
	}
}

func (ssh *SavedSurveyHandler) SaveSurvey(c echo.Context) error {
	return nil
}

func (ssh *SavedSurveyHandler) GetAllSavedSurveys(c echo.Context) error {
	return nil
}

func (ssh *SavedSurveyHandler) GetSavedSurvey(c echo.Context) error {
	return nil
}

func (ssh *SavedSurveyHandler) DeleteSavedSurvey(c echo.Context) error {
	return nil
}
