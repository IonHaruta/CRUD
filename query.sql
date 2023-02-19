-- name: CreatePost :one
INSERT INTO Produs (
  name,
  price,
  quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPostById :one
SELECT * FROM Produs
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM Produs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE Produs
set 
name = coalesce(sqlc.narg('title'), title), 
price = coalesce(sqlc.narg('category'), category), 
quantity = coalesce(sqlc.narg('content'), content)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;