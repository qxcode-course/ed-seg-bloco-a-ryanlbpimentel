package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Set) reserve(newCapacity int) {
	for i := 0; i < newCapacity; i++ {
		v.data = append(v.data, 0)
	}
	v.capacity = newCapacity
}

func (v *Set) insert(value int, index int) error {
	//fmt.Println(v.data, index, value)
	for i := v.size; i > index; i-- {
		//fmt.Println(v.data, i, v.data[i-1])
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.size++
	//fmt.Println(v.data, index, value)
	return nil
}

func (v *Set) Insert(value int) {
	if v.size == v.capacity {
		v.reserve(v.capacity * 2)
	}
	//fmt.Println(v.data)
	for i := 0; i < v.capacity; i++ {
		if v.data[i] < value && v.data[i] != 0 {
			continue
		} else if v.data[i] == value {
			break
		} else {
			v.insert(value, i)
			break
		}
	}

}

func (v *Set) binarySearch(value int) int {
	esquerda, direita := 0, v.size

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if v.data[meio] == value {
			return meio
		}

		if v.data[meio] < value {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}

	return -1
}

func (v *Set) Contains(value int) bool {
	if v.binarySearch(value) != -1 {
		return true
	}
	return false
}

func (v *Set) erase(index int) {
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.data[v.size-1] = 0
	v.size--
}

func (v *Set) Erase(value int) bool {
	for i := 0; i <= v.size; i++ {
		if v.data[i] == value {
			//fmt.Println("achou")
			v.erase(i)
			return true
		} else {
			//fmt.Println("nao achou")
			continue
		}
	}
	return false
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Print("[")

			for i, value := range v.data[:v.size] {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(value)
			}

			fmt.Println("]")
		case "erase":
			value, _ := strconv.Atoi(parts[1])

			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])

			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
