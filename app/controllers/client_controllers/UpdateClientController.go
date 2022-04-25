package client_controllers

import (
	"awesomeProject/app/models"
	"awesomeProject/pkg/utils"
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
)

// UpdateClient func for updates client by given ID.
// @Description Update client.
// @Summary update client
// @Tags Client
// @Accept json
// @Produce json
// @Success 201 {string} status "ok"
// @Router /v1/client [put]
func UpdateClient(c *fiber.Ctx) error {

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

	// Checking, if client with given ID is exists.
	foundedClient, err := db.GetClient(client.Id)
	if err != nil {
		// Return status 404 and client not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "client with this ID not found",
		})
	}

	// Create a new validator for a Transaction model.
	validate := utils.NewValidator()

	// Validate client fields.
	if err := validate.Struct(client); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update client by given ID.
	if err := db.UpdateClient(foundedClient.Id, client); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return c.SendStatus(fiber.StatusCreated)
}
