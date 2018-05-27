CREATE DATABASE mydb;

\connect mydb;

CREATE TABLE users (
  id	SERIAL,
  name	text	NOT NULL,
  email	text	NOT NULL,
  created_at	timestamptz	NOT NULL,
  updated_at	timestamptz	NOT NULL
);


