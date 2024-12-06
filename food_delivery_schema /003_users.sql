-- +goose Up
create table users(UserId INT PRIMARY KEY, UserName TEXT);

-- +goose Down
DROP TABLE users CASCADE;