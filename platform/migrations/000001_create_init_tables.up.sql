-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Moscow";

CREATE TABLE Clients(
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    balance INT NOT NULL
);

CREATE TABLE Transactions(
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    transaction_type VARCHAR (255) NOT NULL,
    status VARCHAR (255) NOT NULL,
    amount INT NOT NULL,
    client_id UUID,
    FOREIGN KEY (client_id) REFERENCES Clients (id) ON DELETE CASCADE
);
