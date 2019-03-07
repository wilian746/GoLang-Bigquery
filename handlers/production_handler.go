package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/GoLang-Bigquery/controllers"

	"github.com/google/uuid"
)

type ProductionHandler struct {
}

type ErrorModel struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func NewProductionHandler() *ProductionHandler {
	return &ProductionHandler{}
}

func (h *ProductionHandler) ProductionHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")

	log.Println("METHOD: ", req.Method)
	log.Println("URL: ", req.URL.Path)
	switch req.Method {
	case http.MethodOptions:
		h.Options(resp, req)
	case http.MethodPost:
		h.Post(resp, req)
	case http.MethodGet:
		h.Get(resp, req)
	case http.MethodDelete:
		h.Delete(resp, req)
	default:
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Status Method Not Allowed"))
	}
}

func (h *ProductionHandler) Options(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	resp.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	resp.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	resp.WriteHeader(http.StatusOK)
}

func (h *ProductionHandler) Post(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if len(body) == 0 {
		log.Printf("failed to read post request: %s", err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("received post body: %s", body)

	production, err := controllers.SaveBQProduction(body)

	if err != nil {
		if err.Error() == "{SaveProductionBQ} Error while trying to save ProductionBQ on BigQuery" {
			log.Printf("error when trying to save Production: %s", err)
			resp.WriteHeader(http.StatusConflict)
			return
		}
		log.Printf("error when trying to save Production: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(production)

	resp.WriteHeader(http.StatusCreated)
	_, err = resp.Write(bytes)
}

func (h *ProductionHandler) Get(resp http.ResponseWriter, req *http.Request) {
	params := strings.Split(req.URL.Path, "/")

	productionId := params[len(params)-1]

	if productionId != "" {
		_, err := uuid.Parse(productionId)

		if err != nil {
			log.Print(err)
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		production, err := controllers.GetOneBQProduction(productionId)

		if err != nil {
			log.Printf("error when trying to get Production: %s", err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		if production == nil {
			resp.WriteHeader(http.StatusNotFound)
			response, _ := json.Marshal("Production Not Found")
			resp.Write(response)
			return
		}

		resp.WriteHeader(http.StatusOK)
		_, err = resp.Write(production)
	} else {
		production, err := controllers.GetAllBQProduction()

		if err != nil {
			log.Printf("error when trying to get Production: %s", err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		if production == nil {
			resp.WriteHeader(http.StatusNotFound)
			response, _ := json.Marshal("Production Not Found")
			resp.Write(response)
			return
		}

		resp.WriteHeader(http.StatusOK)
		resp.Write(production)
	}
}

func (h *ProductionHandler) Delete(resp http.ResponseWriter, req *http.Request) {
	params := strings.Split(req.URL.Path, "/")

	productionId := params[len(params)-1]

	_, err := uuid.Parse(productionId)

	if err != nil {
		log.Print(err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controllers.DeleteBQProduction(productionId)

	if err != nil {
		log.Print(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(204)
}
