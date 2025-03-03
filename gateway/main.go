package main

import (
	"net/http"
	"log"
	"github.com/Desk888/common"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.RegisterRoutes(mux)

	log.Println("starting server on", httpAddr)
	
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to listen and serve: ", err)
	}
}``