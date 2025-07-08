CREATE TABLE showtimes(
  id SERIAL PRIMARY KEY,
  movie_id INT REFERENCES movies(id),
  theater_id INT REFERENCES theaters(id),
  show_date TIMESTAMP,
  show_time TIMESTAMP,
  base_price DECIMAL(10,2),
  available_seats INT,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_showtimes_show_time ON showtimes(show_time);