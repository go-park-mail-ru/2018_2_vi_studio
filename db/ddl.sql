CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

----------- users -----------

CREATE TABLE IF NOT EXISTS users (
  uid       UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  nickname CITEXT NOT NULL UNIQUE,
  email    CITEXT NOT NULL UNIQUE,
  password TEXT   NOT NULL,
  avatar   TEXT   NULL,
  points   INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS sessions (
  user_uid UUID REFERENCES users (uid) UNIQUE,
  token   UUID NOT NULL DEFAULT uuid_generate_v4()
);
