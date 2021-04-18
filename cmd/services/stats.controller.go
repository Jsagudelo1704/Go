package services

/*
//Este componente se crea con el fin de manejar todo el flujo del servicio y desacoplar responsabilidades
//Como buena practica de desarrollo se crea un controlador para cada servicio
*/

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jsagudelo1704/Go/pkg/stats"
	"github.com/Jsagudelo1704/Go/structs"
)

//Funcion principal del controlador para el flujo del servicio mutant\
func statsController(w http.ResponseWriter, r *http.Request) {
	var rta structs.Respuesta
	var statsresp structs.Stats

	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		rta, statsresp = stats.GetStats()
		if rta.Result != "" {
			w.WriteHeader(http.StatusInternalServerError)
			j, err := json.Marshal(rta)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(j)
			return
		}

		w.WriteHeader(http.StatusOK)
		j, err := json.Marshal(statsresp)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)
		return

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		rta.Msg = "Metodo no permitido"
		rta.Result = "Peticion Fallida"
		j, err := json.Marshal(rta)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(j)
		return
	}

}
