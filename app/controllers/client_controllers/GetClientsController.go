package client_controllers

import (
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
)

// GetClients func gets all exists clients.
// @Description Get all exists clients.
// @Summary get all exists clients
// @Tags Clients
// @Accept json
// @Produce json
// @Success 200 {array} models.Client
// @Router /v1/books [get]
func GetClients(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all clients.
	clients, err := db.GetClients()
	if err != nil {
		// Return, if clients not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "clients were not found",
			"count":   0,
			"clients": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(clients),
		"clients": clients,
	})
}
