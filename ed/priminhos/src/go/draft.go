package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x <= 1 {
		return false
	}

	if (x%div == 0) && (x != div) {
		return false
	} else if div >= x {
		return true
	}

	return eh_primo(x, div+1)
}

func enesimo_primo(n int, contador int, primo int, vet []int) []int {
	if eh_primo(primo, 2) {
		contador += 1
		vet = append(vet, primo)
		if contador == n {
			return vet
		}
	}

	primo += 1
	return enesimo_primo(n, contador, primo, vet)
}

func main() {
	var n int
	var vet []int
	fmt.Scan(&n)

	vet = enesimo_primo(n, 0, 1, vet)

	fmt.Print("[")

	for i := 0; i <= len(vet)-1; i++ {
		if i != len(vet)-1 {
			fmt.Printf("%d, ", vet[i])
		} else {
			fmt.Print(vet[i])
		}
	}
	fmt.Println("]")
}
