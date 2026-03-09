package components

import (
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/hajimehoshi/ebiten/v2"
)

type Cursor struct {
	img     *Image
	hotspot basic.Point
}

func NewCursor() (*Cursor, error) {

	img, err := NewImage("assets/images/mouse.png", basic.Point{}, basic.Size{45, 45})
	if err != nil {
		return nil, err
	}

	return &Cursor{
		img:     img,
		hotspot: basic.Point{5, 5},
	}, nil
}

func (c *Cursor) GetPos() basic.Point {
	return c.img.GetPos()
}

func (c *Cursor) SetPos(p basic.Point) {
	c.img.SetPos(p)
}

func (c *Cursor) GetSize() basic.Size {
	return c.img.GetSize()
}

func (c *Cursor) Update(offset basic.Point) {

	x, y := ebiten.CursorPosition()

	mouse := basic.Point{
		X: float32(x),
		Y: float32(y),
	}

	// desloca para que a ponta do cursor seja o ponto real de clique
	pos := mouse.Sub(c.hotspot)

	c.img.SetPos(pos)
	c.img.Update(offset)
}

func (c *Cursor) Draw(screen *ebiten.Image) {
	c.img.Draw(screen)
}
