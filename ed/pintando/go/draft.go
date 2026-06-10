 package main

import (
	"fmt"
	"math"
)

func calcular(a, b, c float64) float64 {
	p := (a + b + c) / 2
	heron := (p * (p - a) * (p - b) * (p - c))
	area := math.Sqrt(math.Abs(heron))

	return area
}

//6 * (6-4) * (6-5) * (6-6) =

func main() {
	var a, b, c float64

	fmt.Scanf("%f\n%f\n%f", &a, &b, &c)
	area := calcular(a, b, c)

	fmt.Printf("%.2f\n", area)
}
