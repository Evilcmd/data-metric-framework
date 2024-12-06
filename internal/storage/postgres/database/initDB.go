package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type MetricDBModel struct {
	Queriess *Queries
	DB       *sql.DB
}

func GetMetricDBModel() *MetricDBModel {
	metricConString := os.Getenv("METRICSCONSTRING")
	db, err := sql.Open("postgres", metricConString)
	if err != nil {
		log.Fatalf("Failed to open a DB connection to metrics: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	q := New(db)
	metricDBmodel := MetricDBModel{
		Queriess: q,
		DB:       db,
	}
	return &metricDBmodel
}

type FoodDeliveryDBModel struct {
	DB *sql.DB
}

func GetFoodDeliveryDBModel() *FoodDeliveryDBModel {
	metricConString := os.Getenv("FOODDELIVERYCONSTRING")
	db, err := sql.Open("postgres", metricConString)
	if err != nil {
		log.Fatalf("Failed to open a DB connection to food delivery: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	foodDeliveryDBModel := FoodDeliveryDBModel{db}
	return &foodDeliveryDBModel
}
