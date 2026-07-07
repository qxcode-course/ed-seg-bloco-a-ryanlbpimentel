package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	type Pos struct {
		l int
		c int
	}

	stack := NewStack[Pos]()
	pos := Pos{l: l, c: c}

	stack.Push(pos)

	if grid[l][c] == '#' {
		grid[l][c] = 'o'
	} else {
		return
	}

	stack.Pop()

	if l+1 < len(grid) && grid[l+1][c] == '#' {
		stack.Push(Pos{l: l + 1, c: c})
		burnTrees(grid, l+1, c)
	}

	if l-1 >= 0 && grid[l-1][c] == '#' {
		stack.Push(Pos{l: l - 1, c: c})
		burnTrees(grid, l-1, c)
	}
	if c+1 < len(grid[0]) && grid[l][c+1] == '#' {
		stack.Push(Pos{l: l, c: c + 1})
		burnTrees(grid, l, c+1)
	}
	if c-1 >= 0 && grid[l][c-1] == '#' {
		stack.Push(Pos{l: l, c: c - 1})
		burnTrees(grid, l, c-1)
	}

	if stack.IsEmpty() {
		return
	}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
