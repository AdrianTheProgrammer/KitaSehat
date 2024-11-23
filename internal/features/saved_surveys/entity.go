package savedsurveys

import (
	"time"

	"github.com/labstack/echo/v4"
)

type SavedSurvey struct {
	SavedSurveyID int
	SurveyTitle   string
	Username      string
	Questions     []SavedQuestion
	AIConclusion  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type SavedQuestion struct {
	SavedQuestionID int
	SavedSurveyID   int
	Question        string
	Answer          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type SSHandler interface {
	SaveSurvey(echo.Context) error
	GetAllSavedSurveys(echo.Context) error
	GetSavedSurvey(echo.Context) error
	DeleteSavedSurvey(echo.Context) error
}

type SSService interface {
	SaveSurvey(SavedSurvey) error
	GetAllSavedSurveys(int) ([]SavedSurvey, error)
	GetSavedSurvey(int) (SavedSurvey, error)
	DeleteSavedSurvey(int) error
}

type SSQuery interface {
	SaveSurvey(SavedSurvey) error
	GetAllSavedSurveys(int) ([]SavedSurvey, error)
	GetSavedSurvey(int) (SavedSurvey, error)
	DeleteSavedSurvey(int) error
}
