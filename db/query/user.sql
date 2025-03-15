-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM public.users
ORDER BY id ASC;

-- name: CreateUser :exec
INSERT INTO public.users (
    email, password, type
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE public.users
SET email = $2 , password = $3, type = $4
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, password, type, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1;

