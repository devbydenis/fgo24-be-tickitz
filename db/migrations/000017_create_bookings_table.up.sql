CREATE TABLE bookings(
  id SERIAL PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  showtime_id INT REFERENCES showtimes(id),
  status  booking_status DEFAULT 'pending',
  total_amount DECIMAL(10,2),
  discount_amount DECIMAL(10,2),
  tax_amount DECIMAL(10,2),
  booking_time TIMESTAMP,
  expires_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);
