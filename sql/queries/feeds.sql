-- name: CreateFeed :one
INSERT INTO feeds(
    id, created_at, updated_at,
    name, url, user_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT feeds.id, feeds.name, feeds.url,
users.name AS creator
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;