CREATE TYPE transaction_status AS ENUM ('pending', 'success', 'failed');

CREATE TABLE transaction_history(
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  transaction_id UUID REFERENCES transactions(id),
  status        transaction_status DEFAULT 'pending',
  note          TEXT,
  created_at    TIMESTAMP DEFAULT now(),
  updated_at    TIMESTAMP DEFAULT now()
)