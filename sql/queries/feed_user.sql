-- name: CreateFeedFollow :one
INSERT INTO users_feeds(id, created_at, updated_at, user_id, feed_id)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM users_feeds 
WHERE id=$1;

-- name: GetAllFeedFollowsForUser :many
SELECT * FROM users_feeds
WHERE user_id=$1;