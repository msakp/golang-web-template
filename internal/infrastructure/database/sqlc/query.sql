-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: CreateUser :exec
insert into users(name, email, password) values ($1, $2, $3);

-- name: GetUserById :one
select * from users where id = $1;
