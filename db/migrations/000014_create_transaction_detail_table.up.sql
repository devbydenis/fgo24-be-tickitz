CREATE TABLE transaction_detail(
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  transaction_id UUID REFERENCES transactions(id),
  seats          VARCHAR(255)
)