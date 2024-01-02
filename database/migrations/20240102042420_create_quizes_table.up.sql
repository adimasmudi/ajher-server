CREATE TABLE quizes(
    ID VARCHAR(100) PRIMARY KEY,
    title VARCHAR(255),
    duration INT NULL,
    end_at TIMESTAMP,
    code VARCHAR(50),
    status VARCHAR(10),
    description TEXT,
    quiz_category_id INT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_quiz_category_id
    FOREIGN KEY(quiz_category_id)
    REFERENCES quiz_categories(id)
    ON DELETE CASCADE
);