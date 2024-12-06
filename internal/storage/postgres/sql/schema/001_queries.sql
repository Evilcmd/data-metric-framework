-- +goose Up
CREATE TABLE queries (
    QueryId SERIAL PRIMARY KEY, 
    QueryDescription TEXT NOT NULL, 
    Query TEXT NOT NULL
);

-- +goose Down
DROP TABLE queries CASCADE;