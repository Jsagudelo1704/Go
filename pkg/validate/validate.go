package validate

/*
// Este pkg contiene las funciones golbales de validacion de los servicios del proyecto
*/

import (
	"encoding/json"
	"io"

	"github.com/Jsagudelo1704/Go/structs"
)

// Funcion que permite validar el body de la peticion recibida. Se valida si este se puede convertir al objeto destino correctamente
func ValidateBody(r io.Reader) (rta structs.Respuesta, dna structs.Dna) {

	rta.Msg = ""
	rta.Result = ""
	err := json.NewDecoder(r).Decode(&dna)
	if err != nil {
		rta.Msg = "Error al interpretar el cuerpo de la petición"
		rta.Result = "Peticion fallida"
	}

	return
}
