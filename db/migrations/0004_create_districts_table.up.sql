CREATE TABLE districts(
	id INT PRIMARY KEY,
	name VARCHAR(60),
	electoral_district_id INT REFERENCES electoral_districts(id)
);