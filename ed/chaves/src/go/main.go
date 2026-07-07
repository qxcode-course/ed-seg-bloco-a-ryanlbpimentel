package main

import "fmt"

func main() {
	fila := NewQueue[int]()

	for i := 0; i < 16; i++ {
		fila.Enqueue('A' + i)
	}

	//fmt.Println(fila.String())

	//teste := fila.Dequeue()

	//fmt.Println(fila.String(), teste)

	for i := 1; i < 16; i++ {
		time1 := fila.Dequeue()
		time2 := fila.Dequeue()

		var esquerda, direita int

		fmt.Scan(&esquerda, &direita)

		if esquerda > direita {
			fila.Enqueue(time1)
		} else {
			fila.Enqueue(time2)
		}

		//fmt.Println(esquerda, direita, time1, time2)

		//fmt.Println(fila.String())
	}

	ganhador := fila.Dequeue()

	fmt.Printf("%c\n", ganhador)
}
