package server

import (
	"net/http"

	"github.com/Evilcmd/data-metric-framework/internal/apis"
	"github.com/Evilcmd/data-metric-framework/internal/middleware"
	"github.com/Evilcmd/data-metric-framework/internal/storage/postgres/database"
)

func GetRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", apis.Root)
	router.HandleFunc("GET /health", apis.CheckHealth)
	router.HandleFunc("GET /err", apis.ErrCheck)

	db := apis.DB{
		MetricDb:       database.GetMetricDBModel(),
		FoodDeliveryDb: database.GetFoodDeliveryDBModel(),
	}
	router.HandleFunc("POST /api/v1/metrics", db.AddMetric)
	router.HandleFunc("GET /api/v1/metrics", db.GetAllMetrics)
	router.HandleFunc("GET /api/v1/metrics/{metricId}", db.GetMetric)
	router.HandleFunc("POST /api/v1/metrics/{metricId}", db.ExecuteQuery)

	return middleware.Cors(router)
}
