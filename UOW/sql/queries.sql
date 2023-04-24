-- name: CreateCategory :exec
INSERT INTO categories(id, name) VALUES (?,?);

-- name: CreateCourses :exec
INSERT INTO courses (id, name, category_id) VALUES (?, ?, ?);
