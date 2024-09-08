package controllers

import (
	"log"
	"time"

	"github.com/YadBro/fiber-mvc-template/app/database"
	"github.com/YadBro/fiber-mvc-template/app/models"
	"github.com/gofiber/fiber/v2"
)

func getMenus(c *fiber.Ctx) ([]models.Menu, error) {
	var menus []models.Menu

	if database.DB == nil {
		log.Println("Database connection is nil")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Database connection error")
	}

	if err := database.DB.Preload("Children").Order("urutan ASC").Find(&menus).Error; err != nil {
		log.Println("Error retrieving menus:", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error retrieving menus")
	}

	return menus, nil
}

func HomeIndex(c *fiber.Ctx) error {
	menus, err := getMenus(c)
	if err != nil {
		log.Fatal("ERRROR")
	}
	return c.Render("pages/home", fiber.Map{
		"Title":       "Home",
		"CurrentTime": time.Now(),
		"Menus":       menus,
	})
}
