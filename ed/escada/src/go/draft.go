package main

import "fmt"

func escadaria(n int, a int) int {
	if a == n {
		return 1
	}

	if a > n {
		return 0
	}

	return escadaria(n, a+1) + escadaria(n, a+3)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(escadaria(n, 0))
}
