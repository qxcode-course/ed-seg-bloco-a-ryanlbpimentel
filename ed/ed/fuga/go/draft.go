package main

import "fmt"

func main() {
	var h, p, f, d int

	fmt.Scanf("%d %d %d %d", &h, &p, &f, &d)
	for {
		if d == 1 {
			f = f - 1
			if f < 0 {
				f = 15
			}
		} else if d == -1 {
			f = f + 1
			if f > 15 {
				f = 0
			}
		}
		//fmt.Println(h,p,f)

		if f == p {
			fmt.Println("S")
			break
		} else if f == h {
			fmt.Println("N")
			break
		}

	}
}
