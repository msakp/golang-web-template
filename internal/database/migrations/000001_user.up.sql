BEGIN;
create table users(
	id uuid default gen_random_uuid() primary key,
	name text not null,
	email text not null,
	password text not null
);
COMMIT;
