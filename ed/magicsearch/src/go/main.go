package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := (low + high) / 2

		if slice[mid] == value {
			for i := mid; i < len(slice); i++ {
				if slice[i] == value {
					mid = i
				} else {
					break
				}
			}
			return mid
		} else if low == high {
			return mid + 1
		}

		if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
