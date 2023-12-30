CREATE TABLE users(
    ID SERIAL PRIMARY KEY,
    full_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    picture VARCHAR(255),
    username VARCHAR(255) UNIQUE,
    gender VARCHAR(20)
);