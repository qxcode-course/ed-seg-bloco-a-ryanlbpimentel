package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var vetFinal = []int{}

	for _, val := range vet {
		if val > 0 {
			vetFinal = append(vetFinal, val)
		}
	}

	return vetFinal
}

func getCalmWomen(vet []int) []int {
	var vetFinal = []int{}

	for _, val := range vet {
		if val > -10 && val <= 0 {
			vetFinal = append(vetFinal, val)
		}
	}
	return vetFinal
}

func sortVet(vet []int) []int {
	slices.Sort(vet)
	return vet
}

func sortStress(vet []int) []int {

	sort.Slice(vet, func(i, j int) bool {
		return math.Abs(float64(vet[i])) < math.Abs(float64(vet[j]))
	})

	return vet
}

func reverse(vet []int) []int {
	vetInvertido := []int{}

	for i := len(vet) - 1; i >= 0; i-- {
		vetInvertido = append(vetInvertido, vet[i])
	}

	return vetInvertido
}

func unique(vet []int) []int {
	vetFinal := []int{}

	var achou = false

	for _, val := range vet {
		for _, val2 := range vetFinal {
			if val == val2 {
				achou = true
				break
			}
		}
		if !achou {
			vetFinal = append(vetFinal, val)
		}
		achou = false
	}

	return vetFinal
}

func repeated(vet []int) []int {
	vetCopia := vet
	vetFinal := []int{}

	var contador = 0

	for _, v1 := range vet {
		if slices.Contains(vetFinal, v1) {
			continue
		}

		for _, v2 := range vetCopia {
			if v1 == v2 {
				contador++
			}
		}

		if contador > 1 {
			for i := 0; i < contador-1; i++ {
				vetFinal = append(vetFinal, v1)
			}
		}
		contador = 0
	}

	slices.Sort(vetFinal)

	return vetFinal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
