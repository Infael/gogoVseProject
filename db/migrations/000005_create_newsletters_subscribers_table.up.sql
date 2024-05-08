CREATE TABLE newsletters_subscribers (
  newsletter_id INT
  subscriber_id INT
  PRIMARY KEY (newsletter_id, subscriber_id)
  CONSTRAINT fk_newsletter FOREIGN KEY(newsletter_id) REFERENCES newsletters(id)
  CONSTRAINT fk_subscriber FOREIGN KEY(subscriber_id) REFERENCES subscribers(id)
)