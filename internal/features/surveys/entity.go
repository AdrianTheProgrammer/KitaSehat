package surveys

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Survey struct {
	SurveyID    int
	SurveyTitle string
	CreatedBy   string
	Questions   []Question
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Question struct {
	QuestionID int
	Order      int
	Question   string
	AnswerType string
	Answers    []Answer
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type Answer struct {
	AnswerID  int
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type SHandler interface {
	AddSurvey(echo.Context) error
	GetAllSurveys(echo.Context) error
	GetSurvey(echo.Context) error
	UpdateSurvey(echo.Context) error
	DeleteSurvey(echo.Context) error
}

type SService interface {
	AddSurvey(Survey) error
	GetAllSurveys() ([]Survey, error)
	GetSurvey(int) (Survey, error)
	UpdateSurvey(Survey) error
	DeleteSurvey(int) error
}

type SQuery interface {
	AddSurvey(Survey) error
	GetAllSurveys() ([]Survey, error)
	GetSurvey(int) (Survey, error)
	UpdateSurvey(Survey) error
	DeleteSurvey(int) error
}
