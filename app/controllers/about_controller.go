package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func AboutIndex(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title":       "About",
		"CurrentTime": time.Now(),
	})
}
