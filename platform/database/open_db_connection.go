package database

import "awesomeProject/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.ClientQueries      // load queries from Client model
	*queries.TransactionQueries // load queries from Transaction model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		ClientQueries:      &queries.ClientQueries{DB: db}, // from Client model
		TransactionQueries: &queries.TransactionQueries{DB: db},
	}, nil
}
