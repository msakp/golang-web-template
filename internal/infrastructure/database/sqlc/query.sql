-- name: GetUser :one
select * from users
where email = $1;

-- name: CreateUser :exec
insert into users(name, email, password)
values ($1, $2, $3);
