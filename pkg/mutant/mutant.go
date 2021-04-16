package mutant

import (
	"regexp"
	"strings"
)

var tam, cont, i, j int
var matriz [][]string

const CONSECUTIVOS = 4
const NUMSEQ = 2

//Funcion que retorna true cuando el areglo de entrada contiene >= NUMSEQ secuencias de CONSECUTIVOS caracteres iguales
// Para esta funcion se crea una matriz con en scope global con el fin de que pueda ser accedidad en las diversas funciones de busqueda
func Ismutant(dna []string) bool {
	i = 0
	j = 0
	cont = 0
	tam = len(dna)

	vector := make([]string, tam)

	//Para llenar la matriz lo que se hace es anexar cada arreglo a la matriz. Cada arreglo contiene un vector de caracteres
	for _, cadena := range dna {
		cadena = strings.ToUpper(cadena)
		busq1(cadena)
		if cont >= NUMSEQ {
			return true
		}
		vector = strings.Split(cadena, "")
		matriz = append(matriz, vector)
	}

	//En cada posición de la matriz realizamos las tres posibles busquedas para verificar las secuencias
	//Luego de cada búsqueda se valida el contador, con el fin de no hacer operaciones innecesarias
	for i = 0; i < tam; i++ {
		for j = 0; j < tam; j++ {
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

//Busqueda inicial con el fin de aprovechar que se está leyendo cada cadena del array
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
func IsDnaValid(dna []string) (msg string, result string) {
	//Validar que la cadena solo contenga los caracteres A,C,G,T
	var regex_dna = regexp.MustCompile("^([ACGT]*)$")
	tam = len(dna)
	result = ""
	msg = ""

	for _, cadena := range dna {
		reg := regex_dna.FindStringSubmatch(strings.ToUpper(cadena))
		if reg == nil {
			msg = "Caracter invalido en la cadena"
			result = "Peticion fallida"
			return
		}
		//Validar que la matriz sea simétrica
		if len(cadena) != tam {
			msg = "Matriz no es simetrica"
			result = "Peticion fallida"
			return
		}
	}
	return
}
