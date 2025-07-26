CREATE TABLE cinemas(
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(255),
  created_at    TIMESTAMP DEFAULT now(),
  updated_at    TIMESTAMP DEFAULT now()
)