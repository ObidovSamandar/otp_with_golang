CREATE TABLE users_db (
	id varchar(80) PRIMARY KEY,
	first_name varchar(80) UNIQUE,
	phone_number varchar(80) UNIQUE,
	user_img varchar(40) DEFAULT NULL
)