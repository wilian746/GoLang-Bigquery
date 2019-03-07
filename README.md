# GOLang and Bigquery API #

This application is intended to make viability on [golang and bigquery](https://godoc.org/cloud.google.com/go/bigquery). Follow the instructions below to run the project

## Requirements

- [GOLANG](https://golang.org/dl/)
- [Account in Google](https://console.cloud.google.com/bigquery)
- [GOOGLE_APPLICATION_CREDENTIALS](https://cloud.google.com/docs/authentication/production?hl=pt-br#obtaining_and_providing_service_account_credentials_manually)

### First steps

- Clone this project

- Inside of project run install all requirements

```bash
  $ go get -t -v -u ./...
```
- And run project

```bash
  $ go run main.go
```

### Requests

- Request to send Api in route POST `http://localhost:8071/production`

```json
  {
    "Weight": 12.34,
    "EquipmentId": "1234",
    "ProductionOrderId": "12345",
    "ProductAttributeId": "123456",
    "CompanyId": "1234567",
    "QrCodeId": "12345678",
    "Active": true,
    "CreatedBy": "123456789",
    "UpdatedBy": "1234567890"
  }
```

- Request to send Api in route GET with ID optional to GetOne (UUID válid) or GetAll if not send id `http://localhost:8071/production/?{UUID}`

- Request to send Api in route DELETE with ID `http://localhost:8071/production/{UUID}`

