
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
SELECT id, name, email, password, restaurant_id FROM users WHERE email = $1 LIMIT 1;

-- -- name: CheckUserExist :one
-- SELECT * FROM

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

-- name: UpdateRestaurant :one
UPDATE restaurants  SET name = coalesce(sqlc.narg(name), name)
 WHERE id = sqlc.arg(id) RETURNING name;

-- name: GetUserResById :one
SELECT restaurant_id FROM users WHERE id = $1;

-- name: CreateResAddress :one
INSERT INTO res_addresses (
  restaurant_id,
  street_name,
  city,
  state,
  phone,
  email,
  is_default
) VALUES (
  sqlc.arg('restaurant_id'),
  sqlc.narg('street_name'),
  sqlc.narg('city'),
  sqlc.narg('state'),
  sqlc.narg('phone'),
  sqlc.narg('email'),
  sqlc.arg('is_default')
) RETURNING *;

-- name: UpdateDefaultResBranch :exec
UPDATE res_addresses SET is_default = false
     WHERE restaurant_id = $1 AND is_default = true;

-- name: GetUserResAddressesById :many
SELECT * FROM res_addresses WHERE restaurant_id = $1 ORDER BY created_at;

-- name: UpdateResAddress :one
UPDATE res_addresses SET
   street_name = coalesce(sqlc.narg(street_name), street_name),
   city = COALESCE(sqlc.narg(city), city),
   state = COALESCE(sqlc.narg(state), state),
   phone = COALESCE(sqlc.narg(phone), phone),
   email = COALESCE(sqlc.narg(email), email),
   is_default = COALESCE(sqlc.narg('is_default'), is_default)
WHERE id = $1 RETURNING * ;

-- name: GetRestaurant :one
SELECT * FROM restaurants WHERE id = $1;

-- name: DeleteResAddress :exec
DELETE FROM res_addresses WHERE id = $1;

-- name: CreateUserAddress :one
INSERT INTO user_addresses (
user_id,
city,
state,
zip_code,
is_default
) VALUES(
sqlc.arg('user_id'),
sqlc.narg('city'),
sqlc.narg('state'),
sqlc.narg('zip_code'),
sqlc.arg('is_default')
) RETURNING *;

-- name: SetDefaultUserAddress :exec
UPDATE user_addresses SET is_default = false
     WHERE user_id = $1 AND is_default = true;
