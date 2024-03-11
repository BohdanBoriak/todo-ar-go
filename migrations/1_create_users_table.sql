CREATE TABLE IF NOT EXISTS public.users (
    id serial PRIMARY KEY,
    name text NOT NULL,
    password text NOT NULL
);