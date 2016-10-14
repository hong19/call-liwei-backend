CREATE TABLE legislators(
	id 				INT 			PRIMARY KEY NOT NULL,
	name 			VARCHAR(60)		NOT NULL,
	english_name 	VARCHAR(60),
	gender 			CHAR(1)			NOT NULL,
	email VARCHAR(255),
	experience VARCHAR(1023),
	facebook VARCHAR(255),
	wiki VARCHAR(511),
	line VARCHAR(255),
	party_id INT REFERENCES parties(id),
	party_group_id INt REFERENCES party_groups(id),
	electoral_district_id INT REFERENCES electoral_districts(id)
);