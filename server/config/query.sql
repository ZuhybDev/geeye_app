
-- name: GetUserList :many
SELECT * FROM users ORDER BY name;

-- name: NewUser :one
INSERT INTO users (
  name,
  email,
  password,
  phone_number,
  image_url,
  restaurant_id
) VALUES (
  sqlc.arg('name'),
  sqlc.arg('email'),
  sqlc.arg('password'),
  sqlc.narg('phone_number'),
  sqlc.narg('image_url'),
  sqlc.narg('restaurant_id')
) RETURNING *;

-- name: CheckPassword :one
SELECT id, name, email, password FROM users WHERE email = $1;
