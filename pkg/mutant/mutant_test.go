package mutant

import (
	"encoding/json"
	"testing"

	"github.com/Jsagudelo1704/Go/structs"
)

//Funcion para probar el algoritmo que valida un dna
//Se plantean 3 escenarios para abarscar la mayor cobertura posible
//1. dna de Mutante. En este caso el escesario cumple con la busqueda 1
//2. dna de Humano. En este caso no encuentra ninguna secuencia
//3. dna de Humano. En este caso encuentra 1 secuencia

func TestIsmutant(t *testing.T) {

	tests := map[string]struct {
		dna            string
		expectedResult bool
	}{
		"Deberia retornar - Mutante": {
			dna:            `{"dna":["ttttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","cACTaA"]}`,
			expectedResult: true,
		},

		"Deberia retornar - Mutante2": {
			dna:            `{"dna":["ttttgA","CTGCat","aCaTGT","TAcAAc","CagCTA","cACTaA"]}`,
			expectedResult: true,
		},

		"Deberia retornar - Mutante3": {
			dna:            `{"dna":["tgttgA","CTGCat","aCaTGc","TAcAcc","CagCTA","cACTaA"]}`,
			expectedResult: true,
		},

		"Deberia retornar - Mutante4": {
			dna:            `{"dna":["tGttgA","tGGCat","tgaTGc","TgcAcc","CagCTA","cACTaA"]}`,
			expectedResult: true,
		},

		"Deberia retornar - Humano 1": {
			dna:            `{"dna":["tcttgA","CTGCat","aCaTGc","TAtAcc","CagCTA","cACTaA"]}`,
			expectedResult: false,
		},
		"Deberia retornar - Humano 2": {
			dna:            `{"dna":["tcttgA","CTGCat","aCaTGt","TAcAcc","CagCTA","cACTaA"]}`,
			expectedResult: false,
		},
		"Deberia retornar - Humano 3": {
			dna:            `{"dna":["tc","CT"]}`,
			expectedResult: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var dna structs.Dna
			json.Unmarshal([]byte(test.dna), &dna)
			retorno := Ismutant(dna.Dna)
			if test.expectedResult != retorno {
				t.Error("No se cumple la prueba de Ismutant")
			}
		})
	}
}

//Funcion para probar el algoritmo que valida una cadena dna antes de enviarla a verificar si es mutante
//Se plantean 3 escenarios para cubrir todos los casos posible:
//1. cadenas con caracteres que no esten dentro de la regex definida
//2. cadenas que no permitan armar una matriz simetrica
//3. una cadena que sea simetrica y contenga solo los caracteres validos - exitoso
func TestIsDnaValid(t *testing.T) {
	var dna structs.Dna
	tests := map[string]struct {
		dna            string
		expectedResult string
		expectedMsg    string
	}{
		"Deberia retornar Caracter invalido en la cadena": {
			dna:            `{"dna":["ftttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","cACTaA"]}`,
			expectedResult: "Peticion fallida",
			expectedMsg:    "Caracter invalido en la cadena",
		},
		"Deberia retornar Caracter invalido en la cadena 1": {
			dna:            `{"dna":["6tttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","cACTaA"]}`,
			expectedResult: "Peticion fallida",
			expectedMsg:    "Caracter invalido en la cadena",
		},
		"Deberia retornar Caracter invalido en la cadena 2": {
			dna:            `{"dna":["atttgA","CTGCat","aCaTGT","TAa+Ac","CagCTA","cACTaA"]}`,
			expectedResult: "Peticion fallida",
			expectedMsg:    "Caracter invalido en la cadena",
		},
		"Deberia retornar Matriz no es simetrica": {
			dna:            `{"dna":["ttttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","cACTa"]}`,
			expectedResult: "Peticion fallida",
			expectedMsg:    "Matriz no es simetrica",
		},
		"Deberia retornar Matriz no es simetrica 1": {
			dna:            `{"dna":["ttttgA","CTGCat"]}`,
			expectedResult: "Peticion fallida",
			expectedMsg:    "Matriz no es simetrica",
		},
		"Deberia retornar vacio - exitoso": {
			dna:            `{"dna":["atttgA","CTGCat","aCaTGT","TAaAAc","CagCTA","cACTaA"]}`,
			expectedResult: "",
			expectedMsg:    "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			json.Unmarshal([]byte(test.dna), &dna)
			msg, rta := IsDnaValid(dna.Dna)
			if test.expectedResult != rta || test.expectedMsg != msg {
				t.Error("No se cumple la prueba de IsDnaValid")
			}
		})
	}
}
