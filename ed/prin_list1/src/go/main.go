package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var sb strings.Builder

	sb.WriteString("[ ")
	for n := l.Front(); n != l.End(); n = n.Next() {
		if n == sword {
			sb.WriteString(fmt.Sprintf("%v>", n.Value))
		} else {
			sb.WriteString(fmt.Sprintf("%v", n.Value))
		}

		if n.Next() != l.End() {
			sb.WriteString(" ")
		}
	}
	sb.WriteString(" ]")
	return sb.String()
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.next == l.End() {
		return it.next.next
	}
	return it.next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
