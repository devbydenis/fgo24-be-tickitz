CREATE TABLE payments(
  id SERIAL PRIMARY KEY,
  method_name VARCHAR(255),
  provider VARCHAR(255),
  fee_process DECIMAL(10,2),
  is_active BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);
