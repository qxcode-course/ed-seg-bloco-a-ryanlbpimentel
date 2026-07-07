package main

import "fmt"

func backtrackingSoma(arrayValores []int, valor int) bool {

	if valor == 0 {
		return true
	}

	if len(arrayValores) == 0 {
		return false
	}

	inclui := backtrackingSoma(arrayValores[1:], valor-arrayValores[0])

	if inclui {
		return true
	}

	naoInclui := backtrackingSoma(arrayValores[1:], valor)

	if naoInclui {
		return true
	}

	return false
}

func main() {
	var qtd, valor int

	fmt.Scan(&qtd)
	fmt.Scan(&valor)

	var valores []int

	for i := 0; i < qtd; i++ {
		var v int
		fmt.Scan(&v)
		valores = append(valores, v)
	}

	resultado := backtrackingSoma(valores, valor)

	fmt.Println(resultado)
}
