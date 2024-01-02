CREATE TABLE questions(
    ID VARCHAR(100) PRIMARY KEY,
    quiz_id VARCHAR(100),
    question TEXT,
    reference_answer TEXT,
    grade_percentage FLOAT,
    status VARCHAR(10),
    duration INT,
    point FLOAT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_quiz_id
    FOREIGN KEY(quiz_id)
    REFERENCES quizes(id)
    ON DELETE CASCADE
);