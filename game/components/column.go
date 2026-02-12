package components

import (
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/hajimehoshi/ebiten/v2"
)

// Column organiza widgets em uma coluna, com espaçamento e largura opcional fixa
// <USAR A LARGURA DO PAI PARA ALINHAMENTO NO EIXO SECUNDARIO (X)>
type Column struct {
	Pos, currentPos basic.Point //posição inicial

	Spacing float32 //espaçamento vertical entre elementos

	cursorY float32 //posição vertical do próximo elemento

	MainAlign basic.Align // alinhamento dos elementos no eixo principal (center, start, end)

	CrossAlign basic.Align //eixo cruzado

	Children []Widget //elementos da column

	size basic.Size //caso necessario
}

// NewColumn cria uma coluna e já calcula a posição de todos os widgets
// alinhando no eixo principal (vertical) e no eixo cruzado (horizontal)
func NewColumn(
	pos basic.Point,
	spacing float32,
	parentSize basic.Size,
	mainAlign basic.Align,
	crossAlign basic.Align,
	children []Widget,
) *Column {

	c := &Column{
		Pos:        pos,
		Spacing:    spacing,
		Children:   children,
		MainAlign:  mainAlign,
		CrossAlign: crossAlign,
	}

	// posicionamento inicial (Start/Start)
	c.init()

	// se ambos forem Start, não faz nada
	if mainAlign == basic.Start && crossAlign == basic.Start {
		return c
	}

	if mainAlign != basic.Start {
		c.alignMain(parentSize)
	}
	if crossAlign != basic.Start {
		c.alignCross(parentSize)
	}

	for _, w := range c.Children {
		p := w.GetPos()
		p.X += c.Pos.X
		p.Y += c.Pos.Y
		w.SetPos(p)
	}

	return c
}

// Update chama Update de todos os filhos
func (c *Column) Update(offset basic.Point) {
	c.currentPos = c.Pos.Add(offset)

	for _, w := range c.Children {
		w.Update(offset)
	}
}

func (c *Column) Draw(screen *ebiten.Image) {
	for _, w := range c.Children {
		w.Draw(screen)
	}
}

func (c *Column) alignMain(parentSize basic.Size) {
	content := c.GetSize()

	var offsetY float32
	switch c.MainAlign {
	case basic.Start:
		return
	case basic.Center:
		offsetY = (parentSize.H - content.H) / 2
	case basic.End:
		offsetY = parentSize.H - content.H
	}

	for _, w := range c.Children {
		p := w.GetPos()
		p.Y += offsetY
		w.SetPos(p)
	}
}

func (c *Column) alignCross(parentSize basic.Size) {
	for _, w := range c.Children {
		size := w.GetSize()
		p := w.GetPos()

		switch c.CrossAlign {
		case basic.Start:
			continue
		case basic.Center:
			p.X = (parentSize.W - size.W) / 2
		case basic.End:
			p.X = parentSize.W - size.W
		}

		w.SetPos(p)
	}
}

func (c *Column) GetPos() basic.Point {
	return c.Pos
}

func (c *Column) SetPos(point basic.Point) {
	c.Pos = point
}

// calcula tamanho (apenas no init)
func (c *Column) calcSize() {
	var totalH float32
	var maxW float32

	for i, w := range c.Children {
		s := w.GetSize()

		totalH += s.H
		if i < len(c.Children)-1 {
			totalH += c.Spacing
		}

		if s.W > maxW {
			maxW = s.W
		}
	}

	//calcula size uma unica vez aqui
	c.size = basic.Size{W: maxW, H: totalH}
}

// GetSize retorna dimensões da column
func (c *Column) GetSize() basic.Size {
	return c.size
}

// init serve para um primeiro posicionamento dos elementos (start x start)
func (c *Column) init() {
	cursorY := float32(0)

	for i, w := range c.Children {
		size := w.GetSize()

		w.SetPos(basic.Point{
			X: 0,
			Y: cursorY,
		})

		cursorY += size.H
		if i < len(c.Children)-1 {
			cursorY += c.Spacing
		}
	}
	c.calcSize()
}
