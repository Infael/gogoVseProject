CREATE TABLE IF NOT EXISTS newsletters (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    creator_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY creator_id REFERENCES users(id),
);