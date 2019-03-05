CREATE TABLE employee (
id INT UNIQUE PRIMARY KEY NOT NULL,
name VARCHAR(255) NULL,
birthdate TIMESTAMP NULL,
gender gender_type NULL,
married boolean NULL
);

CREATE TYPE gender_type AS ENUM ('M', 'F');