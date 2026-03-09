package components

import (
	"fmt"

	"github.com/allanjose001/go-battleship/game/components/basic"
	inputhelper "github.com/allanjose001/go-battleship/game/util"
	"github.com/hajimehoshi/ebiten/v2"
)

// IconButton é um botão circular com ícone
type IconButton struct {
	img             *Image
	pos, currentPos basic.Point
	size            basic.Size
	callback        func()
}

func NewIconButton(iconPath string, pos basic.Point, size basic.Size, cb func()) *IconButton {
	img, err := NewImage(iconPath, basic.Point{}, size)
	if err != nil {
		return nil
	}
	return &IconButton{
		img:      img,
		pos:      pos,
		size:     size,
		callback: cb,
	}
}

func (b *IconButton) GetPos() basic.Point {
	return b.pos
}

func (b *IconButton) SetPos(p basic.Point) {
	b.pos = p
}

func (b *IconButton) GetSize() basic.Size {
	return b.size
}

func (b *IconButton) Update(offset basic.Point) {
	if b.img == nil {
		return
	}

	b.currentPos = b.pos.Add(offset)
	b.img.Update(b.currentPos)

	mx, my := ebiten.CursorPosition()

	if b.callback != nil && inputhelper.IsClicked(mx, my, b.currentPos, b.size) {
		b.callback()
	}
}

func (b *IconButton) Draw(screen *ebiten.Image) {
	if b.img == nil {
		return
	}

	b.img.Draw(screen)
}

// Helpers para criar botões específicos

func NewDeleteIconButton(point basic.Point, size basic.Size, cb func()) *IconButton {
	return NewIconButton("assets/images/apagar-simbolo.png", point, size, cb)
}

func NewPlayIconButton(point basic.Point, size basic.Size, cb func()) *IconButton {
	return NewIconButton("assets/images/botao-play.png", point, size, cb)
}

func (b *IconButton) SetIcon(path string) {
	if b.img == nil {
		return
	}
	err := b.img.SetTexture(path)
	if err != nil {
		fmt.Println("Erro ao trocar ícone:", err)
	}
}
