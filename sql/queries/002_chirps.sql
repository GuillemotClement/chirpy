-- name: CreateChirps :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUEs (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1, 
  $2
)
RETURNING *;

-- name: ListChirps :many
SELECT * FROM chirps ORDER BY created_at;

-- name: DetailChirp :one
SELECT * FROM chirps WHERE id = $1;