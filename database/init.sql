CREATE DATABASE mydb;

\connect mydb;

CREATE TABLE users (
  id	integer,
  name	text,
  email	text,
  created_at	timestamptz,
  updated_at	timestamptz
);


