package factory

import (
	"KitaSehat_Backend/configs"
	ss_hnd "KitaSehat_Backend/internal/features/saved_surveys/handler"
	ss_qry "KitaSehat_Backend/internal/features/saved_surveys/repository"
	ss_srv "KitaSehat_Backend/internal/features/saved_surveys/service"
	s_hnd "KitaSehat_Backend/internal/features/surveys/handler"
	s_qry "KitaSehat_Backend/internal/features/surveys/repository"
	s_srv "KitaSehat_Backend/internal/features/surveys/service"
	u_hnd "KitaSehat_Backend/internal/features/users/handler"
	u_qry "KitaSehat_Backend/internal/features/users/repository"
	u_srv "KitaSehat_Backend/internal/features/users/service"
	"KitaSehat_Backend/internal/routes"
	"KitaSehat_Backend/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	// Initialzie Database Connection
	db := configs.DBConnect()

	// Migrate Struct to Database
	db.AutoMigrate(
		&u_qry.User{},
		&s_qry.Survey{},
		&s_qry.Question{},
		&s_qry.Answer{},
		&ss_qry.SavedSurvey{},
		&ss_qry.SavedQuestion{},
	)

	// Declare Utility Methods
	pu := utils.NewPasswordUtility()
	tu := utils.NewTokenUtility()
	cu := utils.NewCloudinaryUtility()

	// Declare User Methods
	uq := u_qry.NewUserQuery(db)
	us := u_srv.NewUserService(uq, pu, tu, cu)
	uh := u_hnd.NewUserHandler(us)

	// Declare Survey Methods
	sq := s_qry.NewSurveyQuery(db)
	ss := s_srv.NewSurveyService(sq)
	sh := s_hnd.NewSurveyHandler(ss)

	// Declare Saved Survey Methods
	ssq := ss_qry.NewSavedSurveyQuery(db)
	sss := ss_srv.NewSavedSurveyService(ssq)
	ssh := ss_hnd.NewSavedSurveyHandler(sss)

	// Echo Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize API Endpoints
	routes.InitRoute(e, uh, sh, ssh)
}
