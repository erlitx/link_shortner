BEGIN;

CREATE TABLE urls (
    id              SERIAL PRIMARY KEY,
    rawUrl          TEXT,
    shortUrl        VARCHAR(40)
);

COMMIT;