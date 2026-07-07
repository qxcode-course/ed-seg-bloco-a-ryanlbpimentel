package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func verificarAtras(char []string, index int, L int, valor string) bool {
	var contador = 0
	for i := index - 1; i >= 0; i-- {
		if index < 0 {
			break
		}

		contador++
		if char[i] == valor {
			return true
		}

		if contador == L {
			break
		}
	}

	return false
}

func verificarFrente(char []string, index int, L int, valor string) bool {
	var contador = 0
	for i := index + 1; i < len(char); i++ {
		if index >= len(char)-1 {
			break
		}

		contador++
		if char[i] == valor {
			return true
		}

		if contador == L {
			break
		}
	}
	return false
}

func backtracking(char []string, current string, index int, L int) string {
	if index == len(char) {
		return current
	}

	if char[index] != "." {
		current += char[index]
		return backtracking(char, current, index+1, L)
	} else if char[index] == "." {
		for i := 0; i <= L; i++ {
			vlr := strconv.Itoa(i)
			if !verificarAtras(char, index, L, vlr) && !verificarFrente(char, index, L, vlr) {
				char[index] = vlr
				resultado := backtracking(char, current+vlr, index+1, L)

				if resultado != "" {
					return resultado
				}
				char[index] = "."
			}
		}
	}

	return ""
}

func main() {
	buffer := bufio.NewReader(os.Stdin)
	string1, _ := buffer.ReadString('\n')
	trimmed := strings.TrimSpace(string1)

	char := strings.Split(trimmed, "")

	string2, _ := buffer.ReadString('\n')
	L, _ := strconv.Atoi(strings.TrimSpace(string2))

	fmt.Println(backtracking(char, "", 0, L))
}
