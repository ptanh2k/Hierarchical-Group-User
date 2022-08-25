CREATE DATABASE cycir WITH ENCODING 'UTF8' TEMPLATE template0;

CREATE EXTENSION IF NOT EXISTS ltree;

CREATE TABLE group_ (
	name VARCHAR(50) NOT NULL PRIMARY KEY,
	name_in_path VARCHAR(60) NOT NULL,
	path ltree UNIQUE NOT NULL
)

INSERT INTO group_ (name, name_in_path, path) VALUES
('Viettel Ticket', 'Viettel_Ticket', 'Viettel_Ticket'),
('Tier 1', 'Tier_1', 'Viettel_Ticket.Tier_1'),
('Tier 2', 'Tier_2', 'Viettel_Ticket.Tier_2'),

INSERT INTO group_ (name, name_in_path, path) VALUES
('Group Viettel 1', 'Group_Viettel_1', 'Viettel_Ticket.Tier_1.Group_Viettel_1'),
('Group Viettel 2', 'Group_Viettel_2', 'Viettel_Ticket.Tier_1.Group_Viettel_2'),
('Group Viettel 3', 'Group_Viettel_3', 'Viettel_Ticket.Tier_1.Group_Viettel_3')

INSERT INTO group_ (name, name_in_path, path) VALUES
('Group Viettel 1', 'Group_Viettel_1', 'Viettel_Ticket.Tier_2.Group_Viettel_1'),
('Group Viettel 2', 'Group_Viettel_2', 'Viettel_Ticket.Tier_2.Group_Viettel_2'),
('Group Viettel 3', 'Group_Viettel_3', 'Viettel_Ticket.Tier_2.Group_Viettel_3')

SELECT * FROM group_

SELECT * FROM group_ WHERE path @ 'Tier_1'

-- Change PRIMARY KEY for group_ table --
ALTER TABLE group_ DROP CONSTRAINT group__pkey;

ALTER TABLE group_ ADD PRIMARY KEY (path);