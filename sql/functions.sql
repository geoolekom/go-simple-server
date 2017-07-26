CREATE OR REPLACE FUNCTION avg_mark(location_id INTEGER)
RETURNS INTEGER
AS $$
  SELECT avg(mark) FROM visit WHERE "location" = location_id;
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION get_user(user_id INTEGER)
RETURNS "user"
AS $$
  SELECT id, email, first_name, last_name, gender, birth_date FROM "user" WHERE id = user_id;
$$
LANGUAGE sql
STABLE;
