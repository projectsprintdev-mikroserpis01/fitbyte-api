package server

import (
	"path/filepath"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/internal/infra/env"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/internal/middlewares"
	errorhandler "github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/helpers/http/error_handler"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/helpers/http/response"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/jwt"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/log"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/s3"
	timePkg "github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/time"
)

type HttpServer interface {
	Start(part string)
	MountMiddlewares()
	MountRoutes(db *sqlx.DB)
	GetApp() *fiber.App
}

type httpServer struct {
	app *fiber.App
}

func NewHttpServer() HttpServer {
	config := fiber.Config{
		CaseSensitive: true,
		AppName:       "GoGo Manager",
		ServerHeader:  "GoGo Manager",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		ErrorHandler:  errorhandler.ErrorHandler,
	}

	app := fiber.New(config)

	return &httpServer{
		app: app,
	}
}

func (s *httpServer) GetApp() *fiber.App {
	return s.app
}

func (s *httpServer) Start(port string) {
	if port[0] != ':' {
		port = ":" + port
	}

	err := s.app.Listen(port)

	if err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "[SERVER][Start] failed to start server")
	}
}

func (s *httpServer) MountMiddlewares() {
	s.app.Use(middlewares.LoggerConfig())
	s.app.Use(middlewares.Helmet())
	s.app.Use(middlewares.Compress())
	s.app.Use(middlewares.Cors())
	if env.AppEnv.AppEnv != "development" {
		s.app.Use(middlewares.ApiKey())
	}
	s.app.Use(middlewares.RecoverConfig())
}

func (s *httpServer) MountRoutes(db *sqlx.DB) {
	//bcrypt := bcrypt.Bcrypt
	_ = timePkg.Time
	//uuid := uuid.UUID
	//validator := validator.Validator
	jwtManager := jwt.JwtManager
	jwt := jwt.Jwt
	s3 := s3.S3

	middleware := middlewares.NewMiddleware(jwt, jwtManager)

	s.app.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "fitbyte API")
	})

	api := s.app.Group("/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "fitbyte API")
	})

	s.app.Post("/v1/file", middleware.RequireAdmin(), func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return domain.ErrFileNotFound
		}

		extFileOptions := []string{"jpg", "jpeg", "png"}
		maxSize := 100 * 1024 // 100 KiB

		// check file extension
		validExt := false
		for _, ext := range extFileOptions {
			if strings.Contains(filepath.Ext(file.Filename), ext) {
				validExt = true
				break
			}
		}

		if !validExt {
			return domain.ErrInvalidFileExtension
		}

		// check file size
		if file.Size > int64(maxSize) {
			return domain.ErrFileSizeLimitExceeded
		}

		uri, err := s3.Upload(file)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"uri": uri,
		})
	})

	s.app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("./web/not-found.html")
	})
}
