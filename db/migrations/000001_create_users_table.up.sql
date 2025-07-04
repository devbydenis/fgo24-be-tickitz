-- CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users(
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email         VARCHAR(255) UNIQUE NOT NULL,
  password      VARCHAR(255) NOT NULL,
  role          user_role DEFAULT 'user',
  updated_at    TIMESTAMP DEFAULT now(),
  created_at    TIMESTAMP DEFAULT now()
);