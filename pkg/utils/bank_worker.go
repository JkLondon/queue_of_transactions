package utils

import (
	"awesomeProject/platform/database"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func ApproveTransaction(TransactionId uuid.UUID) error {
	// Create database connection.
	fmt.Println("я начал аппрувить транзу")
	db, err := database.OpenDBConnection()
	if err != nil {
		return err
	}
	transaction, err := db.GetTransaction(TransactionId)
	if err != nil {
		return err
	}
	// Get client by ID.
	client, err := db.GetClient(transaction.ClientId)
	if err != nil {
		// Return, if client not found.
		return err
	}
	// bank is working
	//randomTime := rand.Intn(20)
	time.Sleep(30 * time.Second)
	// bank stop working

	if transaction.TransactionType == "+" {
		transaction.Status = "accepted"
		client.Balance += transaction.Amount
	} else if transaction.TransactionType == "-" {
		if client.Balance >= transaction.Amount {
			transaction.Status = "accepted"
			client.Balance -= transaction.Amount
		} else {
			transaction.Status = "blocked"
		}
	} else {
		transaction.Status = "blocked"
	}

	if err := db.UpdateTransaction(TransactionId, &transaction); err != nil {
		return err
	}
	if err := db.UpdateClient(client.Id, &client); err != nil {
		return err
	}
	fmt.Println("я кончил аппрувать транзу", client.Balance)
	return nil

}
