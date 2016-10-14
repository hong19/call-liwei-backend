CREATE TABLE legislators_committees(
	legislator_id 		INT 	REFERENCES legislators(id),
	committee_id 		INT 	REFERENCES committees(id),
	term				INT,
	session 			INT,
	onboard_date		TIMESTAMP
);