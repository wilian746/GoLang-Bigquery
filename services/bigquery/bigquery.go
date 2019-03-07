package bigqueryservice

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/GoLang-Bigquery/utils"
)

var client *bigquery.Client
var ctx = context.Background()

func GetClient(reconnect bool) *bigquery.Client {
	utils.CheckEnvExist("GOOGLE_APPLICATION_CREDENTIALS")

	if client == nil || reconnect {
		connection, err := bigquery.NewClient(ctx, utils.GetEnvOrDefault("PROJECTID", "bigquery"))

		client = connection

		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
	}

	return client
}

func SetClient(client *bigquery.Client) {
	client = client
}

func GetDataSet() *bigquery.Dataset {
	GetClient(false).Dataset("productions").Create(ctx, &bigquery.DatasetMetadata{})

	dataset := GetClient(false).Dataset("productions")

	return dataset
}

func GetTable(tableName string, SchemaProduction bigquery.Schema) (*bigquery.Table, error) {
	table := GetDataSet().Table(tableName)
	_, err := table.Metadata(ctx)

	if err != nil {
		log.Println("Error in get Table:", err)
		err := table.Create(ctx, &bigquery.TableMetadata{Schema: SchemaProduction})

		if err != nil {
			log.Println(err)
		}

		return table, err
	} else {
		return table, nil
	}
}
