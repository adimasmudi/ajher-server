CREATE TABLE question_images(
    ID VARCHAR(100) PRIMARY KEY,
    question_id VARCHAR(100),
    image_path VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_question_id
    FOREIGN KEY(question_id)
    REFERENCES questions(id)
    ON DELETE CASCADE
);