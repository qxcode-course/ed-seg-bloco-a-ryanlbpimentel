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

func enesimo_primo(n int, contador int, primo int) int {
	if eh_primo(primo, 2) {
		contador += 1
		if contador == n {
			return primo
		}
	}

	primo += 1
	return enesimo_primo(n, contador, primo)
}

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Println(enesimo_primo(n, 0, 1))
}
