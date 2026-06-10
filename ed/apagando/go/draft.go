package main

import "fmt"

func main() {
	var n,nSairam int
	fmt.Scan(&n)

	var filaOriginal = make([]int, n)
	var filaFinal = []int{}

	for i := 0; i < n; i++ {
		fmt.Scan(&filaOriginal[i])
	}

	fmt.Scan(&nSairam)

	var filaSairam = make([]int, nSairam)

	for i := 0; i < nSairam; i++ {
		fmt.Scan(&filaSairam[i])
	}

	for i := range filaOriginal{
		achou := false
		
		for j := range filaSairam{
			if filaOriginal[i] == filaSairam[j]{
				achou = true
			}
		}
		
		if achou == false{
			filaFinal = append(filaFinal, filaOriginal[i])
		}
	}

	for i := range filaFinal{
		fmt.Print(filaFinal[i], " ")
	}

	fmt.Println()
}
