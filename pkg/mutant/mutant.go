package mutant

/*
// Este pkg contiene las funciones correspondientes al metodo mutant\
// Estan funciones de validacion para este contexto, así como las funciones propias del negocio
*/

import (
	"regexp"
	"strings"

	"github.com/Jsagudelo1704/Go/structs"
)

const CONSECUTIVOS = 4
const NUMSEQ = 2

//Funcion que retorna si el dna ingreado es mutante. Para esta validacion se buscan 2 (NUMSEQ) o más secuencias de 4 (CONSECUTIVOS) caracteres iguales.
//Para esta funcion se crea una matriz con en scope global con el fin de que pueda ser accedida en las diversas funciones privadas de busqueda
func Ismutant(dna []string) bool {

	matrizData := structs.MatrizData{}

	matrizData.Pos_i = 0
	matrizData.Pos_j = 0
	matrizData.Tamano = len(dna)

	cont := 0

	vector := make([]string, matrizData.Tamano)

	//Para llenar la matriz lo que se hace es anexar cada arreglo a la matriz. Cada arreglo contiene un vector de caracteres
	//Para permitir retorno tempranos, se implementa la busqueda 1 que valida en primera instancia los cadenas ya recibidas.
	for _, cadena := range dna {
		cadena = strings.ToUpper(cadena)
		cont = cont + busq1(cadena)
		if cont >= NUMSEQ {
			return true
		}
		vector = strings.Split(cadena, "")
		matrizData.Matriz = append(matrizData.Matriz, vector)
	}

	//En cada posición de la matriz realizamos las tres posibles busquedas para verificar las secuencias
	//Luego de cada búsqueda se valida el contador, con el fin de no hacer operaciones innecesarias
	for matrizData.Pos_i = 0; matrizData.Pos_i < matrizData.Tamano; matrizData.Pos_i++ {
		for matrizData.Pos_j = 0; matrizData.Pos_j < matrizData.Tamano; matrizData.Pos_j++ {
			cont = cont + busq2(matrizData)
			if cont >= NUMSEQ {
				return true
			}
			cont = cont + busq3(matrizData)
			if cont >= NUMSEQ {
				return true
			}
			cont = cont + busq4(matrizData)
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
func busq2(matrizData structs.MatrizData) (count int) {
	if (matrizData.Pos_i + CONSECUTIVOS) <= matrizData.Tamano {
		if matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i+1][matrizData.Pos_j] &&
			matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i+2][matrizData.Pos_j] &&
			matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i+3][matrizData.Pos_j] {
			count++
		}
	}
	return count
}

//Busqueda en direccion derecha y abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq3(matrizData structs.MatrizData) (count int) {
	if (matrizData.Pos_i+CONSECUTIVOS) <= matrizData.Tamano && (matrizData.Pos_j+CONSECUTIVOS) <= matrizData.Tamano {
		if matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i+1][matrizData.Pos_j+1] &&
			matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i+2][matrizData.Pos_j+2] &&
			matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i+3][matrizData.Pos_j+3] {
			count++
		}
	}
	return count
}

//Busqueda en direccion derecha y arriba
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq4(matrizData structs.MatrizData) (count int) {
	if (matrizData.Pos_i-CONSECUTIVOS+1) >= 0 && (matrizData.Pos_j+CONSECUTIVOS) <= matrizData.Tamano {
		if matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i-1][matrizData.Pos_j+1] &&
			matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i-2][matrizData.Pos_j+2] &&
			matrizData.Matriz[matrizData.Pos_i][matrizData.Pos_j] == matrizData.Matriz[matrizData.Pos_i-3][matrizData.Pos_j+3] {
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
	var tam = len(dna)
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
