package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (v *Vector) Reserve(newCapacity int) {
	//fmt.Println("teste1", v.capacity)
	if newCapacity <= v.capacity {
		return
	}

	if newCapacity == 0 {
		newCapacity = 1
	}

	v.data = append(v.data, make([]int, newCapacity-v.capacity)...)
	v.capacity = newCapacity

}
func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		v.Reserve(v.capacity * 2)
	}
	//fmt.Println("teste2", v.data)
	v.data[v.size] = value
	v.size++

	//fmt.Println("teste3", v.data)
}

func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}

func (v *Vector) Clear() {
	for i := 0; i < v.size; i++ {
		v.data[i] = 0
	}
	v.size = 0
}

func (v *Vector) PopBack() (int, error) {
	if v.size == 0 {
		return 0, fmt.Errorf("vector is empty")
	}
	v.size--
	return v.data[v.size], nil
}

func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}

	if v.size == v.capacity {
		v.Reserve(v.capacity * 2)
	}

	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[index] = value
	v.size++

	return nil
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.size--

	return nil
}

func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1
}

func (v *Vector) Slice(start int, end int) []int {
	if start < 0 {
		start = (start % v.size + v.size) % v.size
	}
	if end < 0 {
		end = (end % v.size + v.size) % v.size
	}
	
	if start > end {
		end += v.size
	}

	slice := NewVector(end - start)

	for i := start; i < end; i++ {
		value, _ := v.At(i % v.size)
		slice.PushBack(value)
	}

	return slice.data[:slice.size]
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Print("[")
			for i, value := range v.data[:v.size] {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(value)
			}
			fmt.Print("]\n")
		case "status":
			fmt.Printf("size:%d capacity:%d\n", v.size, v.capacity)
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)

			fmt.Print("[")
			for i, value := range slice {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(value)
			}
			fmt.Print("]\n")
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
