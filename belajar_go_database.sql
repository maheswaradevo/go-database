-- auto-generated definition
create schema belajar_go_database collate utf8mb4_general_ci;

use belajar_go_database;

SELECT *
FROM customer;

DROP TABLE customer;

DELETE
FROM customer;

CREATE TABLE customer
(
    id varchar(100),
    name varchar(100)
) ENGINE = InnoDB;

ALTER TABLE customer
    ADD COLUMN email varchar(100),
    add column balance int default 0,
    add column rating double default 0.0,
    add column created_at timestamp default current_timestamp,
    add column birth_date date,
    add column married boolean default false;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES('devo', 'Devo', 'pundadevo21@gmail.com', 2500000, 90.0, '2002-05-25', false),
       ('reza', 'Reza', 'rezap@gmail.com', 5000000, 86.0, '1998-06-20', false),
       ('vio', 'vio', 'vio22@gmail.com', 100000, 82.6, '2005-07-12', false);

UPDATE customer
SET email = 'devo2@gmail.com'
WHERE id = 'devo';

CREATE TABLE user
(
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY(username)
)ENGINE = InnoDB;

SELECT * FROM user;

INSERT INTO user(username, password) VALUES('admin', 'salah');

CREATE TABLE comments
(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comments TEXT,
    PRIMARY KEY (id)
) ENGINE InnoDB;

DESC comments;

SELECT * FROM comments;

SELECT COUNT(*) FROM comments;