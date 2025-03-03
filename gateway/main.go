package main

import (
	"net/http"
	"log"
	"github.com/Desk888/common"
	_ "github.com/joho/godotenv/autoload"
	"github.com/gorilla/mux"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	router := mux.NewRouter()
	handler := NewHandler()
	handler.RegisterRoutes(router)

	log.Println("starting server on", httpAddr)
	
	if err := http.ListenAndServe(httpAddr, router); err != nil {
		log.Fatal("failed to listen and serve: ", err)
	}
}