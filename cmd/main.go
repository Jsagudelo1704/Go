package main

import (
	"log"
	"net/http"

	"github.com/Jsagudelo1704/Go/cmd/services"
)

const rutabase = "/api"

func main() {

	services.SetupRoutes(rutabase)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
