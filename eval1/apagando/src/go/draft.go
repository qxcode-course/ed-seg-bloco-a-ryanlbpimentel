package main

import "fmt"

func main() {
	var qtd, qtdRemover int
	fmt.Scan(&qtd)

	numeros := make([]int, qtd)

	for i := 0; i < qtd; i++ {
		fmt.Scan(&numeros[i])
	}

	fmt.Scan(&qtdRemover)

	remover := make([]int, qtdRemover)

	for i := 0; i < qtdRemover; i++ {
		fmt.Scan(&remover[i])
	}

	deveRemover := false

	final := make([]int, 0)

	for i := 0; i < qtd; i++ {
		for j := 0; j < qtdRemover; j++ {
			if numeros[i] == remover[j] {
				deveRemover = true
				break
			}
		}
		if !deveRemover {
			final = append(final, numeros[i])
		}
		deveRemover = false
	}

	for i := 0; i < len(final); i++ {
		fmt.Printf("%d ", final[i])
	}

	fmt.Println("")
}
