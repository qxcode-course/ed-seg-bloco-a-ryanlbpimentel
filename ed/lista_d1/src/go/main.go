package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int   // Valor é público
	next  *Node // o próximo nó da lista
	prev  *Node // o nó anterior
}

type LList struct {
	root *Node // Nó sentinela que marca o começo e o fim da lista
}

func NewLList() *LList {
	ll := &LList{}
	ll.root = &Node{}
	ll.root.next = ll.root
	ll.root.prev = ll.root
	return ll
}

func (ll *LList) String() string {
	fmt.Print("[")
	node := ll.root.next

	for node != ll.root {
		fmt.Print(node.Value)
		node = node.next

		if node != ll.root {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
	return ""
}

func (ll *LList) Size() int {
	c := 0

	node := ll.root.next

	for node != ll.root {
		c++
		node = node.next
	}

	return c
}

func (ll *LList) PushFront(value int) {
	node := &Node{Value: value}
	node.next = ll.root.next
	node.prev = ll.root
	ll.root.next.prev = node
	ll.root.next = node
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *LList) PushBack(value int) {
	node := &Node{Value: value}
	node.next = ll.root
	node.prev = ll.root.prev
	ll.root.prev.next = node
	ll.root.prev = node
}

func (ll *LList) PopFront() {
	ll.root.next = ll.root.next.next
	ll.root.next.prev = ll.root
}

func (ll *LList) PopBack() {
	ll.root.prev = ll.root.prev.prev
	ll.root.prev.next = ll.root
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
