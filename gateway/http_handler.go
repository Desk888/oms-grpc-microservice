package main

import (
	"net/http"
	"errors"
	"github.com/Desk888/common"
	pb "github.com/Desk888/common/api"
	"github.com/gorilla/mux"
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
	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	if err := ValidateItems(items) ; err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	customerID := vars["customerID"]

	order, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerID,
		Items:      items,
	})

	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	common.WriteJSON(w, http.StatusCreated, order)
}

func ValidateItems(items []*pb.ItemsWithQuantity) error {
	for _, item := range items {
		if item.Id == "" {
			return errors.New("item_id is required")
		}

		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
	}

	return nil
}