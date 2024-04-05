CREATE TABLE IF NOT EXISTS public.tasks (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id),
    title varchar(255) NOT NULL,
    description text,
    status varchar(100) NOT NULL,
    date timestamptz
);