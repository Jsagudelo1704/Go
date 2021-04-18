package services

/*
//Este componente se crea con el fin de manejar todo el flujo del servicio y desacoplar responsabilidades
//Como buena practica de desarrollo se crea un controlador para cada servicio
*/
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jsagudelo1704/Go/pkg/mutant"
	"github.com/Jsagudelo1704/Go/pkg/stats"
	"github.com/Jsagudelo1704/Go/pkg/validate"
	"github.com/Jsagudelo1704/Go/structs"
)

//Funcion principal del controlador para el flujo del servicio mutant\
func mutantController(w http.ResponseWriter, r *http.Request) {
	var rta structs.Respuesta

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

		// Verificar si las cadenas ingresadas son validas
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

		//Verificar si la cadena recibida ya ha sido procesada para no volverla a analizar
		//Este punto se debe evaluar con bastante análisis ya que puede ser mas eficiente procesar la cadena que realizar la consulta a la bd
		result := stats.GetDna(newDna.Dna)
		if result != "" {
			w.WriteHeader(http.StatusForbidden)
			if result == "Mutante" {
				w.WriteHeader(http.StatusOK)
			}
			rta.Msg = "Dna ya procesado"
			rta.Result = result
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(j)
			return
		}

		//Verificar si el dna ingresado es de un mutante
		retorno := mutant.Ismutant(newDna.Dna)

		// De forma asíncrona se envía la petición para guardar el documento en la bd
		go stats.InsertDna(newDna.Dna, retorno)

		//Validar retorno de la funcion Ismutant()
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
