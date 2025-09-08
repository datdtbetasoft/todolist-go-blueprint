package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"my_project/internal/config"
	"my_project/internal/handler"
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

	// ==== Handler ====
	HealthHandler := handler.NewCheckHealthHandler()
	HelloHandler := handler.NewHelloWorldHandler()
	UserHandler := handler.NewUserHandler()
	AuthHandler := handler.NewAuthHandler()
	TaskHandler := handler.NewTaskHandler()

	version := e.Group("/api/" + c.GetString("server.version"))
	// ==== Routes ====
	api := e.Group("") // nếu sau này muốn versioning: e.Group("/api/v1")
	{
		api.GET("/", HelloHandler.HelloWorldHandler)
		api.GET("/health", HealthHandler.HealthHandler)
		api.POST("/register", UserHandler.Register) // thêm route đăng ký
		api.POST("/login", AuthHandler.Login)       // thêm route đăng ký

	}

	// Protected routes
	authGroup := version.Group("/auth")
	authGroup.Use(ApiMiddleware.TokenCheckMiddleware)
	{
		authGroup.POST("/task/create", TaskHandler.CreateATask)
		authGroup.PUT("/task/:id", TaskHandler.UpdateATask)
		authGroup.DELETE("/task/:id", TaskHandler.DeleteATask)
	}

	return e
}
