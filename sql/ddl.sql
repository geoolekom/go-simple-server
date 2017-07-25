--- CREATE TYPE GENDER AS ENUM ('m', 'f');

CREATE TABLE IF NOT EXISTS "user" (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) UNIQUE NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  gender GENDER NOT NULL,
  birth_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS location (
  id SERIAL PRIMARY KEY,
  place TEXT NOT NULL,
  country VARCHAR(50) NOT NULL,
  city VARCHAR(50) NOT NULL,
  distance INTEGER CHECK (distance > 0)
);

CREATE TABLE IF NOT EXISTS visit (
  id SERIAL PRIMARY KEY,
  "user" INTEGER REFERENCES "user" (id),
  location INTEGER REFERENCES location (id),
  visited_at TIMESTAMP NOT NULL,
  mark INTEGER CHECK (mark > 0 AND mark <= 5)
);
