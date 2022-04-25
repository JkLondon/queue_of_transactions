package queries

import (
	"awesomeProject/app/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// TransactionQueries struct for queries from Book model.
type TransactionQueries struct {
	*sqlx.DB
}

// GetTransactions method for getting all transactions.
func (q *TransactionQueries) GetTransactions() ([]models.Transaction, error) {
	// Define transactions variable.
	transactions := []models.Transaction{}

	// Define query string.
	query := `SELECT * FROM Transactions`

	// Send query to database.
	err := q.Select(&transactions, query)
	if err != nil {
		fmt.Println(err)
		// Return empty object and error.
		return transactions, err
	}

	// Return query result.
	return transactions, nil
}

// GetTransaction method for getting one transaction by given ID.
func (q *TransactionQueries) GetTransaction(id uuid.UUID) (models.Transaction, error) {
	// Define transaction variable.
	transaction := models.Transaction{}

	// Define query string.
	query := `SELECT * FROM Transactions WHERE id = $1`

	// Send query to database.
	err := q.Get(&transaction, query, id)
	if err != nil {
		// Return empty object and error.
		return transaction, err
	}

	// Return query result.
	return transaction, nil
}

// GetClientsTransactions method for getting all transactions by given client_id.
func (q *TransactionQueries) GetClientsTransactions(client_id uuid.UUID) ([]models.Transaction, error) {
	// Define transaction variable.
	transactions := []models.Transaction{}

	// Define query string.
	query := `SELECT * FROM Transactions WHERE client_id = $1`

	// Send query to database.
	err := q.Select(&transactions, query, client_id)
	if err != nil {
		// Return empty object and error.
		return transactions, err
	}

	// Return query result.
	return transactions, nil
}

// CreateTransaction method for creating transaction by given Transaction object.
func (q *TransactionQueries) CreateTransaction(b *models.Transaction) error {
	// Define query string.
	query := `INSERT INTO Transactions(id, created_at, transaction_type, status, amount, client_id) VALUES ($1, $2, $3, $4, $5, $6)`

	// Send query to database.
	_, err := q.Exec(query, b.Id, b.CreatedAt, b.TransactionType, b.Status, b.Amount, b.ClientId)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateTransaction method for updating transaction by given Transaction object.
func (q *TransactionQueries) UpdateTransaction(id uuid.UUID, b *models.Transaction) error {
	// Define query string.
	query := `UPDATE Transactions SET status = $2 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.Status)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
