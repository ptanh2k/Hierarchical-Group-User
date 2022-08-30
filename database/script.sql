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

SELECT path, nlevel(path) FROM group_ WHERE path @ 'Tier_1'

SELECT * FROM group_

ALTER TABLE group_ ADD gid SERIAL

-- Change PRIMARY KEY for group_ table --
-- ALTER TABLE group_ DROP CONSTRAINT group__pkey;

-- ALTER TABLE group_ ADD PRIMARY KEY (path);

SELECT name, path, nlevel(path), subpath(path, 0, -1) as parent from group_

SELECT path, nlevel(path) FROM group_ WHERE name='Tier 1'

--- TABLE for User
CREATE TABLE user_ (
	uid SERIAL PRIMARY KEY,
	username VARCHAR(30) UNIQUE NOT NULL,
	firstname VARCHAR(50) NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	email TEXT UNIQUE NOT NULL,
	gid INT REFERENCES group_(gid)
)

DROP TABLE user_

INSERT INTO user_ (username, firstname, lastname, email, gid) VALUES
('feelingBlue00', 'Tom', 'Hansen', 'feelingBlue@gmail.com', 4),
('Chris B Bacon', 'Chris', 'Dao', 'Ilovebacon@gmail.com', 5),
('Harvey Nguyen', 'Harvey', 'Nguyen', 'tibi2k@gmail.com', 5),
('Viettel', 'Viet', 'tel', 'Viettel@gmail.com', 6),
('Henry', 'Henry', 'Nguyen', 'lmh@gmail.com', 7),
('NB_Prince', 'Hoang Anh', 'Duong Minh', 'dmha@gmail.com', 9)

SELECT u.username, u.firstname, u.lastname, u.email, g.path
FROM user_ u
INNER JOIN group_ g ON u.gid = g.gid WHERE u.uid = 1
