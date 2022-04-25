package transaction_controllers

import (
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
)

// GetTransactions func gets all exists transactions.
// @Description Get all exists transactions.
// @Summary get all exists transactions
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {array} models.Transaction
// @Router /v1/transactions [get]
func GetTransactions(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all transactions.
	transactions, err := db.GetTransactions()
	if err != nil {
		// Return, if transactions not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":        true,
			"msg":          "transactions were not found",
			"count":        0,
			"transactions": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"count":        len(transactions),
		"transactions": transactions,
	})
}
