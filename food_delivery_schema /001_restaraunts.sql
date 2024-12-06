-- +goose Up
CREATE TABLE restaraunts(RestarauntId INT PRIMARY KEY, RestarauntName TEXT);

-- +goose Down
DROP TABLE restaraunts CASCADE;