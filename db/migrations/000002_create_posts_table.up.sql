CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    newsletter_id INT NOT NULL,
    CONSTRAINT fk_newsletter FOREIGN KEY newsletter_id REFERENCES newsletters(id),
);