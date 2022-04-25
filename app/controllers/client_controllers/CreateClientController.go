package client_controllers

import (
	"awesomeProject/app/models"
	"awesomeProject/pkg/utils"
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateClient func for creates a new client.
// @Description Create a new client.
// @Summary create a new client
// @Tags Client
// @Accept json
// @Produce json
// @Success 200 {object} models.Client
// @Router /v1/client [post]
func CreateClient(c *fiber.Ctx) error {

	// Create new Client struct
	client := &models.Client{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(client); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	// Create a new validator for a Client model.
	validate := utils.NewValidator()

	// Set initialized default data for client:
	client.Id = uuid.New()
	client.Balance = 10000

	// Validate client fields.
	if err := validate.Struct(client); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete client by given ID.
	if err := db.CreateClient(client); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"client": client,
	})
}
