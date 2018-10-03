CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

----------- users -----------

CREATE TABLE IF NOT EXISTS users (
  user_id       SERIAL PRIMARY KEY,
  nickname CITEXT NOT NULL UNIQUE,
  email    CITEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens (
  user_id       INTEGER REFERENCES users(user_id) UNIQUE ,
  uuid UUID NOT NULL DEFAULT uuid_generate_v4()
);
