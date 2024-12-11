CREATE TABLE IF NOT EXISTS comment (
    id SERIAL PRIMARY KEY,
    article_id INTEGER REFERENCES article (id),
    author VARCHAR(255),
    content VARCHAR(65536),
    creation_timestamp TIMESTAMP
);
