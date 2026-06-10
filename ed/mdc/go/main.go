package main

import (
	"fmt"
)

func mdc(a, b int) string {
	if b == 0 {
		return fmt.Sprintf("%d", a)
	}
	return mdc(b, a%b)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
