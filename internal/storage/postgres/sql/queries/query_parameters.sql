-- name: AddMetricParams :one
INSERT INTO query_parameters(QueryId, ParameterName, DataType, Ordered) VALUES($1, $2, $3, $4) RETURNING *;

-- name: GetMetricParameters :many
SELECT * FROM query_parameters WHERE QueryId = $1;