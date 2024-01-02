CREATE TABLE participant_rankings(
    ID VARCHAR(100) PRIMARY KEY,
    participation_id VARCHAR(100),
    ranking_id INT,
    grade INT,
    note TEXT,
    position INT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_participation_id
    FOREIGN KEY(participation_id)
    REFERENCES participations(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_ranking_id
    FOREIGN KEY(ranking_id)
    REFERENCES rankings(id)
    ON DELETE CASCADE
);