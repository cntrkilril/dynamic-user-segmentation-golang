CREATE TABLE IF NOT EXISTS segments
(
    slug TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users_segments
(
    user_id      INT NOT NULL,
    segment_slug TEXT
);