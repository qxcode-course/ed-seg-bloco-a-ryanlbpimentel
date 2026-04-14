package main

import "fmt"

func main() {
	var qtd, qtdBaruel int

	original := make([]int, 0)

	fmt.Scan(&qtd, &qtdBaruel)

	for i := 0; i < qtdBaruel; i++ {
		var valor int
		fmt.Scan(&valor)
		original = append(original, valor)
	}

	repetidas := make(map[int]int, 0)

	for _, valor := range original {
		repetidas[valor]++
	}

	var primeiro = false

	qtdRepetidas := 0

	for valor, count := range repetidas {
		if count > 1 {
			for i := 0; i < count-1; i++ {
				if primeiro == false {
					primeiro = true
					fmt.Printf("%d", valor)
				} else {
					fmt.Printf(" %d", valor)
				}
				qtdRepetidas++
			}
		}
	}

	if qtdRepetidas == 0 {
		fmt.Printf("N\n")
	} else {
		fmt.Println("")
	}

	qtdFaltam := 0

	for i := 1; i <= qtd; i++ {
		if repetidas[i] == 0 {
			if i > 1 && qtdFaltam > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", i)
			qtdFaltam++
		}
	}

	if qtdFaltam == 0 {
		fmt.Printf("N\n")
	} else {
		fmt.Println("")
	}
}
