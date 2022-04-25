package routes

import (
	"awesomeProject/app/controllers/client_controllers"
	"awesomeProject/app/controllers/transaction_controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/transactions", transaction_controllers.GetTransactions)   // get list of all transactions
	route.Get("/transaction/:id", transaction_controllers.GetTransaction) // get one transaction by ID
	route.Get("/clients", client_controllers.GetClients)                  // get list of all clients
	route.Get("/client/:id", client_controllers.GetClient)                // get one client by ID

	// Routes for POST method:
	route.Post("/transaction", transaction_controllers.CreateTransaction) // create a new transaction
	route.Post("/client", client_controllers.CreateClient)                // create a new client

	// Routes for PUT method:
	route.Put("/client", client_controllers.UpdateClient)                // update one book by ID
	route.Put("/transaction", transaction_controllers.UpdateTransaction) // update one transaction by ID
}
