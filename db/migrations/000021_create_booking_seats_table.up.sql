CREATE TABLE booking_seats(
  id SERIAL PRIMARY KEY,
  booking_id INT REFERENCES bookings(id),
  seat_id INT REFERENCES seats(id),
  seat_price DECIMAL(10,2),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_booking_seats_seat_id ON booking_seats(seat_id);