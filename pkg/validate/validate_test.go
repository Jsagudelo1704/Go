package validate

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Jsagudelo1704/Go/structs"
)

func TestValidateBody(t *testing.T) {
	tests := map[string]struct {
		payload        string
		expectedResult string
		expectedMsg    string
	}{
		"Deberia retornar Peticion fallida": {
			payload:        `{"dna":[ "ttttgA" CTGCat", "aCaTGT", "TATAAc", "CagCTA", "cACTaA" ]}`,
			expectedResult: "Peticion fallida",
			expectedMsg:    "Error al interpretar el cuerpo de la petición",
		},
		"Deberia retornar vacio (exitoso)": {
			payload:        `{"dna":[ "ttttgA", "CTGCat", "aCaTGT", "TATAAc", "CagCTA", "cACTaA" ]}`,
			expectedResult: "",
			expectedMsg:    "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var rta structs.Respuesta
			r := ioutil.NopCloser(strings.NewReader(test.payload))
			rta, _ = ValidateBody(r)
			if test.expectedMsg != rta.Msg || test.expectedResult != rta.Result {
				t.Error("No se cumple la prueba de ValidateBody")
			}
		})
	}
}
