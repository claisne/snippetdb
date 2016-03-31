
CREATE TABLE snippets (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGSERIAL NOT NULL REFERENCES users(id),
    language_id BIGSERIAL NOT NULL REFERENCES languages(id),
    title TEXT NOT NULL,
    code TEXT NOT NULL,
    views_count INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL
);

