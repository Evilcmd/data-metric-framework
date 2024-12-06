-- +goose Up
CREATE TABLE query_parameters (
    ParameterId SERIAL PRIMARY KEY, 
    QueryId INT NOT NULL, 
    ParameterName TEXT NOT NULL, 
    DataType TEXT NOT NULL, 
    Ordered INT NOT NULL, 
    FOREIGN KEY (QueryId) REFERENCES queries (QueryId)
);

-- +goose Down
DROP TABLE query_parameters CASCADE;