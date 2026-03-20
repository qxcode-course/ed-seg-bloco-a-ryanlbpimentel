package main

import "fmt"

func classify(i int) string {
	switch {
	case i < 12:
		return " eh crianca"
	case i < 18:
		return " eh jovem"
	case i < 65:
		return " eh adulto"
	case i < 1000:
		return " eh idoso"
	default:
		return " eh mumia"
	}
}

func main() {
	var n string
	var i int

	fmt.Scanf("%s\n%d", &n, &i)

	fmt.Println(n + classify(i))
}
