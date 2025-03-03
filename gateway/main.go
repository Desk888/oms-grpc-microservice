package main

import (
	"log"
	"net/http"

	"github.com/Desk888/common"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Desk888/common/api"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
	OrderServiceAddr = common.EnvString("ORDER_SERVICE_ADDR", ":3000")
)

func main() {
	conn, err := grpc.Dial(OrderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to order service: ", err)
	}
	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)
	log.Println("Dialing order service on", OrderServiceAddr)

	router := mux.NewRouter()
	handler := NewHandler(c)
	handler.RegisterRoutes(router)

	log.Println("Starting server on", httpAddr)
	
	if err := http.ListenAndServe(httpAddr, router); err != nil {
		log.Fatal("failed to listen and serve: ", err)
	}
}