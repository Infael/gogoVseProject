ALTER TABLE newsletters_subscribers
ADD COLUMN verified BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN verification_token VARCHAR(255) DEFAULT (substr(md5(random()::text), 1, 20)) NOT NULL,
ADD COLUMN subscribed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
ADD COLUMN verified_at TIMESTAMP;