-- init.sql

CREATE DATABASE dbwallet;

\c dbwallet;

CREATE TABLE IF NOT EXISTS clients (
    client_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255),
    zipcode VARCHAR(10),
    status VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS wallets (
    wallet_id SERIAL PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL,
    wallet_type VARCHAR(255) NOT NULL,
    amount NUMERIC(15, 2) NOT NULL
);

CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    wallet_id VARCHAR(255) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
