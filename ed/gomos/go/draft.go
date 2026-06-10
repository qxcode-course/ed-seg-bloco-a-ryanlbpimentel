package main

import "fmt"

func main() {
	var qtdGomos int
	var direcao string

	type gomo struct {
		x int
		y int
	}

	cobra := make([]gomo, 0)

	fmt.Scan(&qtdGomos, &direcao)

	for i := 0; i < qtdGomos; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		cobra = append(cobra, gomo{x, y})
	}

	for i := qtdGomos - 1; i >= 0; i-- {
		if i > 0 {
			cobra[i].x = cobra[i-1].x
			cobra[i].y = cobra[i-1].y
		} else {
			switch direcao {
			case "L":
				cobra[i].x--
			case "R":
				cobra[i].x++
			case "U":
				cobra[i].y--
			case "D":
				cobra[i].y++
			}
		}

	}

	for i := 0; i < qtdGomos; i++ {
		fmt.Printf("%d %d\n", cobra[i].x, cobra[i].y)
	}

}
