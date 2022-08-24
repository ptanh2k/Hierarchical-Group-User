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
('Tier 2', 'Tier_2', 'Viettel_Ticket.Tier_2')