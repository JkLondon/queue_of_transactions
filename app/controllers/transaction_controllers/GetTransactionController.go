package transaction_controllers

import (
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetTransaction func gets transaction by given ID or 404 error.
// @Description Get transaction by given ID.
// @Summary get transaction by given ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Router /v1/transaction/{id} [get]
func GetTransaction(c *fiber.Ctx) error {
	// Catch transaction ID from URL.
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

	// Get transaction by ID.
	transaction, err := db.GetTransaction(id)
	if err != nil {
		// Return, if transaction not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":       true,
			"msg":         "transaction with the given ID is not found",
			"transaction": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":       false,
		"msg":         nil,
		"transaction": transaction,
	})
}
