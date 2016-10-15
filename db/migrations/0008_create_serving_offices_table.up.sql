CREATE TABLE serving_offices(
	id			SERIAL 	PRIMARY KEY,
	address   	VARCHAR(255),
	telephone	VARCHAR(20),
	tax			VARCHAR(20),
	legislator_id INT REFERENCES legislators(id)
);