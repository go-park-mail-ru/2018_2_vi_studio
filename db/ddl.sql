CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

----------- users -----------

CREATE TABLE IF NOT EXISTS users (
  id  SERIAL PRIMARY KEY,
  nickname CITEXT NOT NULL UNIQUE,
  email    CITEXT NOT NULL UNIQUE,
  password TEXT   NOT NULL,
  points   INTEGER
);

CREATE TABLE IF NOT EXISTS sessions (
  user_id INTEGER REFERENCES users (id) UNIQUE,
  token   UUID NOT NULL DEFAULT uuid_generate_v4()
);
