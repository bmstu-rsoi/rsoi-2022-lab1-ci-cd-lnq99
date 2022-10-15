package server

import (
	"context"
	"os"
	"os/signal"
	"rsoi-1/config"
	"rsoi-1/internal/controller"
	"rsoi-1/internal/service"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	BaseServer
	echo *echo.Echo
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}

func NewEchoServer(services *service.Services, cfg *config.ServerConfig) Server {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1K"))

	//e.Logger.SetLevel(log.INFO)
	e.Validator = &CustomValidator{Validator: validator.New()}

	ctrl := controller.NewEchoController(services)
	RegisterEchoHandlers(e, ctrl)
	return &EchoServer{BaseServer{cfg}, e}
}

func RegisterEchoHandlers(e *echo.Echo, handlers *controller.EchoController) {
	e.GET("/api/v1/persons", handlers.ListPersons)
	e.POST("/api/v1/persons", handlers.CreatePerson)
	e.DELETE("/api/v1/persons/:id", handlers.DeletePerson)
	e.GET("/api/v1/persons/:id", handlers.GetPerson)
	e.PATCH("/api/v1/persons/:id", handlers.EditPerson)
}

func (s *EchoServer) Run() error {
	server := s.createHttpServer()
	e := s.echo

	go func() {
		e.Logger.Infof("EchoServer is listening on PORT: %s", s.cfg.Port)
		if err := e.StartServer(server); err != nil {
			e.Logger.Fatal("Error starting EchoServer: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	e.Logger.Info("EchoServer Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
