CREATE TABLE rankings(
    ID SERIAL PRIMARY KEY,
    name VARCHAR(100),
    icon VARCHAR(255),
    note TEXT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);