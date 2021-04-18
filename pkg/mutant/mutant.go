package mutant

/*
// Este pkg contiene las funciones correspondientes al metodo mutant\
// Estan funciones de validacion para este contexto, así como las funciones propias del negocio
*/

import (
	"regexp"
	"strings"
)

var tam, i, j int
var matriz [][]string

const CONSECUTIVOS = 4
const NUMSEQ = 2

//Funcion que retorna si el dna ingreado es mutante. Para esta validacion se buscan 2 (NUMSEQ) o más secuencias de 4 (CONSECUTIVOS) caracteres iguales.
//Para esta funcion se crea una matriz con en scope global con el fin de que pueda ser accedida en las diversas funciones privadas de busqueda
func Ismutant(dna []string) bool {
	i = 0
	j = 0
	cont := 0
	tam = len(dna)

	vector := make([]string, tam)

	//Para llenar la matriz lo que se hace es anexar cada arreglo a la matriz. Cada arreglo contiene un vector de caracteres
	//Para permitir retorno tempranos, se implementa la busqueda 1 que valida en primera instancia los cadenas ya recibidas.
	for _, cadena := range dna {
		cadena = strings.ToUpper(cadena)
		cont = cont + busq1(cadena)
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
			cont = cont + busq2()
			if cont >= NUMSEQ {
				return true
			}
			cont = cont + busq3()
			if cont >= NUMSEQ {
				return true
			}
			cont = cont + busq4()
			if cont >= NUMSEQ {
				return true
			}
		}
	}
	return false
}

//Busqueda inicial con el fin de aprovechar que se está leyendo cada cadena del array
//Un resultado mayor a 1 en esta busqueda permitira que el algoritmo termine de forma pronta y exitosa consumiendo así menos recursos de procesamiento
func busq1(s string) (count int) {
	count = strings.Count(s, "AAAA") + strings.Count(s, "CCCC") + strings.Count(s, "GGGG") + strings.Count(s, "TTTT")
	return count
}

//Busqueda en direccion vertical hacia abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud vertical suficiente
func busq2() (count int) {
	if (i + CONSECUTIVOS) <= tam {
		if matriz[i][j] == matriz[i+1][j] && matriz[i][j] == matriz[i+2][j] && matriz[i][j] == matriz[i+3][j] {
			count++
		}
	}
	return count
}

//Busqueda en direccion derecha y abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq3() (count int) {
	if (i+CONSECUTIVOS) <= tam && (j+CONSECUTIVOS) <= tam {
		if matriz[i][j] == matriz[i+1][j+1] && matriz[i][j] == matriz[i+2][j+2] && matriz[i][j] == matriz[i+3][j+3] {
			count++
		}
	}
	return count
}

//Busqueda en direccion derecha y arriba
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq4() (count int) {
	if (i-CONSECUTIVOS+1) >= 0 && (j+CONSECUTIVOS) <= tam {
		if matriz[i][j] == matriz[i-1][j+1] && matriz[i][j] == matriz[i-2][j+2] && matriz[i][j] == matriz[i-3][j+3] {
			count++
		}
	}
	return count
}

//Aplicar validaciones generales a la cadena recibida para verificar que cumple con las condiciones
//Estas se obtienen de la especificacion entregada
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
