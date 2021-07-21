package stats

import (
	"encoding/json"
	"testing"

	"github.com/Jsagudelo1704/Go/structs"
)

var testconf structs.MongoConfiguration

func init() {
	testconf = structs.MongoConfiguration{
		Server:     "mongodb://localhost:27017",
		Database:   "Mutant",
		Collection: "Test",
	}
	conf = structs.MongoConfiguration{
		Server:     "mongodb://localhost:27017",
		Database:   "Mutant",
		Collection: "dnaverified",
	}
}

//Funcion para probar la consulta de un dna en la BD
//Se plantean 3 casos para cubrir todos los escenarios posible:
//1. dna de Mutante existene
//2. dna de Humano existente
//3. dna NO existente
func TestGetDna(t *testing.T) {
	var dna structs.Dna
	tests := map[string]struct {
		dna            string
		expectedResult string
	}{
		"Deberia retornar dna existente - Mutante": {
			dna:            `{"dna":["ttttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","cACTaA"]}`,
			expectedResult: "Mutante",
		},
		"Deberia retornar dna existente - Humano": {
			dna:            `{"dna":["ttttgA","CTGCat","aCaTGT","TATAAc","CagCTA","cACTaA"]}`,
			expectedResult: "Humano",
		},
		"Deberia retornar vacio ya que NO existe dna": {
			dna:            `{"dna":["acttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","cACTaA"]}`,
			expectedResult: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			json.Unmarshal([]byte(test.dna), &dna)
			result := GetDna(dna.Dna)
			if test.expectedResult != result {
				t.Error("No se cumple la prueba de GetDna")
			}
		})
	}
}

//Funcion para probar la consulta a la BD para traer un documento
//Se plantean los siguientes dos casos para cubrir los posibles escenarios:
//1. dna existente de Mutante
//2. dna existente de Humano
//NO se plantea un escenario cuando la BD retorna error ya que esta funcion es asincrona y este caso no detiene la ejecucion.
//Un documento que no se grabe en la BD no afecta la funcionalidad, ya que la cadena se evalua igualmente
func TestInsertDna(t *testing.T) {
	var dna structs.Dna
	tests := map[string]struct {
		dna      string
		ismutant bool
		err      error
	}{
		"Deberia guardar dna - Mutante": {
			dna:      `{"dna":["ttttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","ccCcaA"]}`,
			ismutant: true,
			err:      nil,
		},
		"Deberia guardar dna - Humano": {
			dna:      `{"dna":["ttttgA","CTGCat","aCaTGT","TATAAc","CagCTA","cACTaA"]}`,
			ismutant: false,
			err:      nil,
		},
		"Deberia retornar error por cadena no JSON": {
			dna:      `{"dna":"ttttgA,"CTGCat","aCaTGT","TATAAc","CagCTA","cACTaA"]`,
			ismutant: false,
			err:      nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			json.Unmarshal([]byte(test.dna), &dna)
			err := InsertDna(dna.Dna, test.ismutant)
			if test.err != err {
				t.Error("No se cumple la prueba de InsertDna")
			}
		})
	}
}

//Funcion para probar la consulta de un dna en la BD
//Se plantean 3 casos para cubrir todos los escenarios posible:
//1. dna de Mutante existene
//2. dna de Humano existente
//3. dna NO existente
func TestGetStats(t *testing.T) {
	rta, stats := GetStats()
	if rta.Result != "" {
		if rta.Msg == "" {
			t.Error("No se cumple la prueba de GetStats, datos de respuesta no congruentes")
		}
		if stats.Ratio > 0 {
			t.Error("No se cumple la prueba de GetStats, ratio de respuesta no congruente")
		}
		if stats.CountHuman > 0 && stats.CountMutant > 0 {
			t.Error("No se cumple la prueba de GetStats, Result deberia estar vacio")
		}
	} else {
		if stats.CountHuman == 0 && stats.Ratio != float32(stats.CountMutant) {
			t.Error("No se cumple la prueba de GetStats, Ratio incorrecto")
		}
		if stats.CountHuman > 0 && stats.Ratio != (float32(stats.CountMutant)/float32(stats.CountHuman)) {
			t.Error("No se cumple la prueba de GetStats, Ratio no corresponde a conteos")
		}
	}

}
