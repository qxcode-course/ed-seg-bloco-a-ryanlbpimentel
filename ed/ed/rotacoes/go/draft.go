package main

import "fmt"

func main() {
	var n, rotacoes int
	fmt.Scan(&n, &rotacoes)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < rotacoes; i++ {
		var ultimo = arr[len(arr)-1]

		for j := 0; j < len(arr)-1; j++ {
			arr[len(arr)-1-j] = arr[len(arr)-2-j]
		}
		arr[0] = ultimo
	}

	fmt.Print("[ ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("]\n")

}
