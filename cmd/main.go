package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jsagudelo1704/Go/cmd/services"
)

const rutabase = "/api"

func main() {

	services.SetupRoutes(rutabase)
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
	}
	return ":" + port
}
