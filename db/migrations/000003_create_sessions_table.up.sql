CREATE TABLE sessions(
  id          UUID PRIMARY KEY,
  user_id     UUID NOT NULL REFERENCES users(id),
  token       VARCHAR(255) UNIQUE NOT NULL,
  device_info VARCHAR(255),
  is_active   BOOLEAN DEFAULT FALSE,
  created_at  TIMESTAMP DEFAULT now(),
  expired_at  TIMESTAMP DEFAULT now()
);