package validate

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jsagudelo1704/Go/structs"
)

func validateBody(req *http.Request) (resp http.ResponseWriter) {

	var dna structs.Dna
	var rta structs.Respuesta

	err := json.NewDecoder(req.Body).Decode(&dna)
	if err != nil {
		log.Print(err)
		resp.WriteHeader(http.StatusBadRequest)

		rta.Msg = "Error al interpretar el cuerpo de la petición"
		rta.Result = "Peticion fallida"
	}

	j, err := json.Marshal(rta)
	if err != nil {
		log.Fatal(err)
	}

	_, err = resp.Write(j)
	if err != nil {
		log.Fatal(err)
	}

	return resp

}
