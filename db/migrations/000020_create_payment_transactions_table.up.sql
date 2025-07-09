CREATE TABLE payments_transactions(
  id SERIAL PRIMARY KEY,
  payment_id INT REFERENCES payments(id),
  booking_id INT REFERENCES bookings(id),
  status payment_status DEFAULT 'pending',
  amount DECIMAL(10,2),
  completed_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT now()
);
