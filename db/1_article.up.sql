CREATE TABLE IF NOT EXISTS article (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    content VARCHAR(65536),
    creation_timestamp TIMESTAMP
);