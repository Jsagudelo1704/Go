package mutant

import (
	"regexp"
	"strings"
	"fmt"
)

var len const 
var cont int
const consecutivos = 4
const numseq = 2

func Ismutant(dna []string) bool {
	i := 0
	j := 0
	cont = 0
	len = len(dna)
	
	matriz := make([][len]string, len, len)

	for _, cadena := range dna {
		//busq1(i, j, matriz)
		letras := strings.Split(cadena, "")
		j = 0
		for _, letra := range letras {
			matriz[i][j] = letra
			j++
		}
		i++
	}

	for i = 0; i < len; i++ {
		for j = 0; j < len; j++ {
			busq1(i, j, matriz)
			if cont >= numseq {
				return true
			}
			busq2(i, j, matriz)
			if cont >= numseq {
				return true
			}
			busq3(i, j, matriz)
			if cont >= numseq {
				return true
			}
			busq4(i, j, matriz)
			if cont >= numseq {
				return true
			}
		}
	}
	return false
}

//Busqueda en direccion horizontal a la derecha
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud horizontal suficiente
func busq1(i int, j int, matriz [][len]string) {
	if (j + consecutivos) <= len {
		if matriz[i][j] == matriz[i][j+1] && matriz[i][j] == matriz[i][j+2] && matriz[i][j] == matriz[i][j+3] {
			cont++
		}
	}
}

//Busqueda en direccion vertical hacia abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud vertical suficiente
func busq2(i int, j int, matriz [][len]string) {
	if (i + consecutivos) <= len {
		if matriz[i][j] == matriz[i+1][j] && matriz[i][j] == matriz[i+2][j] && matriz[i][j] == matriz[i+3][j] {
			cont++
		}
	}
}

//Busqueda en direccion derecha y abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq3(i int, j int, matriz [][len]string) {
	if (i+consecutivos) <= len && (j+consecutivos) <= len {
		if matriz[i][j] == matriz[i+1][j+1] && matriz[i][j] == matriz[i+2][j+2] && matriz[i][j] == matriz[i+3][j+3] {
			cont++
		}
	}
}

//Busqueda en direccion derecha y arriba
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq4(i int, j int, matriz [][len]string) {
	if (i-consecutivos+1) >= 0 && (j+consecutivos) <= len {
		if matriz[i][j] == matriz[i-1][j+1] && matriz[i][j] == matriz[i-2][j+2] && matriz[i][j] == matriz[i-3][j+3] {
			cont++
		}
	}
}

//Aplicar validaciones generales a la cadena recibida para verificar que cumple con las condiciones
func IsDnaValid(dna []string) (bool, string) {
//Validar que la cadena solo contenga los caracteres A,C,G,T	
	var regex_dna = regexp.MustCompile("^([ACGT]*)$")
	len = len(dna)

	for _, cadena := range dna {
		reg := regex_dna.FindStringSubmatch(strings.ToUpper(cadena))
		if reg == nil {
			return (false,"Caracter invlaido en la cadena")
		}
//Validar que la matriz sea simétrica
		if len(cadena) != len{
			return (false,"Matriz no es simetrica")
		}
	}
	return (true,"")
}
