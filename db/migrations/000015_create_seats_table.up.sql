CREATE TABLE seats(
  id SERIAL PRIMARY KEY,
  theater_id INT REFERENCES theaters(id),
  seat_letter VARCHAR(1),
  seat_number INT,
  is_active BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);