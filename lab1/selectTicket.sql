SELECT 
	ticket.id,
	movie.title as movie,
	hall.title as title,
	customer.first_name,
	price,
	movie.duration,
	row.number_in_hall,
	position.number_in_row
FROM ticket
JOIN session ON ticket.session_id = session.id
JOIN customer ON ticket.customer_id = customer.id
JOIN movie ON session.movie_id = movie.id
JOIN hall ON session.hall_id = hall.id
JOIN row ON ticket.row_id = row.id
JOIN position ON ticket.position_id = position.id