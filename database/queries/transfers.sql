-- name: CreateTransfers :one
INSERT INTO transfers (
  owner, balance, currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id;

-- name: UpdateTransfers :one
UPDATE transfers
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1;