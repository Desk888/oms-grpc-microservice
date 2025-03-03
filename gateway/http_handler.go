package main

import (
	"net/http"
	"github.com/gorilla/mux"
	pb "github.com/Desk888/common/api"
	"github.com/Desk888/common"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/api/customers/{customerID}/orders", h.HandleCreateOrder).Methods("POST")
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var items[]*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, "failed to read request body")
		return
	}

	vars := mux.Vars(r)
	customerID := vars["customerID"]

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerID,
		Items: items,
	})
}