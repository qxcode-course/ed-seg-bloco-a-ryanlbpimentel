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
	root  *Node // aponta para o nó sentinela da lista da qual ele faz parte
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type LList struct {
	root *Node // Nó sentinela que marca o começo e o fim da lista
	size int   // tamanho da lista
}

func NewLList() *LList {
	ll := &LList{}
	ll.root = &Node{}
	ll.root.root = ll.root
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
	node := &Node{Value: value, root: ll.root}
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
	node := &Node{Value: value, root: ll.root}
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

func (ll *LList) Front() *Node {
	for node := ll.root.next; node != ll.root; node = node.next {
		return node
	}
	return nil
}

func (ll *LList) Back() *Node {
	for node := ll.root.prev; node != ll.root; node = node.prev {
		return node
	}
	return nil
}

func (ll *LList) Search(value int) *Node {
	for node := ll.Front(); node != nil; node = node.Next() {
		if node.Value == value {
			return node
		}
	}
	return nil
}

func (ll *LList) Insert(n *Node, value int) {
	count := 1
	for node := ll.Front(); node != nil; node = node.Next() {
		if node == n {
			newNode := &Node{Value: value, root: ll.root}
			newNode.next = node
			newNode.prev = node.prev
			node.prev.next = newNode
			node.prev = newNode
			return
		}
		count++
	}
}

func (ll *LList) Remove(n *Node) {
	for node := ll.Front(); node != nil; node = node.Next() {
		if node == n {
			node.prev.next = node.next
			node.next.prev = node.prev
			return
		}
	}
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
