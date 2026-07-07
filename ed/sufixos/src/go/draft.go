package main

import "fmt"

func processa(s string, pos int) {
	if pos < 0 {
		return
	}
	fmt.Println(s[pos:])
	processa(s, pos-1)
}

func main() {
	var s string
	fmt.Scan(&s)

	processa(s, len(s)-1)
}
