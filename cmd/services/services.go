package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jsagudelo1704/Go/pkg/mutant"
)

const rutaMutant = "mutant/"

//const rutaStats = "stats"

type Cadenas struct {
	Dna []string `json:"dna"`
}

func handleMutant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		var newDna Cadenas
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error al leer la entrada"))
			return
		}

		err = json.Unmarshal(body, &newDna)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error al convertir la entrada"))
			return
		}

		// Validar si las cadenas ingresadas son validas
		retorno, msjretorno := mutant.IsDnaValid(newDna.Dna)
		if !retorno {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(msjretorno))
			return
		}

		//Verificar si el dna ingresado es mutante
		retorno = mutant.Ismutant(newDna.Dna)
		if !retorno {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Humano"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Mutante"))

		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método inválido"))
	}
}

// SetupRoutes :
func SetupRoutes(rutabase string) {
	mutantHandler := http.HandlerFunc(handleMutant)
	//statsHandler := http.HandlerFunc(handleProduct)
	http.Handle(fmt.Sprintf("%s/%s", rutabase, rutaMutant), mutantHandler)
	//http.Handle(fmt.Sprintf("%s/%s", rutabase, rutaStats), cors.Middleware(productHandler))
}
