-- name: ListCategories :many
SELECT * FROM Categories;

-- name: GetCategory :one
SELECT * FROM Categories 
WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO Categories (id, name,description)
VALUES(?,?,?);

-- name: UpdateCategory :exec
UPDATE Categories SET name = ?, description = ?
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM Categories WHERE id = ?;