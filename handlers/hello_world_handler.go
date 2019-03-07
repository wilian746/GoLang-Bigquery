package handlers

import (
	"net/http"
)

type HelloWorldHandler struct {
}

func NewHelloWorldHandler() *HelloWorldHandler {
	return &HelloWorldHandler{}
}

func (h *HelloWorldHandler) HelloWorldHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")

	switch req.Method {
	case http.MethodOptions:
		h.Options(resp, req)
	case http.MethodGet:
		h.Get(resp, req)
	default:
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Status Method Not Allowed"))
	}
}

func (h *HelloWorldHandler) Options(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	resp.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	resp.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	resp.WriteHeader(http.StatusOK)
}

func (h *HelloWorldHandler) Get(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Hello World!"))
}
