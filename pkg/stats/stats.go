package stats

/*
// Este pkg contiene las funciones correspondientes al metodo stats\
// Estan funciones de validacion para este contexto, así como las funciones propias del negocio
*/
import (
	"github.com/Jsagudelo1704/Go/database"
	"github.com/Jsagudelo1704/Go/structs"
)

func InsertDna() {
	cli, db, col := database.ConnectDB()

}

func GetStats() structs.Stats {
	cli, db, col := database.ConnectDB()

	return structs.Stats{}
}
