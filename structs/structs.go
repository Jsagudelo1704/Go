package structs

/*
// Este componente se encarga de agrupar las diferentes estructuras usadas en el proyecto.
// Para realizar la inicialización de alguna estructura, este pkg debe ser importado
*/

//Estructura del objeto JSON que se recibe en el body del servicio mutant/
type Dna struct {
	Dna []string `json:"dna"`
}

//Estructura de respuesta genércia para el servicio mutant/
type Respuesta struct {
	Msg    string `json:"message"`
	Result string `json:"result"`
}

//Estructura del JSON de respuesta para el servicio de stats/
type Stats struct {
	CountMutant int     `json:"count_mutant_dna"`
	CountHuman  int     `json:"count_human_dna"`
	Ratio       float32 `json:"ratio"`
}

//Estructura JSON para guardar en la BD
type DnaBd struct {
	Dna    []string `json:"dna" bson:"dna"`
	Result string   `json:"result" bson:"result"`
}

//Estructura para el archivo de configuracion de la BD
type Configuration struct {
	Environment string
	Mongo       MongoConfiguration
}

// Tipo configuracion para el archivo de configuracion de la BD
type MongoConfiguration struct {
	Server     string
	Database   string
	Collection string
}

type MatrizData struct {
	Matriz [][]string
	Tamano int
	Pos_i  int
	Pos_j  int
}
