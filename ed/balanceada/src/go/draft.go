package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stack := NewStack[string]()

	buffer := bufio.NewReader(os.Stdin)
	text, _ := buffer.ReadString('\n')
	trimmed := strings.TrimSpace(text)

	char := strings.Split(trimmed, "")

	for _, c := range char {
		if c == "(" || c == "[" || c == "{" {
			stack.Push(c)
		} else {
			if stack.IsEmpty() {
				fmt.Println("nao balanceado")
				return
			}
			top := stack.Pop()

			if (c == ")" && top != "(") || (c == "]" && top != "[") || (c == "}" && top != "{") {
				fmt.Println("nao balanceado")
				return
			}
		}
	}

	if stack.IsEmpty() {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}
}
