package transaction_controllers

import (
	"awesomeProject/app/models"
	"awesomeProject/pkg/utils"
	"awesomeProject/platform/database"
	"github.com/gofiber/fiber/v2"
)

// UpdateTransaction func for updates transaction by given ID.
// @Description Update transaction.
// @Summary update transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 201 {string} status "ok"
// @Router /v1/transaction [put]
func UpdateTransaction(c *fiber.Ctx) error {

	// Create new Book struct
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

	// Checking, if book with given ID is exists.
	foundedTransaction, err := db.GetTransaction(transaction.Id)
	if err != nil {
		// Return status 404 and transaction not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "transaction with this ID not found",
		})
	}

	// Create a new validator for a Transaction model.
	validate := utils.NewValidator()

	// Validate transaction fields.
	if err := validate.Struct(transaction); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update transaction by given ID.
	if err := db.UpdateTransaction(foundedTransaction.Id, transaction); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return c.SendStatus(fiber.StatusCreated)
}
