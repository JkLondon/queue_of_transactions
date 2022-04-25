package queries

import (
	"awesomeProject/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ClientQueries struct for queries from Client model.
type ClientQueries struct {
	*sqlx.DB
}

// GetClients method for getting all clients.
func (q *ClientQueries) GetClients() ([]models.Client, error) {
	// Define transactions variable.
	clients := []models.Client{}

	// Define query string.
	query := `SELECT * FROM Clients`

	// Send query to database.
	err := q.Select(&clients, query)
	if err != nil {
		// Return empty object and error.

		return clients, err
	}

	// Return query result.
	return clients, nil
}

// GetClient method for getting one client by given ID.
func (q *ClientQueries) GetClient(id uuid.UUID) (models.Client, error) {
	// Define client variable.
	client := models.Client{}

	// Define query string.
	query := `SELECT * FROM Clients WHERE id = $1`

	// Send query to database.
	err := q.Get(&client, query, id)
	if err != nil {
		// Return empty object and error.
		return client, err
	}

	// Return query result.
	return client, nil
}

// CreateClient method for creating client by given Client object.
func (q *ClientQueries) CreateClient(b *models.Client) error {
	// Define query string.
	query := `INSERT INTO Clients VALUES ($1, $2)`

	// Send query to database.
	_, err := q.Exec(query, b.Id, b.Balance)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateClient method for updating client by given Client object.
func (q *ClientQueries) UpdateClient(id uuid.UUID, b *models.Client) error {
	// Define query string.
	query := `UPDATE Clients SET balance = $2 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.Balance)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
