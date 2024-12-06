package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Evilcmd/data-metric-framework/internal/model"
	"github.com/Evilcmd/data-metric-framework/internal/storage/postgres/database"
	"github.com/Evilcmd/data-metric-framework/internal/utilities"
)

type DB struct {
	MetricDb       *database.MetricDBModel
	FoodDeliveryDb *database.FoodDeliveryDBModel
}

func (db *DB) AddMetric(res http.ResponseWriter, req *http.Request) {
	// utilities.RespondWithJson(res, 200, "Add metrics")
	query := model.Queries{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&query)
	if err != nil {
		utilities.RespondWithError(res, http.StatusBadRequest, "error reading request body")
		return
	}

	tx, err := db.MetricDb.DB.Begin()
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, "error starting transaction")
	}
	qtx := db.MetricDb.Queriess.WithTx(tx)

	addmetricparams := database.AddMetricParams{
		Querydescription: query.QueryDescription,
		Query:            query.Query,
	}
	addmetrics, err := qtx.AddMetric(context.Background(), addmetricparams)
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, "error inserting into metrics")
		return
	}

	newQuery := model.Queries{
		QueryId:          int(addmetrics.Queryid),
		QueryDescription: addmetrics.Querydescription,
		Query:            addmetrics.Query,
		Params:           make([]model.QueryParameter, len(query.Params)),
	}

	for i := 0; i < len(query.Params); i++ {
		addmetricparamsparams := database.AddMetricParamsParams{
			Queryid:       addmetrics.Queryid,
			Parametername: query.Params[i].ParameterName,
			Datatype:      query.Params[i].DataType,
			Ordered:       int32(query.Params[i].Ordered),
		}
		addmetricparam, err := qtx.AddMetricParams(context.Background(), addmetricparamsparams)
		if err != nil {
			utilities.RespondWithError(res, http.StatusInternalServerError, "error inserting into metrics parameters")
			return
		}
		newQuery.Params[i].ParameterId = int(addmetricparam.Parameterid)
		newQuery.Params[i].QueryId = int(addmetricparam.Queryid)
		newQuery.Params[i].ParameterName = addmetricparam.Parametername
		newQuery.Params[i].DataType = addmetricparam.Datatype
		newQuery.Params[i].Ordered = int(addmetricparam.Ordered)
	}
	tx.Commit()
	utilities.RespondWithJson(res, http.StatusOK, newQuery)

}

func (db *DB) GetAllMetrics(res http.ResponseWriter, req *http.Request) {
	// utilities.RespondWithJson(res, 200, "Get all metrics")
	queries, err := db.MetricDb.Queriess.GetAllMetrics(context.Background())
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, "error fetching queries")
		return
	}

	utilities.RespondWithJson(res, http.StatusOK, queries)
}

func (db *DB) GetMetric(res http.ResponseWriter, req *http.Request) {
	// utilities.RespondWithJson(res, 200, "Get metric")
	queryId := req.PathValue("metricId")
	if queryId == "" {
		utilities.RespondWithError(res, http.StatusBadRequest, "error reading metricId")
		return
	}
	qId, err := strconv.Atoi(queryId)
	if err != nil {
		utilities.RespondWithError(res, http.StatusBadRequest, "error converting metricId to int")
		return
	}

	metric, err := db.MetricDb.Queriess.GetMetric(context.Background(), int32(qId))
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, fmt.Sprintf("error while fetching metric: %v", err.Error()))
		return
	}

	metricparams, err := db.MetricDb.Queriess.GetMetricParameters(context.Background(), int32(qId))
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, fmt.Sprintf("error while fetching metric params: %v", err.Error()))
		return
	}

	query := model.Queries{
		QueryId:          int(metric.Queryid),
		QueryDescription: metric.Querydescription,
		Query:            metric.Query,
		Params:           make([]model.QueryParameter, len(metricparams)),
	}
	for i := 0; i < len(metricparams); i++ {
		query.Params[i].ParameterId = int(metricparams[i].Parameterid)
		query.Params[i].QueryId = int(metricparams[i].Queryid)
		query.Params[i].ParameterName = metricparams[i].Parametername
		query.Params[i].DataType = metricparams[i].Datatype
		query.Params[i].Ordered = int(metricparams[i].Ordered)
	}

	utilities.RespondWithJson(res, http.StatusOK, query)

}

func (db *DB) ExecuteQuery(res http.ResponseWriter, req *http.Request) {
	// utilities.RespondWithJson(res, 200, "Execute Query")
	queryId := req.PathValue("metricId")
	if queryId == "" {
		utilities.RespondWithError(res, http.StatusBadRequest, "error reading metricId")
		return
	}
	qId, err := strconv.Atoi(queryId)
	if err != nil {
		utilities.RespondWithError(res, http.StatusBadRequest, "error converting metricId to int")
		return
	}

	var data map[string]string
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&data)
	if err != nil {
		utilities.RespondWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}

	metric, err := db.MetricDb.Queriess.GetMetric(context.Background(), int32(qId))
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, fmt.Sprintf("error while fetching metrics: %v", err.Error()))
		return
	}
	queryparams, err := db.MetricDb.Queriess.GetMetricParameters(context.Background(), int32(qId))
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, fmt.Sprintf("error while fetching metric params: %v", err.Error()))
		return
	}

	query := metric.Query
	fmt.Println(query)
	for i := 0; i < len(queryparams); i++ {
		// fmt.Print(i)
		if queryparams[i].Datatype == "TEXT" {
			query = strings.Replace(query, "$"+strconv.Itoa(i+1), "'"+data[strconv.Itoa(i+1)]+"'", 1)
			// fmt.Println(" " + "$" + strconv.Itoa(i+1) + " " + "'" + data[strconv.Itoa(i+1)] + "'")
		} else {
			query = strings.Replace(query, "$"+strconv.Itoa(i+1), data[strconv.Itoa(i+1)], 1)
			// fmt.Println(" " + "$" + strconv.Itoa(i+1))
		}
	}
	fmt.Println(query)
	rows, err := db.FoodDeliveryDb.DB.Query(query)
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, fmt.Sprintf("error while executing query: %v", err.Error()))
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, fmt.Sprintf("error fetching columns: %v", err.Error()))
		return
	}

	results := []map[string]interface{}{}

	for rows.Next() {
		// Create a slice of `interface{}` to hold each column value
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the row into the value pointers
		if err := rows.Scan(valuePtrs...); err != nil {
			utilities.RespondWithError(res, http.StatusInternalServerError, fmt.Sprintf("error scanning row: %v", err.Error()))
			return
		}

		// Create a map for the row
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]

			// Convert `[]byte` to `string` if necessary
			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}

		results = append(results, rowMap)
	}
	utilities.RespondWithJson(res, http.StatusOK, results)
}
