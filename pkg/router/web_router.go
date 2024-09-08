package router

import (
	"github.com/YadBro/fiber-mvc-template/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type WebRouter struct{}

func (h WebRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New(), csrf.New())

	group.Get("/", controllers.HomeIndex)
}

func NewWebRouter() *WebRouter {
	return &WebRouter{}
}
