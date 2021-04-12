package mutant

import (
	"regexp"
	"strings"
)

var tam, cont, i, j int
var matriz [][]string

const CONSECUTIVOS = 4
const NUMSEQ = 2

func Ismutant(dna []string) bool {
	i = 0
	j = 0
	cont = 0
	tam = len(dna)
	//matriz = make([][]string, tam)

	vector := make([]string, tam)

	for _, cadena := range dna {
		cadena = strings.ToUpper(cadena)
		busq1(cadena)
		if cont >= NUMSEQ {
			return true
		}
		vector = strings.Split(cadena, "")
		matriz = append(matriz, vector)
	}

	for i = 0; i < tam; i++ {
		for j = 0; j < tam; j++ {
			/*
				busq1(i, j, matriz)
				if cont >= NUMSEQ {
					return true
				}
			*/
			busq2()
			if cont >= NUMSEQ {
				return true
			}
			busq3()
			if cont >= NUMSEQ {
				return true
			}
			busq4()
			if cont >= NUMSEQ {
				return true
			}
		}
	}
	return false
}

//Busqueda en direccion horizontal a la derecha
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud horizontal suficiente
/*
func busq1(i int, j int, matriz [][]string) {
	if (j + CONSECUTIVOS) <= tam {
		if matriz[i][j] == matriz[i][j+1] && matriz[i][j] == matriz[i][j+2] && matriz[i][j] == matriz[i][j+3] {
			cont++
		}
	}
}
*/
//Busqueda
func busq1(s string) {
	cont = cont + strings.Count(s, "AAAA") + strings.Count(s, "CCCC") + strings.Count(s, "GGGG") + strings.Count(s, "TTTT")
}

//Busqueda en direccion vertical hacia abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud vertical suficiente
func busq2() {
	if (i + CONSECUTIVOS) <= tam {
		if matriz[i][j] == matriz[i+1][j] && matriz[i][j] == matriz[i+2][j] && matriz[i][j] == matriz[i+3][j] {
			cont++
		}
	}
}

//Busqueda en direccion derecha y abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq3() {
	if (i+CONSECUTIVOS) <= tam && (j+CONSECUTIVOS) <= tam {
		if matriz[i][j] == matriz[i+1][j+1] && matriz[i][j] == matriz[i+2][j+2] && matriz[i][j] == matriz[i+3][j+3] {
			cont++
		}
	}
}

//Busqueda en direccion derecha y arriba
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq4() {
	if (i-CONSECUTIVOS+1) >= 0 && (j+CONSECUTIVOS) <= tam {
		if matriz[i][j] == matriz[i-1][j+1] && matriz[i][j] == matriz[i-2][j+2] && matriz[i][j] == matriz[i-3][j+3] {
			cont++
		}
	}
}

//Aplicar validaciones generales a la cadena recibida para verificar que cumple con las condiciones
func IsDnaValid(dna []string) (ret bool, msj string) {
	//Validar que la cadena solo contenga los caracteres A,C,G,T
	var regex_dna = regexp.MustCompile("^([ACGT]*)$")
	tam = len(dna)
	ret = true
	msj = ""

	for _, cadena := range dna {
		reg := regex_dna.FindStringSubmatch(strings.ToUpper(cadena))
		if reg == nil {
			ret = false
			msj = "Caracter invalido en la cadena"
			return
		}
		//Validar que la matriz sea simétrica
		if len(cadena) != tam {
			ret = false
			msj = "Matriz no es simetrica"
			return
		}
	}
	return
}
