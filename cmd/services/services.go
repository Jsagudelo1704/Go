package services

import (
	"fmt"
	"net/http"
)

const rutaMutant = "mutant/"

//const rutaStats = "stats"

func handleMutant(w http.ResponseWriter, r *http.Request) {
	mutantController(w, r)

}

// SetupRoutes :
func SetupRoutes(rutabase string) {
	mutantHandler := http.HandlerFunc(handleMutant)
	//statsHandler := http.HandlerFunc(handleStats)
	http.Handle(fmt.Sprintf("%s/%s", rutabase, rutaMutant), mutantHandler)
	//http.Handle(fmt.Sprintf("%s/%s", rutabase, rutaStats), statsHandler)
}
