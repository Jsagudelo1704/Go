package validate

import (
	"encoding/json"
	"net/http"

	"github.com/Jsagudelo1704/Go/structs"
)

func validateBody(req *http.Request) (rta structs.Respuesta, dna structs.Dna) {

	rta.Msg = ""
	rta.Result = ""
	err := json.NewDecoder(req.Body).Decode(&dna)
	if err != nil {
		rta.Msg = "Error al interpretar el cuerpo de la petición"
		rta.Result = "Peticion fallida"
	}

	return
}
