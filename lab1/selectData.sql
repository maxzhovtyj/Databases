--- SELECT and JOIN to get all tickets 
SELECT 
	ticket.id,
	movie.title as movie,
	session.start_at,
	hall.title as hall,
	customer.first_name,
	price,
	movie.duration,
	row.number_in_hall as row,
	position.number_in_row as position
FROM ticket
JOIN session ON ticket.session_id = session.id
JOIN customer ON ticket.customer_id = customer.id
JOIN movie ON session.movie_id = movie.id
JOIN hall ON session.hall_id = hall.id
JOIN row ON ticket.row_id = row.id
JOIN position ON ticket.position_id = position.id;

--- SELECT all customers
SELECT * FROM customer;

--- SELECT all movies
SELECT * FROM movie;

--- SELECT all halls;
SELECT * FROM hall;

--- SELECT all existings row in all halls
SELECT 
	row.id,
	row.number_in_hall,
	row.hall_id,
	hall.title as hall 
FROM row
JOIN hall ON row.hall_id = hall.id;

--- SELECT all existings position in all rows
SELECT * FROM position;

SELECT 
	position.id,
	row.number_in_hall as row,
	position.number_in_row as position,
	hall.title as hall
FROM position
JOIN row ON position.row_id = row.id
JOIN hall ON row.hall_id = hall.id;

--- SELECT all sessions
SELECT 
	session.id,
	session.movie_id,
	session.hall_id
FROM session;
