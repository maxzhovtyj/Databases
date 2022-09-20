INSERT INTO customer (first_name, last_name) values ('Максим', 'Жовтанюк');
INSERT INTO customer (first_name, last_name) values ('Володимир', 'Зеленський');
INSERT INTO customer (first_name, last_name) values ('Петро', 'Порошенко');

INSERT INTO movie (title, description, duration) values ('Shawshank Redemption', 'Movie about prison escape', '2:30:25');
INSERT INTO movie (title, description, duration) values ('The Lord of the Rings', 'Movie about hobbits', '3:45:40');
INSERT INTO movie (title, description, duration) values ('Interstellar', 'Movie about space', '2:40:00');

INSERT INTO hall (title, description) values ('Paradise', 'Good hall for good films');
INSERT INTO hall (title, description) values ('Second hall', 'Some good hall');

INSERT INTO row (hall_id, number_in_hall) values (1, 1);
INSERT INTO row (hall_id, number_in_hall) values (1, 2);
INSERT INTO row (hall_id, number_in_hall) values (1, 3);

INSERT INTO row (hall_id, number_in_hall) values (2, 1);
INSERT INTO row (hall_id, number_in_hall) values (2, 2);

INSERT INTO position (row_id, number_in_row) values (1, 1);
INSERT INTO position (row_id, number_in_row) values (1, 2);
INSERT INTO position (row_id, number_in_row) values (1, 3);
INSERT INTO position (row_id, number_in_row) values (1, 4);

INSERT INTO position (row_id, number_in_row) values (2, 1);
INSERT INTO position (row_id, number_in_row) values (2, 2);
INSERT INTO position (row_id, number_in_row) values (2, 3);
INSERT INTO position (row_id, number_in_row) values (2, 4);

INSERT INTO position (row_id, number_in_row) values (3, 1);
INSERT INTO position (row_id, number_in_row) values (3, 2);
INSERT INTO position (row_id, number_in_row) values (3, 3);
INSERT INTO position (row_id, number_in_row) values (3, 4);

INSERT INTO position (row_id, number_in_row) values (4, 1);
INSERT INTO position (row_id, number_in_row) values (4, 2);
INSERT INTO position (row_id, number_in_row) values (5, 1);
INSERT INTO position (row_id, number_in_row) values (5, 2);

--- insert separately because of unique constraint
INSERT INTO session (movie_id, hall_id, start_at) values (1, 1, (now() AT TIME ZONE 'utc-3')); 
INSERT INTO session (movie_id, hall_id, start_at) values (2, 1, (now() AT TIME ZONE 'utc-3'));
INSERT INTO session (movie_id, hall_id, start_at) values (3, 1, (now() AT TIME ZONE 'utc-3'));

INSERT INTO session (movie_id, hall_id, start_at) values (1, 2, (now() AT TIME ZONE 'utc-3')); 
INSERT INTO session (movie_id, hall_id, start_at) values (3, 2, (now() AT TIME ZONE 'utc-3'));

INSERT INTO ticket (customer_id, session_id, price, row_id, position_id) values (1, 1, 120, 1, 1);
INSERT INTO ticket (customer_id, session_id, price, row_id, position_id) values (1, 1, 120, 1, 2);
INSERT INTO ticket (customer_id, session_id, price, row_id, position_id) values (2, 2, 120, 2, 2);

INSERT INTO ticket (customer_id, session_id, price, row_id, position_id) values (2, 1, 120, 1, 1);
INSERT INTO ticket (customer_id, session_id, price, row_id, position_id) values (3, 1, 140, 1, 2);
INSERT INTO ticket (customer_id, session_id, price, row_id, position_id) values (3, 2, 140, 2, 2);