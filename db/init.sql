CREATE TABLE IF NOT EXISTS segments
(
    slug TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users_segments
(
    user_id      INT NOT NULL,
    segment_slug TEXT
);

CREATE TYPE OPERATION AS ENUM ('add', 'delete');

CREATE TABLE IF NOT EXISTS users_segments_history
(
    user_id      INT       NOT NULL,
    segment_slug TEXT      NOT NULL,
    operation    OPERATION NOT NULL,
    datetime     TIMESTAMP NOT NULL
)