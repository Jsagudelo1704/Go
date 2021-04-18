package stats

/*
// Este pkg contiene las funciones correspondientes al metodo stats\
*/
import (
	"context"
	"log"
	"time"

	"github.com/Jsagudelo1704/Go/database"
	"github.com/Jsagudelo1704/Go/structs"
	"go.mongodb.org/mongo-driver/bson"
)

//Insertar un nuevo documento en la collecion con la informacion de la cadena validada
func InsertDna(dna []string, ismutant bool) {

	var document structs.DnaBd
	document.Dna = dna
	document.Result = "Humano"
	if ismutant {
		document.Result = "Mutante"
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, col := database.ConnectDB(ctx)
	_, err := col.InsertOne(ctx, document)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

}

//Traer las estadisticas de la collecion dnaverified
func GetStats() (rta structs.Respuesta, stats structs.Stats) {
	var ret1, ret2 bool
	stats.CountMutant, ret1 = getcount("Mutante")
	stats.CountHuman, ret2 = getcount("Humano")

	if !ret1 || !ret2 {
		rta.Msg = "Error Interno con la BD"
		rta.Result = "Peticion Fallida"
		return
	}

	//Se valida que el denominador de la operación no sea 0, en ese caso se deja el ratio igual a la cantidad de mutantes
	if stats.CountMutant > 0 {
		stats.Ratio = (float32(stats.CountMutant) / float32(stats.CountHuman))
	} else {
		stats.Ratio = float32(stats.CountMutant)
	}

	return
}

//Obtener un documento especifico de la coleccion
func GetDna(dna []string) string {
	var document structs.DnaBd
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, col := database.ConnectDB(ctx)
	col.FindOne(ctx, bson.M{"dna": dna}).Decode(&document)

	defer client.Disconnect(ctx)
	return document.Result
}

//funcion para obtener la cantidad de documentos que cumplen con el parametro de busqueda
func getcount(s string) (int, bool) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, col := database.ConnectDB(ctx)
	count, err := col.CountDocuments(ctx, bson.M{"result": s})
	if err != nil {
		return 0, false
	}
	defer client.Disconnect(ctx)
	return int(count), true
}
