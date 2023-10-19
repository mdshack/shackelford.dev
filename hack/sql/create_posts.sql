CREATE TABLE posts (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title varchar(255) NOT NULL,
    slug varchar(255) UNIQUE NOT NULL,
    contents TEXT NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp
);