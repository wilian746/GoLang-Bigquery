package main

import (
	"github.com/GoLang-Bigquery/api"
	bigqueryservice "github.com/GoLang-Bigquery/services/bigquery"
	"github.com/GoLang-Bigquery/utils"

	"log"
	"net/http"
)

func main() {
	bigqueryservice.GetDataSet()

	log.Println("Service running on port 8071")
	log.Fatal(http.ListenAndServe(":"+utils.GetEnvOrDefault("PORT", "8071"), api.Server()))
}
