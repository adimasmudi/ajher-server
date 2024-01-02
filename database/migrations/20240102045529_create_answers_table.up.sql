CREATE TABLE answers(
    ID VARCHAR(100) PRIMARY KEY,
    user_id INT,
    question_id VARCHAR(100),
    grade FLOAT,
    label VARCHAR(10),
    answer TEXT,
    answer_duration INT,
    status VARCHAR(10),
    generated_suggestion TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_question_id
    FOREIGN KEY(question_id)
    REFERENCES questions(id)
    ON DELETE CASCADE
);