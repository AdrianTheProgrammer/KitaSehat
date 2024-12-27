package routes

import (
	ss_hnd "KitaSehat_Backend/internal/features/saved_surveys"
	s_hnd "KitaSehat_Backend/internal/features/surveys"
	u_hnd "KitaSehat_Backend/internal/features/users"
	"KitaSehat_Backend/internal/helper"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh u_hnd.UHandler, sh s_hnd.SHandler, ssh ss_hnd.SSHandler) {
	// Endpoint Test
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, helper.ResponseFormat(200, "endpoint test success", "hello world!", nil))
	})

	// Non JWT Endpoints
	e.POST("/register", uh.Register)
	e.POST("/login", uh.Login)

	// JWT Endpoints
	UserRoutes(e, uh)
}

func UserRoutes(e *echo.Echo, uh u_hnd.UHandler) {
	u := e.Group("/user")
	u.Use(JWTConfig())
	u.GET("", uh.GetAllUsers)
	u.GET("/:id", uh.GetUser)
	u.PUT("/:id", uh.UpdateUser)
	u.DELETE("/:id", uh.DeleteUser)
	u.PUT("/password/:id", uh.ChangePassword)
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(os.Getenv("jwtkey")),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
