package board

//responsabilidade de desenhar o tabuleiro.

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

func (b *Board) Draw(screen *ebiten.Image) {
	cellSize := b.Size / Cols

	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			x := b.X + float64(j)*cellSize
			y := b.Y + float64(i)*cellSize

			col := color.RGBA{173, 216, 230, 255} // azul claro

			ebitenutil.DrawRect(screen, x, y, cellSize-2, cellSize-2, col)
		}
	}

	// labels
	labelSize := float64(cellSize) * 0.35
	if labelSize < 10 {
		labelSize = 10
	}
	tt, _ := opentype.Parse(goregular.TTF)
	face, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    labelSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	labelColor := color.Black

	// topo: letras A-H
	for j := 0; j < Cols; j++ {
		ch := string(rune('A' + j))
		bounds := text.BoundString(face, ch)
		w := bounds.Dx()
		x := b.X + float64(j)*cellSize + cellSize/2 - float64(w)/2
		y := b.Y - 6
		text.Draw(screen, ch, face, int(x), int(y), labelColor)
	}

	// esquerda: números 1-7
	for i := 0; i < Rows; i++ {
		num := strconv.Itoa(i + 1)
		bounds := text.BoundString(face, num)
		h := bounds.Dy()
		x := b.X - 14
		y := b.Y + float64(i)*cellSize + cellSize/2 + float64(h)/2 - 2
		text.Draw(screen, num, face, int(x), int(y), labelColor)
	}
}
