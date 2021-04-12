package mutant

import (
	"regexp"
	"strings"
)

var fil, col, cont int

const consecutivos = 4
const numseq = 2

func Ismutant(dna []string) bool {
	i := 0
	j := 0
	cont = 0
	fil = len(dna)
	col = len(dna[0])

	matriz := make([][6]string, fil, col)

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

	for i = 0; i < fil; i++ {
		for j = 0; j < col; j++ {
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
func busq1(i int, j int, matriz [][6]string) {
	if (j + consecutivos) <= col {
		if matriz[i][j] == matriz[i][j+1] && matriz[i][j] == matriz[i][j+2] && matriz[i][j] == matriz[i][j+3] {
			cont++
		}
	}
}

//Busqueda en direccion vertical hacia abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud vertical suficiente
func busq2(i int, j int, matriz [][6]string) {
	if (i + consecutivos) <= fil {
		if matriz[i][j] == matriz[i+1][j] && matriz[i][j] == matriz[i+2][j] && matriz[i][j] == matriz[i+3][j] {
			cont++
		}
	}
}

//Busqueda en direccion derecha y abajo
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq3(i int, j int, matriz [][6]string) {
	if (i+consecutivos) <= fil && (j+consecutivos) <= col {
		if matriz[i][j] == matriz[i+1][j+1] && matriz[i][j] == matriz[i+2][j+2] && matriz[i][j] == matriz[i+3][j+3] {
			cont++
		}
	}
}

//Busqueda en direccion derecha y arriba
//Se realiza esta busqueda cuando en la posicion que estamos hay longitud suficiente
func busq4(i int, j int, matriz [][6]string) {
	if (i-consecutivos+1) >= 0 && (j+consecutivos) <= col {
		if matriz[i][j] == matriz[i-1][j+1] && matriz[i][j] == matriz[i-2][j+2] && matriz[i][j] == matriz[i-3][j+3] {
			cont++
		}
	}
}

// Validar mediante regex si cada una de las cadenas ingresadas son validas
func IsDnaValid(dna []string) bool {
	var regex_dna = regexp.MustCompile("^([ACGT]*)$")
	for _, cadena := range dna {
		reg := regex_dna.FindStringSubmatch(strings.ToUpper(cadena))
		if reg == nil {
			return false
		}
	}
	return true
}
