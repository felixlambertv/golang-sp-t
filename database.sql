CREATE TABLE users (
	id serial PRIMARY KEY,
	phone_number VARCHAR (13) UNIQUE NOT NULL,
	full_name VARCHAR (60) NOT NULL,
	password VARCHAR (100) NOT NUll,
    login_count integer DEFAULT 0
);
