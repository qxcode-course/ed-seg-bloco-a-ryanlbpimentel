package main

import "fmt"

func main() {
	var qtdFigurinhas, qtdBaruel int
	fmt.Scan(&qtdFigurinhas, &qtdBaruel)

	var figurinhas = make(map[int]int)

	for i := 0; i < qtdBaruel; i++ {
		var figurinha int
		fmt.Scan(&figurinha)
		figurinhas[figurinha]++
	}

	var repetiram = false

	// As figurinhas que repetiram
	for i := 0; i <= qtdFigurinhas; i++ {
		//fmt.Println(figurinhas)
		if figurinhas[i] > 1 {
			for j := 1; j < figurinhas[i]; j++ {
				if j == 1 && repetiram == true {
					fmt.Print(" ")
				}

				repetiram = true
				fmt.Print(i)

				if i < qtdBaruel && j < figurinhas[i]-1 {
					fmt.Print(" ")
				}

				if j == figurinhas[i]-1 {
					break
				}
			}

		}
	}

	if repetiram == false {
		fmt.Println("N")
	} else {
		fmt.Println()
	}

	var faltam = false

	for i := 1; i <= qtdFigurinhas; i++ {
		if figurinhas[i] == 0 {
			if i <= qtdFigurinhas && faltam == true {
				fmt.Print(" ")
			}

			faltam = true
			fmt.Print(i)
		}
	}

	if faltam == false {
		fmt.Println("N")
	} else {
		fmt.Println()
	}

}
