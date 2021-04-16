package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jsagudelo1704/Go/pkg/mutant"
	"github.com/Jsagudelo1704/Go/pkg/validate"
	"github.com/Jsagudelo1704/Go/structs"
)

var rta structs.Respuesta

func mutantController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")

		//Validacion del body del request
		rta, newDna := validate.ValidateBody(r)
		if rta.Msg != "" {
			w.WriteHeader(http.StatusBadRequest)
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(j)

			return
		}

		// Varificar si las cadenas ingresadas son validas
		rta.Msg, rta.Result = mutant.IsDnaValid(newDna.Dna)

		if rta.Msg != "" {
			w.WriteHeader(http.StatusBadRequest)
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(j)

			return
		}

		//Verificar si el dna ingresado es de un mutante
		retorno := mutant.Ismutant(newDna.Dna)
		if !retorno {
			w.WriteHeader(http.StatusForbidden)
			rta.Msg = "Dna procesado"
			rta.Result = "Humano"
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(j)
			return
		}

		w.WriteHeader(http.StatusOK)
		rta.Msg = "Dna procesado"
		rta.Result = "Mutante"
		j, err := json.Marshal(rta)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)
		return

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		rta.Msg = "Metodo no permitido"
		rta.Result = "Petición Fallida"
		j, err := json.Marshal(rta)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)
		return
	}

}
