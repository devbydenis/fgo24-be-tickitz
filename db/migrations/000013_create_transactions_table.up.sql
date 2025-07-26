CREATE TABLE transactions(
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id       UUID REFERENCES users(id),
  cinema_id     INT REFERENCES cinemas(id),
  movie_id      INT REFERENCES movies(id),
  payment_methods_id    INT REFERENCES payment_methods(id),
  time_booking  TIME,
  date_booking  DATE,
  total_price   DECIMAL,
  created_at    TIMESTAMP DEFAULT now(),
  updated_at    TIMESTAMP DEFAULT now()
)