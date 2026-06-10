package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var arca = make(map[int]int)
	var casais = 0

	for i := 0; i < n; i++ {
		var animal int
		fmt.Scan(&animal)

		if animal < 0 {
			arca[-animal]--
			if arca[-animal] == 0 {
				casais++
			}
		} else {
			arca[animal]++
			if arca[animal] == 0 {
				casais++
			}
		}
	}

	fmt.Println(casais)
}
