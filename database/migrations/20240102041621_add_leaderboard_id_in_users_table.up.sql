ALTER TABLE users 
ADD COLUMN leaderboard_id INT,
ADD CONSTRAINT fk_leaderboard_id 
FOREIGN KEY (leaderboard_id) REFERENCES leaderboards (id)
ON DELETE CASCADE;