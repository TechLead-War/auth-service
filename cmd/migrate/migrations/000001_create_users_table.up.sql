CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password BYTEA NOT NULL, -- "byte array" is a data type in PostgreSQL used to store binary data. It can store any kind of binary data, such as images, files, or encrypted data.
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email citext UNIQUE NOT NULL -- case-insensitive
);
