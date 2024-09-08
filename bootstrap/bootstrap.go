package bootstrap

import (
	"embed"
	"net/http"

	"github.com/YadBro/fiber-mvc-template/app/database"
	"github.com/YadBro/fiber-mvc-template/pkg/env"
	"github.com/YadBro/fiber-mvc-template/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
)

func NewApplication(viewsPath embed.FS, envScope env.Environment) *fiber.App {
	env.SetupEnvFile(envScope)
	database.ConnectDB()

	engine := django.NewPathForwardingFileSystem(http.FS(viewsPath), "/views", ".django")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/base_layout",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	if envScope == env.Development {
		app.Get("/monitor", monitor.New())
	}

	router.InstallRouter(app)

	return app
}
