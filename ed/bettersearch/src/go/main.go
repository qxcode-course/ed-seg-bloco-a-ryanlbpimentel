package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := (low + high) / 2

		if slice[mid] == value {
			return true, mid
		} else if low == high {
			return false, mid + 1
		}

		if slice[mid] < value {
			low = mid + 1
		} else if slice[mid] > value {
			high = mid - 1
		}
	}
	return false, 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
