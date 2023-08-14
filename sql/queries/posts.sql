-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, description, published_at, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT p.title, p.description, p.published_at, p.url, f.name as feed_name FROM posts p
JOIN feed_follows ff ON p.feed_id = ff.feed_id
JOIN feeds f ON f.id = ff.feed_id
where ff.user_id = $1;