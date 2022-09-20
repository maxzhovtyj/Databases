DROP TABLE IF EXISTS customer;
CREATE TABLE customer
(
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(128) NOT NULL,
	last_name VARCHAR(128) NOT NULL,
	UNIQUE(first_name, last_name)
);

DROP TABLE IF EXISTS movie;
CREATE TABLE movie
(
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL UNIQUE,
	description TEXT,
	duration INTERVAL NOT NULL
);

DROP TABLE IF EXISTS hall;
CREATE TABLE hall
(
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL UNIQUE,
	description TEXT,
	capacity INT
);

DROP TABLE IF EXISTS session;
CREATE TABLE session
(
	id SERIAL PRIMARY KEY,
	movie_id INT REFERENCES movie (id) ON DELETE CASCADE NOT NULL,
	hall_id INT REFERENCES hall (id) ON DELETE CASCADE NOT NULL,
	start_at TIMESTAMPTZ,
	UNIQUE(hall_id, start_at)
);

DROP TABLE IF EXISTS row;
CREATE TABLE row
(
	id SERIAL PRIMARY KEY,
	hall_id INT REFERENCES hall (id) ON DELETE CASCADE NOT NULL,
	number_in_hall INT NOT NULL,
	UNIQUE(hall_id, number_in_hall)
);

DROP TABLE IF EXISTS position;
CREATE TABLE position
(
	id SERIAL PRIMARY KEY,
	row_id INT REFERENCES row (id) ON DELETE CASCADE NOT NULL,
	number_in_row INT NOT NULL,
	UNIQUE(row_id, number_in_row)
);

DROP TABLE IF EXISTS ticket;
CREATE TABLE ticket
(
	id SERIAL PRIMARY KEY,
	customer_id INT REFERENCES customer (id) ON DELETE CASCADE NOT NULL,
	session_id INT REFERENCES session (id) ON DELETE CASCADE NOT NULL,
	price DECIMAL(12, 2) NOT NULL,
	row_id INT REFERENCES row (id) ON DELETE CASCADE NOT NULL,
	position_id INT REFERENCES position (id) ON DELETE CASCADE NOT NULL,
	UNIQUE(session_id, row_id, position_id)
);