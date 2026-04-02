
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

-- name: UserLogin :one
SELECT id, name, email, password, restaurant_id FROM users WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET name = coalesce(sqlc.narg(name), name),
    email = coalesce(sqlc.narg(email), email),
    password = coalesce(sqlc.narg(password), password),
    phone_number = coalesce(sqlc.narg(phone_number), phone_number),
    image_url = coalesce(sqlc.narg(image_url), image_url),
    restaurant_id = coalesce(sqlc.narg(restaurant_id), restaurant_id)
WHERE id = sqlc.arg(id) RETURNING name, email, phone_number, image_url, restaurant_id;

-- name: CheckEmail :one
SELECT email FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: NewResTaurant :one
INSERT INTO restaurants ( name ) VALUES ( sqlc.arg('name')) RETURNING *;

-- name: DeleteRestaurant :exec
DELETE FROM restaurants WHERE id = $1;

-- name: CheckRestaurantID :one
SELECT id FROM restaurants WHERE id = $1;
