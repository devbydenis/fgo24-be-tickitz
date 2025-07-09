CREATE TYPE profile_gender AS ENUM ('male', 'female');

CREATE TABLE profiles(
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id     UUID NOT NULL UNIQUE REFERENCES users(id),
  username    VARCHAR(255) UNIQUE,
  firstname   VARCHAR(100),
  lastname    VARCHAR(100),
  birthday    DATE,
  gender      profile_gender,
  profile_picture VARCHAR(255),
  phone_number  VARCHAR(20),
  is_verified   BOOLEAN DEFAULT FALSE,
  created_at  TIMESTAMP DEFAULT now(),
  updated_at  TIMESTAMP DEFAULT now()
);