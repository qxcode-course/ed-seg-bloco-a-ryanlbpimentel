package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	for i := 0; i < ms.capacity; i++ {
		ms.data = append(ms.data, 0)
	}
	ms.capacity *= 2
}

func (ms *MultiSet) insert(value int, index int) error {
	//fmt.Println(ms.data, index, value)
	for i := ms.size; i > index; i-- {
		//fmt.Println(v.data, i, v.data[i-1])
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++

	//fmt.Println(ms.data, index, value)
	return nil
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}

	for i := 0; i < ms.capacity; i++ {
		if ms.data[i] == 0 {
			ms.insert(value, i)
			break
		}

		if ms.data[i] >= value {
			ms.insert(value, i)
			break
		}

	}
}

func (ms *MultiSet) Contains(value int) bool {
	for i := 0; i < ms.capacity; i++ {
		if ms.data[i] == value {
			return true
		}

	}
	return false
}

func (ms *MultiSet) erase(index int) {
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.data[ms.size-1] = 0
	ms.size--
}

func (ms *MultiSet) Erase(value int) bool {
	for i := 0; i <= ms.size; i++ {
		if ms.data[i] == value {
			//fmt.Println("achou")
			ms.erase(i)
			return true
		} else {
			//fmt.Println("nao achou")
			continue
		}
	}
	return false
}

func (ms *MultiSet) Count(value int) int {
	soma := 0
	for i := 0; i < ms.capacity; i++ {
		if ms.data[i] == value {
			soma++
		}

	}
	return soma
}

func (ms *MultiSet) Unique() int {
	soma := 0
	valor := 0

	for i := 0; i < ms.capacity; i++ {
		if ms.data[i] == 0 {
			break
		}

		if i == 0 {
			valor = ms.data[i]
			soma++
			continue
		}

		if ms.data[i] != valor {
			//fmt.Println(ms.data[i], valor)
			valor = ms.data[i]
			soma++
		}

	}

	return soma
}

func (ms *MultiSet) Clear() {
	for i := 0; i < ms.capacity; i++ {
		ms.data[i] = 0
	}

	ms.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}

		case "show":
			fmt.Print("[")

			for i, value := range ms.data[:ms.size] {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(value)
			}

			fmt.Println("]")
		case "erase":
			value, _ := strconv.Atoi(args[1])

			if !ms.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])

			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])

			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
