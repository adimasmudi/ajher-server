CREATE TABLE participations(
    ID VARCHAR(100) PRIMARY KEY,
    user_id INT,
    quiz_id VARCHAR(100),
    status VARCHAR(10),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_quiz_id
    FOREIGN KEY(quiz_id)
    REFERENCES quizes(id)
    ON DELETE CASCADE
);