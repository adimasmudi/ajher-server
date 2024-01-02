CREATE TABLE leaderboards(
    ID SERIAL PRIMARY KEY,
    user_id INT,
    ranking_id INT,
    total_point INT DEFAULT 0,
    note TEXT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_ranking_id
    FOREIGN KEY(ranking_id)
    REFERENCES rankings(id)
    ON DELETE CASCADE
);