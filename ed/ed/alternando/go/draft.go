package main

import "fmt"

func printar(fila []bool, e int) {
	fmt.Print("[ ")

	for i := range fila {

		if fila[i] == true {
			if (i) == e {
				fmt.Print(i+1, "> ")
			} else {
				fmt.Print(i+1, " ")
			}
		}
	}

	fmt.Println("]")
}

func contar(fila []bool) int {
	contador := 0

	for i := range fila {

		if fila[i] == true {
			contador++
		}
	}

	return contador
}

func matar(fila []bool, n int, e int) {
	printar(fila, e)

	for {
		if fila[(e+1)%n] == true {
			fila[(e+1)%n] = false

			for {
				if fila[(e+2)%n] == true {
					e = (e + 2) % n
					break
				}
				e = (e + 1) % n
			}

			break
		} else {
			e = (e + 1) % n
			continue
		}
	}

	if contar(fila) == 1 {
		printar(fila, e)
		return
	}

	matar(fila, n, e)
}

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	fila := make([]bool, n)

	for i := 0; i < n; i++ {
		fila[i] = true
	}

	e = e - 1

	matar(fila, n, e)

}
