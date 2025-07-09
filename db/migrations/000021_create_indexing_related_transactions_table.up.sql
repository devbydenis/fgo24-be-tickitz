CREATE INDEX idx_cities_name ON cities(name);
CREATE INDEX idx_cinemas_name ON cinemas(name);
CREATE INDEX idx_seats_seat_number ON seats(seat_number);
CREATE INDEX idx_bookings_showtime_id ON bookings(showtime_id);
CREATE INDEX idx_payments_transactions_payment_id ON payments_transactions(payment_id);
