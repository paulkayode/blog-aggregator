-- +goose Up
CREATE TABLE users_feeds(
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    user_id UUID NOT NULL,
    feed_id UUID NOT NULL,
    UNIQUE(user_id, feed_id),
    CONSTRAINT fk_user
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_feed
    FOREIGN KEY(feed_id)
    REFERENCES feeds(id)
    ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE users_feeds;