package main

import (
	"fmt"
)

func main() {
	var qtd int
	fmt.Scan(&qtd)

	casais := make(map[int]int)
	qtdCasais := 0

	for i := 0; i < qtd; i++ {
		var animal int
		fmt.Scan(&animal)

		var numNeutro int

		if animal < 0 {
			numNeutro = -animal
		} else {
			numNeutro = animal
		}

		if casais[numNeutro] >= 1 && animal <= -1 {
			qtdCasais++
			casais[numNeutro] -= 1
		} else if casais[numNeutro] <= -1 && animal >= 1 {
			qtdCasais++
			casais[numNeutro] += 1
		} else {
			if animal < 0 {
				casais[numNeutro] -= 1
			} else {
				casais[numNeutro] += 1
			}
		}
	}

	fmt.Println(qtdCasais)

}
