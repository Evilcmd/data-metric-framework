// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query_parameters.sql

package database

import (
	"context"
)

const addMetricParams = `-- name: AddMetricParams :one
INSERT INTO query_parameters(QueryId, ParameterName, DataType, Ordered) VALUES($1, $2, $3, $4) RETURNING parameterid, queryid, parametername, datatype, ordered
`

type AddMetricParamsParams struct {
	Queryid       int32
	Parametername string
	Datatype      string
	Ordered       int32
}

func (q *Queries) AddMetricParams(ctx context.Context, arg AddMetricParamsParams) (QueryParameter, error) {
	row := q.db.QueryRowContext(ctx, addMetricParams,
		arg.Queryid,
		arg.Parametername,
		arg.Datatype,
		arg.Ordered,
	)
	var i QueryParameter
	err := row.Scan(
		&i.Parameterid,
		&i.Queryid,
		&i.Parametername,
		&i.Datatype,
		&i.Ordered,
	)
	return i, err
}

const getMetricParameters = `-- name: GetMetricParameters :many
SELECT parameterid, queryid, parametername, datatype, ordered FROM query_parameters WHERE QueryId = $1
`

func (q *Queries) GetMetricParameters(ctx context.Context, queryid int32) ([]QueryParameter, error) {
	rows, err := q.db.QueryContext(ctx, getMetricParameters, queryid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QueryParameter
	for rows.Next() {
		var i QueryParameter
		if err := rows.Scan(
			&i.Parameterid,
			&i.Queryid,
			&i.Parametername,
			&i.Datatype,
			&i.Ordered,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}