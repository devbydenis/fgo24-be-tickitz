CREATE TABLE cities(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  province VARCHAR(255),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

