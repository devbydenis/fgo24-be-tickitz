CREATE TABLE theaters(
  id SERIAL PRIMARY KEY,
  cinema_id INT REFERENCES cinemas(id),
  name VARCHAR(255),
  capacity INT,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);