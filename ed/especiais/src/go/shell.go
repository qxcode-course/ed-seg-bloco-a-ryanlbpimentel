package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {

	m := make(map[int]int)

	for _, v := range vet {
		m[int(math.Abs(float64(v)))] += 1
	}

	var Pares []Pair

	for key, value := range m {
		Pares = append(Pares, Pair{One: key, Two: value})
	}

	sort.Slice(Pares, func(i, j int) bool {
		return Pares[i].One < Pares[j].One
	})

	return Pares
}

func teams(vet []int) []Pair {
	var Pares []Pair
	contador := 0

	for i, v := range vet {
		if i == 0 {
			if len(vet) == 1 {
				Pares = append(Pares, Pair{One: v, Two: 1})
			}
			contador++
		} else {
			if i == len(vet)-1 {
				if v == vet[i-1] {
					contador++
					Pares = append(Pares, Pair{One: v, Two: contador})
				} else {
					Pares = append(Pares, Pair{One: vet[i-1], Two: contador})
					contador = 0
					Pares = append(Pares, Pair{One: v, Two: 1})
				}

			} else {
				if v == vet[i-1] {
					contador++
				} else {
					Pares = append(Pares, Pair{One: vet[i-1], Two: contador})
					contador = 1
				}
			}
		}

	}

	/*sort.Slice(Pares, func(i, j int) bool {
		return Pares[i].One < Pares[j].One
	})*/

	return Pares
}

func mnext(vet []int) []int {
	mnext := []int{}
	for i, v := range vet {

		if i == 0 {
			mnext = append(mnext, 0)
		} else {
			anterior := vet[i-1]

			if anterior > 0 && v < 0 {
				mnext[i-1] = 1
				mnext = append(mnext, 0)
			} else if anterior < 0 && v > 0 {
				mnext = append(mnext, 1)
			} else {
				mnext = append(mnext, 0)
			}
		}
	}
	return mnext
}

func alone(vet []int) []int {
	alone := []int{}
	for i, v := range vet {

		if v >= 0 {
			alone = append(alone, 1)
		} else {
			alone = append(alone, 0)
		}

		if i > 0 {
			vAnt := vet[i-1]

			if vAnt > 0 && v < 0 {
				alone[i-1] = 0
				alone[i] = 0
			} else if vAnt < 0 && v > 0 {
				alone[i] = 0
			}
		}
	}
	return alone
}

func couple(vet []int) int {
	count := 0
	for i, v := range vet {
		if i > 0 {
			vAnt := vet[i-1]

			if vAnt < 0 && v > 0 || vAnt > 0 && v < 0 {
				count++
				vet[i] = 0
			}
		}
	}
	return count
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	count := -1
	tamanho := len(seq)
	tamanhovet := len(vet) - 1
	for i, v := range vet {
		if (i + tamanho - 1) > tamanhovet {
			break
		}

		if v == seq[0] && vet[i+tamanho-1] == seq[tamanho-1] {
			count = i
			break
		}
	}

	return count
}

func erase(vet []int, posList []int) []int {
	resto := []int{}
	ihTem := false
	for i, v := range vet {
		for _, v := range posList {
			if i == v {
				ihTem = true
			}
		}

		if ihTem == false {
			resto = append(resto, v)
		}
		ihTem = false
	}

	return resto
}

func clear(vet []int, value int) []int {
	resto := []int{}
	for _, v := range vet {
		if v != value {
			resto = append(resto, v)
		}
	}
	return resto
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
