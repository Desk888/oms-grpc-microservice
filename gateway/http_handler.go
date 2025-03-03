package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type handler struct {
	// gateway
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/api/customers/{customerID}/orders", h.HandleCreateOrder).Methods("POST")
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// call service
}