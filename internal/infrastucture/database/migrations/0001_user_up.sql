create table users(
	id uuid default gen_random_uuid(),
	name text,
	email text,
	password text,
);
