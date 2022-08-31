CREATE DATABASE cycir WITH ENCODING 'UTF8' TEMPLATE template0;

--- TABLE for Group
CREATE TABLE group_ (
    gid INT PRIMARY KEY NOT NULL,
	name VARCHAR(50) NOT NULL,
	parent_id INT
)

INSERT INTO group_ (gid, name, parent_id) VALUES
(1, 'Viettel Ticket', NULL),
(2, 'Tier 1', 1),
(3, 'Tier 2', 1),
(4, 'Group Viettel 1', 2),
(5, 'Group Viettel 2', 2),
(6, 'Group Viettel 3', 2),
(7, 'Group Viettel 1', 3),
(8, 'Group Viettel 2', 3),
(9, 'Group Viettel 3', 3),
(10, 'Viettel Cycir', NULL)

INSERT INTO group_ (gid, name, parent_id) VALUES
(11, 'Viettel Ticket', 2),
(20, 'Tier 1', 2)

INSERT INTO group_ (gid, name, parent_id) VALUES
(12, 'Viettel Ticket', 11),
(21, 'Tier 1', 12)

SELECT * FROM group_

--- TABLE for User
CREATE TABLE user_ (
	uid SERIAL PRIMARY KEY,
	username VARCHAR(30) UNIQUE NOT NULL,
	firstname VARCHAR(50) NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	email TEXT UNIQUE NOT NULL,
	gid INT REFERENCES group_(gid)
)

INSERT INTO user_ (username, firstname, lastname, email, gid) VALUES
('feelingBlue00', 'Tom', 'Hansen', 'feelingBlue@gmail.com', 4),
('Chris B Bacon', 'Chris', 'Dao', 'Ilovebacon@gmail.com', 5),
('Harvey Nguyen', 'Harvey', 'Nguyen', 'tibi2k@gmail.com', 5),
('Viettel', 'Viet', 'tel', 'Viettel@gmail.com', 6),
('Henry', 'Henry', 'Nguyen', 'lmh@gmail.com', 7),
('NB_Prince', 'Hoang Anh', 'Duong Minh', 'dmha@gmail.com', 9)

-- Recursive
WITH RECURSIVE group_tree(gid, name, parent_id, lvl) AS (
    SELECT gid, name, parent_id, 1 AS level
    FROM group_ 
    WHERE parent_id = 1
  UNION ALL
    SELECT bg.gid, bg.name, bg.parent_id, gt.lvl + 1
    FROM group_ bg
    JOIN group_tree gt ON bg.parent_id = gt.gid
)
SELECT * FROM group_tree;