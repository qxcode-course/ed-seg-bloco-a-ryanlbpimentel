package main

import (
	"fmt"
	"math"
)

type Jogador struct {
	PedraA          int
	PedraB          int
	Pontuacao       int
	Desclassificado bool
}

func main() {
	var n int
	fmt.Scan(&n)

	jogadores := make([]Jogador, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&jogadores[i].PedraA, &jogadores[i].PedraB)
		jogadores[i].Desclassificado = false

		if jogadores[i].PedraA < 10 || jogadores[i].PedraB < 10 {
			jogadores[i].Desclassificado = true
		}

		jogadores[i].Pontuacao = int(math.Abs(float64(jogadores[i].PedraA - jogadores[i].PedraB)))
		//fmt.Print(jogadores[i])
	}

	var ganhador = 100
	var ganhadorIndex = 100

	for i := 0; i < n; i++ {
		if !jogadores[i].Desclassificado {
			if jogadores[i].Pontuacao < ganhador {
				ganhador = jogadores[i].Pontuacao
				ganhadorIndex = i
			}
		}
	}

	if ganhadorIndex == 100 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(ganhadorIndex)
	}

}
