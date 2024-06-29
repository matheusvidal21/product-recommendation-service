-- User queries

-- name: CreateUser :one
INSERT INTO users (id, name, email, password)
VALUES ($1, $2, $3, $4)
    RETURNING id, name, email;

-- name: GetUserByID :one
SELECT id, name, email, password
FROM users
WHERE id = $1;

-- name: GetAllUsers :many
SELECT id, name, email
FROM users;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, password = $4
WHERE id = $1
RETURNING id, name, email;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- Product queries

-- name: CreateProduct :one
INSERT INTO products (id, name, price, category_id)
VALUES ($1, $2, $3, $4)
RETURNING id, name, price, category_id;

-- name: GetProductByID :one
SELECT id, name, price, category_id
FROM products
WHERE id = $1;

-- name: GetAllProducts :many
SELECT id, name, price, category_id
FROM products;

-- name: UpdateProduct :one
UPDATE products
SET name = $2, price = $3, category_id = $4
WHERE id = $1
RETURNING id, name, price, category_id;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- Category queries

-- name: CreateCategory :one
INSERT INTO categories (id, name, description)
VALUES ($1, $2, $3)
    RETURNING id, name, description;

-- name: GetCategoryByID :one
SELECT id, name, description
FROM categories
WHERE id = $1;

-- name: GetAllCategories :many
SELECT id, name, description
FROM categories;

-- name: UpdateCategory :one
UPDATE categories
SET name = $2, description = $3
WHERE id = $1
RETURNING id, name, description;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- UserActivity queries

-- name: SaveActivity :exec
INSERT INTO user_activities (user_id, product_id, action)
VALUES ($1, $2, $3)
    ON CONFLICT (user_id, product_id) DO UPDATE
                                             SET action = $3;

-- name: GetActivityByUserId :many
SELECT user_id, product_id, action
FROM user_activities
WHERE user_id = $1;

-- name: GetAllActivities :many
SELECT user_id, product_id, action
FROM user_activities;
