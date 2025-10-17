SET TIME ZONE 'Asia/Ho_Chi_Minh';

CREATE TABLE users (
  id UUID PRIMARY KEY,
  avatar TEXT,
  birthday DATE,
  phone_number VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
