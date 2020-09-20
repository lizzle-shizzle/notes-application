CREATE TABLE IF NOT EXISTS note (
    id serial PRIMARY KEY,
    text text NOT NULL,
    created_timestamp timestamp with time zone NOT NULL default now()
);