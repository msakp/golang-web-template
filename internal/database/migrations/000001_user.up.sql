CREATE TABLE users(
	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
	name text NOT NULL,
	email text NOT NULL,
	password text NOT NULL
);
