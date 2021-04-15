package services

import (
	"net/http"

	"github.com/Jsagudelo1704/Go/pkg/mutant"
)

func mutantController(w http.ResponseWriter, r *http.Request) (response http.ResponseWriter, request *http.Request) {
	switch r.Method {
	case http.MethodPost:

		//Validacion del body del request
		response = validate.validateBody(r)

		// Varificar si las cadenas ingresadas son validas
		retorno, msjretorno := mutant.IsDnaValid(newDna.Dna)
		if !retorno {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(msjretorno))
			return
		}

		//Verificar si el dna ingresado es de un mutante
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
