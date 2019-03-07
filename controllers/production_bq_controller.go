package controllers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/bigquery"
	"github.com/GoLang-Bigquery/entities"
	bigqueryservice "github.com/GoLang-Bigquery/services/bigquery"
	"github.com/pkg/errors"
)

var ctx = context.Background()
var ErrorUnmarshalProductionBQ = errors.New("Error while trying to Unmarshal ProductionBQ Type")
var ErrorGenerateUUIDProductionBQ = errors.New("Error while trying to Generate UUID ProductionBQ Type")
var ErrorInferSchemaProductionBQ = errors.New("Error while trying to InferSchema ProductionBQ Type")
var ErrorInvalidProductionBQId = errors.New("{SaveProductionBQ} ProductionBQId is not a uuid")
var ErrorSaveProductionBQOnBigQuery = errors.New("{SaveProductionBQ} Error while trying to save ProductionBQ on BigQuery")
var ErrorGetLocationTimeZoneProductionBQ = errors.New("{SaveProductionBQ} Error while trying get location in time zone ProductionBQ")
var ErrorProductioBQDoesNotExist = errors.New("productionBQ does not exist")
var ErrorDeleteProductionBQOnBigQuery = errors.New("{DeleteProductionBQ} Error while trying to delete ProductionBQ on BigQuery")

type ProductionBQControllerI interface {
	SaveBQProduction(bytes []byte) (*entities.Production, error)
	GetAllBQProduction() ([]byte, error)
	GetOneBQProduction(ProductionId string) ([]byte, error)
	DeleteBQProduction(ProductionId string) error
}

func SaveBQProduction(bytes []byte) (*entities.Production, error) {
	Production := &entities.Production{}

	err := json.Unmarshal(bytes, Production)

	if err != nil {
		return Production, errors.Wrap(err, ErrorUnmarshalProductionBQ.Error())
	}

	err = json.Unmarshal(bytes, Production)

	if err != nil {
		return Production, errors.Wrap(err, ErrorUnmarshalProductionBQ.Error())
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")

	Production.ProductionId = uuid.New().String()
	Production.CreatedAt = time.Now().In(loc)
	Production.UpdatedAt = time.Now().In(loc)

	SchemaProduction, err := bigquery.InferSchema(Production)

	if err != nil {
		return Production, errors.Wrap(err, ErrorInferSchemaProductionBQ.Error())
	}

	tableInstance, err := bigqueryservice.GetTable("production", SchemaProduction)

	err = tableInstance.Uploader().Put(ctx, Production)

	if err != nil {
		return Production, errors.Wrap(err, ErrorSaveProductionBQOnBigQuery.Error())
	}

	log.Println("Production: ", Production)

	return Production, err
}

func GetOneBQProduction(ProductionId string) ([]byte, error) {
	SchemaProduction, err := bigquery.InferSchema(&entities.Production{})

	if err != nil {
		return nil, errors.Wrap(err, ErrorInferSchemaProductionBQ.Error())
	}

	tableInstance, err := bigqueryservice.GetTable("production", SchemaProduction)

	ProductionRead := tableInstance.Read(ctx)

	ProductionSelected := entities.Production{}
	for {
		var row entities.Production
		err := ProductionRead.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if row.ProductionId == ProductionId {
			ProductionSelected = row
		}
	}

	if ProductionSelected.ProductionId == "" {
		return nil, nil
	}

	bytes, _ := json.Marshal(ProductionSelected)

	return bytes, nil
}

func GetAllBQProduction() ([]byte, error) {
	SchemaProduction, err := bigquery.InferSchema(&entities.Production{})

	if err != nil {
		return nil, errors.Wrap(err, ErrorInferSchemaProductionBQ.Error())
	}

	tableInstance, err := bigqueryservice.GetTable("production", SchemaProduction)

	Production := tableInstance.Read(ctx)
	var ProductionFormatted []bigquery.Value
	for {
		var row entities.Production
		err := Production.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		ProductionFormatted = append(ProductionFormatted, row)
	}

	response, _ := json.Marshal(ProductionFormatted)

	log.Println("response.length: ", len(response))

	return response, nil
}

func DeleteBQProduction(ProductionId string) error {
	query := bigqueryservice.GetClient(false).Query("DELETE FROM productions.production WHERE productionId = '" + ProductionId + "'")

	_, err := query.Read(ctx)

	if err != nil {
		return errors.Wrap(err, ErrorDeleteProductionBQOnBigQuery.Error())
	}

	return nil
}
