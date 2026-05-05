package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	novaLinha := NewList[rune]()

	for no := e.cursor; no != e.line.Value.End(); no = e.cursor {
		novaLinha.PushBack(no.Value)
		e.cursor = e.line.Value.Erase(no)
	}

	e.line = e.lines.Insert(e.line.Next(), novaLinha)
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
	fimDaLinha := e.line.Value.End()

	if e.cursor == fimDaLinha {
		if e.line.Next() != e.lines.End() {
			e.line = e.line.Next()
			e.cursor = e.line.Value.Front()
		}
		return
	}

	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyUp() {
	if e.line == e.lines.Front() {
		return
	}

	coluna := e.line.Value.IndexOf(e.cursor)
	if coluna < 0 {
		coluna = 0
	}

	e.line = e.line.Prev()
	linhaDestino := e.line.Value

	if coluna <= 0 {
		e.cursor = linhaDestino.Front()
		return
	}
	if coluna >= linhaDestino.Size() {
		e.cursor = linhaDestino.End()
		return
	}

	no := linhaDestino.Front()
	for i := 0; i < coluna; i++ {
		no = no.Next()
	}
	e.cursor = no
}

func (e *Editor) KeyDown() {
	if e.line.Next() == e.lines.End() {
		return
	}

	coluna := e.line.Value.IndexOf(e.cursor)
	if coluna < 0 {
		coluna = 0
	}

	e.line = e.line.Next()
	linhaDestino := e.line.Value

	if coluna <= 0 {
		e.cursor = linhaDestino.Front()
		return
	}
	if coluna >= linhaDestino.Size() {
		e.cursor = linhaDestino.End()
		return
	}

	no := linhaDestino.Front()
	for i := 0; i < coluna; i++ {
		no = no.Next()
	}
	e.cursor = no
}

func (e *Editor) KeyBackspace() {
	if e.cursor != e.line.Value.Front() {
		e.line.Value.Erase(e.cursor.Prev())
		return
	}

	if e.line == e.lines.Front() {
		return
	}

	noLinhaAnterior := e.line.Prev()
	linhaAnterior := noLinhaAnterior.Value
	tamanhoAntes := linhaAnterior.Size()

	for noChar := e.line.Value.Front(); noChar != e.line.Value.End(); noChar = noChar.Next() {
		linhaAnterior.PushBack(noChar.Value)
	}

	noParaRemover := e.line
	e.line = noLinhaAnterior
	e.lines.Erase(noParaRemover)

	if tamanhoAntes <= 0 {
		e.cursor = linhaAnterior.Front()
		return
	}
	if tamanhoAntes >= linhaAnterior.Size() {
		e.cursor = linhaAnterior.End()
		return
	}

	no := linhaAnterior.Front()
	for i := 0; i < tamanhoAntes; i++ {
		no = no.Next()
	}
	e.cursor = no
}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.line.Value.Erase(e.cursor)
		return
	}

	if e.line.Next() == e.lines.End() {
		return
	}

	tamanhoAntes := e.line.Value.Size()
	noProximaLinha := e.line.Next()

	for noChar := noProximaLinha.Value.Front(); noChar != noProximaLinha.Value.End(); noChar = noChar.Next() {
		e.line.Value.PushBack(noChar.Value)
	}

	e.lines.Erase(noProximaLinha)

	linha := e.line.Value
	if tamanhoAntes <= 0 {
		e.cursor = linha.Front()
		return
	}
	if tamanhoAntes >= linha.Size() {
		e.cursor = linha.End()
		return
	}

	no := linha.Front()
	for i := 0; i < tamanhoAntes; i++ {
		no = no.Next()
	}
	e.cursor = no
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair	
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}