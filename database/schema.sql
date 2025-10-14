-- sql/schema.sql

CREATE TABLE users (
  id UUID PRIMARY KEY,
  avatar TEXT,
  first_name VARCHAR(20),
  last_name VARCHAR(20),
  username VARCHAR(20) UNIQUE NOT NULL,
  email VARCHAR(30) UNIQUE NOT NULL,
  birthday DATE,
  phone_number VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
