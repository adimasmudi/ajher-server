CREATE TABLE quiz_categories(
    ID SERIAL PRIMARY KEY,
    category_name VARCHAR(255),
    icon VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);