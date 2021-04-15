package structs

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
