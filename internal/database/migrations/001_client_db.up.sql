-- client_db.sql

-- Criação da tabela clients
CREATE TABLE IF NOT EXISTS clients (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);