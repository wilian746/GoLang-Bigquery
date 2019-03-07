package entities

import "time"

type Production struct {
	ProductionId       string    `bigquery:"productionId" json:"productionId" bson:"productionId"`
	Weight             float32   `bigquery:"weight" json:"weight" bson:"weight"`
	EquipmentId        string    `bigquery:"equipmentId" json:"equipmentId" bson:"equipmentId"`
	ProductionOrderId  string    `bigquery:"productionOrderId" json:"productionOrderId" bson:"productionOrderId"`
	ProductAttributeId string    `bigquery:"productAttributeId" json:"productAttributeId" bson:"productAttributeId"`
	CompanyId          string    `bigquery:"companyId" json:"companyId" bson:"companyId"`
	QrCodeId           string    `bigquery:"qrCodeId" json:"qrCodeId" bson:"qrCodeId"`
	Active             bool      `bigquery:"active" json:"active" bson:"active"`
	CreatedBy          string    `bigquery:"createdBy" json:"createdBy" bson:"createdBy"`
	UpdatedBy          string    `bigquery:"updatedBy" json:"updatedBy" bson:"updatedBy"`
	CreatedAt          time.Time `bigquery:"createdAt" json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time `bigquery:"updatedAt" json:"updatedAt" bson:"updatedAt"`
}
