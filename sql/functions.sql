CREATE OR REPLACE FUNCTION avg_mark(id INTEGER)
RETURNS INTEGER
AS $$
  SELECT avg(mark) FROM visit WHERE location_id = id;
$$
LANGUAGE sql;