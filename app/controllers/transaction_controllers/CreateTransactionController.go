package transaction_controllers

import (
	"awesomeProject/app/models"
	"awesomeProject/pkg/utils"
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

// CreateTransaction func for creates a new transaction.
// @Description Create a new transaction.
// @Summary create a new transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.Transaction
// @Router /v1/transaction [post]
func CreateTransaction(c *fiber.Ctx) error {

	// Create new Transaction struct
	transaction := &models.Transaction{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(transaction); err != nil {
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

	// Create a new validator for a Transaction model.
	validate := utils.NewValidator()

	// Set initialized default data for transaction:
	transaction.Id = uuid.New()
	transaction.CreatedAt = time.Now()
	transaction.Status = "pending" // pending / accepted / blocked

	// Validate transaction fields.
	if err := validate.Struct(transaction); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete transaction by given ID.
	if err := db.CreateTransaction(transaction); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	utils.Cm[transaction.ClientId] <- transaction.Id
	/*
		if err := utils.ApproveTransaction(transaction.Id); err != nil {
			// Return status 500 and error message.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		} */

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":       false,
		"msg":         nil,
		"transaction": transaction,
	})
}
