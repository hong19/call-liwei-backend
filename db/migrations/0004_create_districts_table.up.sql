CREATE TABLE districts(
	id SERIAL PRIMARY KEY,
	name VARCHAR(60),
	electoral_district_id INT REFERENCES electoral_districts(id)
);