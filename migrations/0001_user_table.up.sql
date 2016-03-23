
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_activity_at TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX idx_users_username on users (username);

