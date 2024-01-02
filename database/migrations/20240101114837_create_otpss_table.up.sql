CREATE TABLE otps(
    ID SERIAL PRIMARY KEY,
    user_id INT,
    otpCode VARCHAR(8),
    status VARCHAR(10),
    valid_until TIMESTAMP,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);