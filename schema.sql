CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users(
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username      VARCHAR(255) UNIQUE,
  email         VARCHAR(255) UNIQUE NOT NULL,
  password      VARCHAR(255) NOT NULL,
  role          user_role DEFAULT 'user',
  updated_at    TIMESTAMP DEFAULT now(),
  created_at    TIMESTAMP DEFAULT now()
);

CREATE TABLE sessions(
  id          UUID PRIMARY KEY,
  user_id     UUID NOT NULL REFERENCES users(id),
  token       VARCHAR(255) UNIQUE NOT NULL,
  device_info VARCHAR(255),
  is_active   BOOLEAN DEFAULT FALSE,
  created_at  TIMESTAMP DEFAULT now(),
  expired_at  TIMESTAMP DEFAULT now()
);
CREATE TABLE profiles(
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id     UUID NOT NULL UNIQUE REFERENCES users(id),
  firstname   VARCHAR(100),
  lastname    VARCHAR(100),
  birthday    DATE,
  gender      VARCHAR(10),
  profile_picture VARCHAR(255),
  phone_number  VARCHAR(20),
  is_verified   BOOLEAN DEFAULT FALSE,
  created_at  TIMESTAMP DEFAULT now(),
  updated_at  TIMESTAMP DEFAULT now()
);

DROP TABLE profiles;
DROP TABLE sessions;
DROP TABLE users;