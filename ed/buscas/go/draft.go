package main

import "fmt"

func main() {
	var nString int
	fmt.Scan(&nString)

	original := make([]string, nString)

	for i := 0; i < nString; i++ {
		fmt.Scan(&original[i])
	}

	var nStringConsulta int
	fmt.Scan(&nStringConsulta)
	consulta := make(map[string]int, nStringConsulta)
	palavras := make([]string, nStringConsulta)

	for i := 0; i < nStringConsulta; i++ {
		var palavra string
		fmt.Scan(&palavra)
		consulta[palavra] = 0
		palavras[i] = palavra
	}

	for i := 0; i < nString; i++ {
		consulta[original[i]]++
	}

	for i := 0; i < nStringConsulta; i++ {
		fmt.Print(consulta[palavras[i]])
		if i < nStringConsulta-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}
