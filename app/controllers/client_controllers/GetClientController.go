package client_controllers

import (
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetClient func gets client by given ID or 404 error.
// @Description Get client by given ID.
// @Summary get client by given ID
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Success 200 {object} models.Client
// @Router /v1/client/{id} [get]
func GetClient(c *fiber.Ctx) error {
	// Catch client ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get client by ID.
	client, err := db.GetClient(id)
	if err != nil {
		// Return, if client not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "client with the given ID is not found",
			"client": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"client": client,
	})
}
