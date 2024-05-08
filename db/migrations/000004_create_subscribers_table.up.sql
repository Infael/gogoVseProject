CREATE TABLE IF NOT EXISTS subscribers (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    CONSTRAINT fk_newsletter FOREIGN KEY newsletter_id REFERENCES newsletters(id) NOT NULL,
    UNIQUE (email)
);