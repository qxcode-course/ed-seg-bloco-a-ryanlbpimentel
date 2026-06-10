package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func count(vet []int, i int, str string, reverse bool) string {
	str += strconv.Itoa(vet[i])

	if !reverse {
		if i == len(vet)-1 {
			return str
		} else {
			str += ", "
		}

		str = count(vet, (i + 1), str, false)
	} else {
		if i == 0 {
			return str
		} else {
			str += ", "
		}

		str = count(vet, (i - 1), str, true)
	}

	return str
}

func tostr(vet []int) string {
	str := "["

	if len(vet) > 0 {
		str += count(vet, 0, "", false)
	}

	str += "]"

	return str
}

func tostrrev(vet []int) string {
	str := "["

	if len(vet) > 0 {
		str += count(vet, len(vet)-1, "", true)
	}

	str += "]"

	return str
}

func trocar(vet []int, i, j int) []int {
	vet[i], vet[j] = vet[j], vet[i]

	if i < j-1 {
		return trocar(vet, i+1, j-1)
	}

	return vet
}

func reverse(vet []int) []int {
	if len(vet) == 0 {
		return vet
	}

	vet = trocar(vet, 0, len(vet)-1)
	return vet
}

func somar(vet []int, i int, soma int) int {
	if i == len(vet)-1 {
		return soma + vet[i]
	}

	return somar(vet, i+1, soma+vet[i])
}

func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return somar(vet, 0, 0)
}

func multiplicar (vet []int, i int, mult int) int {
	if i == len(vet)-1 {
		return mult * vet[i]
	}

	return multiplicar(vet, i+1, mult*vet[i])
}

func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return multiplicar(vet, 0, 1)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice

func minIndex(vet []int) (int, int) {
	var rec func(v []int) (int, int)
	
	rec = func(v []int) (int, int){
		if len(v) == 1 {
			return 0, v[0]
		}
		
		indiceMenor, menor := rec(v[1:])
		if v[0] < menor {
			return 0, v[0]
		}
			
		return indiceMenor + 1, menor
	}
	return rec(vet)
}
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	indice, _ := minIndex(vet)
	return indice
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			vet = reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
