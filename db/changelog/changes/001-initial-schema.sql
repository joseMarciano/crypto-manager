--liquibase formatted sql

--changeset crypto-manager:001-initial-schema
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL,
    document_number VARCHAR(255) NOT NULL
);

CREATE TABLE exchanges (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    minimum_age INTEGER NOT NULL,
    maximum_transfer_amount DECIMAL(15,2) NOT NULL
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    exchange_id UUID NOT NULL REFERENCES exchanges(id),
    balance DECIMAL(15,2) NOT NULL
);

CREATE TABLE exchange_transactions (
    id UUID PRIMARY KEY,
    exchange_id UUID NOT NULL REFERENCES exchanges(id),
    type VARCHAR(255) NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    executed_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE VIEW exchange_transactions_view AS
SELECT
    exchange_id,
    executed_at::DATE AS executed_at,
    type,
    SUM(amount) AS amount
FROM exchange_transactions
GROUP BY exchange_id, executed_at::DATE, type;
