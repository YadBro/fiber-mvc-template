package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func HomeIndex(c *fiber.Ctx) error {
	return c.Render("pages/home", fiber.Map{
		"Title":       "Home",
		"CurrentTime": time.Now(),
	})
}
