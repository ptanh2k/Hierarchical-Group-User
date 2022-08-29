CREATE DATABASE cycir WITH ENCODING 'UTF8' TEMPLATE template0;

CREATE EXTENSION IF NOT EXISTS ltree;

--- TABLE for Group
CREATE TABLE group_ (
	name VARCHAR(50) NOT NULL,
	name_in_path VARCHAR(60) NOT NULL,
	path ltree UNIQUE NOT NULL PRIMARY KEY
)

INSERT INTO group_ (name, name_in_path, path) VALUES
('Viettel Ticket', 'Viettel_Ticket', 'Viettel_Ticket'),
('Tier 1', 'Tier_1', 'Viettel_Ticket.Tier_1'),
('Tier 2', 'Tier_2', 'Viettel_Ticket.Tier_2'),
('Group Viettel 1', 'Group_Viettel_1', 'Viettel_Ticket.Tier_1.Group_Viettel_1'),
('Group Viettel 2', 'Group_Viettel_2', 'Viettel_Ticket.Tier_1.Group_Viettel_2'),
('Group Viettel 3', 'Group_Viettel_3', 'Viettel_Ticket.Tier_1.Group_Viettel_3'),
('Group Viettel 1', 'Group_Viettel_1', 'Viettel_Ticket.Tier_2.Group_Viettel_1'),
('Group Viettel 2', 'Group_Viettel_2', 'Viettel_Ticket.Tier_2.Group_Viettel_2'),
('Group Viettel 3', 'Group_Viettel_3', 'Viettel_Ticket.Tier_2.Group_Viettel_3')

SELECT * FROM group_ WHERE path @ 'Tier_1'

-- Change PRIMARY KEY for group_ table --
-- ALTER TABLE group_ DROP CONSTRAINT group__pkey;

-- ALTER TABLE group_ ADD PRIMARY KEY (path);

--- TABLE for User
CREATE TABLE user_ (
	uid SERIAL PRIMARY KEY,
	username VARCHAR(30) UNIQUE NOT NULL,
	firstname VARCHAR(50) NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	email TEXT UNIQUE NOT NULL,
	path ltree REFERENCES group_(path)
)

INSERT INTO user_ (username, firstname, lastname, email, path) VALUES
('feelingBlue00', 'Tom', 'Hansen', 'feelingBlue@gmail.com', 'Viettel_Ticket.Tier_1.Group_Viettel_1'),
('Chris B Bacon', 'Chris', 'Dao', 'Ilovebacon@gmail.com', 'Viettel_Ticket.Tier_1.Group_Viettel_2'),
('Harvey Nguyen', 'Harvey', 'Nguyen', 'tibi2k@gmail.com', 'Viettel_Ticket.Tier_1.Group_Viettel_1'),
('Viettel', 'Viet', 'tel', 'Viettel@gmail.com', 'Viettel_Ticket'),
('Henry', 'Henry', 'Nguyen', 'lmh@gmail.com', 'Viettel_Ticket.Tier_2.Group_Viettel_1'),
('NB_Prince', 'Hoang Anh', 'Duong Minh', 'dmha@gmail.com', 'Viettel_Ticket.Tier_2.Group_Viettel_3')

SELECT * FROM user_ WHERE path @ 'Tier_1'
