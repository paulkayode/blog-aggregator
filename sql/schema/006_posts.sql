-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    published_at TEXT NOT NULL,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feed
    FOREIGN KEY(feed_id)
    REFERENCES feeds(id)
    ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE posts;