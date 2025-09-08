package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"my_project/internal/config"
	"my_project/internal/controller"
	ApiMiddleware "my_project/internal/middleware"
	validator "my_project/internal/service"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	c := config.GetConfig()

	// ==== Middlewares ====
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	e.Validator = validator.NewValidator()

	// ==== Controllers ====
	HealthCtrl := controller.NewCheckHealthController()
	HelloCtrl := controller.NewHelloWorldController()
	UserCtrl := controller.NewUserController()
	AuthCtrl := controller.NewAuthController()
	TaskCtrl := controller.NewTaskController()

	version := e.Group("/api/" + c.GetString("server.version"))
	// ==== Routes ====
	api := e.Group("") // nếu sau này muốn versioning: e.Group("/api/v1")
	{
		api.GET("/", HelloCtrl.HelloWorldHandler)
		api.GET("/health", HealthCtrl.HealthHandler)
		api.POST("/register", UserCtrl.Register) // thêm route đăng ký
		api.POST("/login", AuthCtrl.Login)       // thêm route đăng ký

	}

	// Protected routes
	authGroup := version.Group("/auth")
	authGroup.Use(ApiMiddleware.TokenCheckMiddleware)
	{
		authGroup.POST("/task/create", TaskCtrl.CreateATask)
		authGroup.PUT("/task/:id", TaskCtrl.UpdateATask)
		authGroup.DELETE("/task/:id", TaskCtrl.DeleteATask)
	}

	return e
}
