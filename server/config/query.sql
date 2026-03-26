
-- name: GetUserList :many
SELECT * FROM users ORDER BY name;

-- name: NewUser :one
INSERT INTO users (
 id, name, email, password, phone_number, image_url, restaurant_id ,created_at, updated_at
) VALUES ( $1, $2, $3, $4, $5, $6,null, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP ) RETURNING *;
