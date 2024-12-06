-- name: AddMetric :one
INSERT INTO queries(QueryDescription, Query) VALUES($1, $2) RETURNING *;

-- name: GetAllMetrics :many
SELECT * FROM queries;

-- name: GetMetric :one
SELECT * FROM queries WHERE QueryId = $1;
