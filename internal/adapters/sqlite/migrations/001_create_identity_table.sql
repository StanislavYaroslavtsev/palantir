-- +goose Up
CREATE TABLE IF NOT EXISTS identity (
    id          INTEGER     PRIMARY KEY CHECK (id = 1),
    peer_id     TEXT        NOT NULL,
    priv_key    BLOB        NOT NULL,
    created_at  INTEGER     NOT NULL DEFAULT (unixepoch())
);

-- +goose Down
DROP TABLE IF EXISTS identity;