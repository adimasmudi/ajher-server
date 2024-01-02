CREATE TABLE participant_questions(
    ID VARCHAR(100) PRIMARY KEY,
    participation_id VARCHAR(100),
    question_id VARCHAR(100),
    number INT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_participation_id
    FOREIGN KEY(participation_id)
    REFERENCES participations(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_question_id
    FOREIGN KEY(question_id)
    REFERENCES questions(id)
    ON DELETE CASCADE
);