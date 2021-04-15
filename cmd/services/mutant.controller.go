package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jsagudelo1704/Go/pkg/mutant"
	"github.com/Jsagudelo1704/Go/structs"
)

var rta structs.Respuesta

func mutantController(w http.ResponseWriter, r *http.Request) (response http.ResponseWriter) {

	switch r.Method {

	case http.MethodPost:

		//Validacion del body del request
		rta, newDna := validate.validateBody(r)
		if rta.Msg != "" {
			response.WriteHeader(http.StatusBadRequest)
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			response.Write(j)

			return
		}

		// Varificar si las cadenas ingresadas son validas
		rta = mutant.IsDnaValid(newDna.Dna)
		if rta.Msg != "" {
			response.WriteHeader(http.StatusBadRequest)
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			response.Write(j)

			return
		}

		//Verificar si el dna ingresado es de un mutante
		retorno := mutant.Ismutant(newDna.Dna)
		if !retorno {
			response.WriteHeader(http.StatusForbidden)
			rta.Msg = "Dna procesado"
			rta.Result = "Humano"
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			response.Write(j)
			return
		}

		response.WriteHeader(http.StatusOK)
		rta.Msg = "Dna procesado"
		rta.Result = "Mutante"
		j, err := json.Marshal(rta)
		if err != nil {
			log.Fatal(err)
		}
		response.Write(j)
		return

	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		rta.Msg = "Metodo no permitido"
		rta.Result = "Petición Fallida"
		j, err := json.Marshal(rta)
		if err != nil {
			log.Fatal(err)
		}
		response.Write(j)
		return
	}
}
