CREATE TYPE payment_status AS ENUM ('pending', 'success', 'failed', 'refunded');

CREATE TABLE payments_transactions(
  id SERIAL PRIMARY KEY,
  payment_id INT REFERENCES payments(id),
  booking_id INT REFERENCES bookings(id),
  status payment_status DEFAULT 'pending',
  amount DECIMAL(10,2),
  completed_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT now(),
);

CREATE INDEX idx_payments_transactions_payment_id ON payments_transactions(payment_id);