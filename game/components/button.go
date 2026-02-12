package components

import (
	"image/color"
	"strings"

	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	inputhelper "github.com/allanjose001/go-battleship/game/util"
	"github.com/hajimehoshi/ebiten/v2"
)

// Button struct que encapsula comportamento por meio de função callback, label, e um corpo que é um container
// não é necessario preencher posição se estiver sendo alinhado em container, row ou column
type Button struct {
	pos, currentPos            basic.Point
	size                       basic.Size
	label                      string
	backgroundColor, textColor color.Color
	CallBack                   func(*Button) //função que o botão chama
	hoverColor                 color.Color
	disabled, hovered, clicked bool
	scale                      float32        // para animar click
	body                       StylableWidget //um container por ex
}

func NewButton(
	pos basic.Point, //opcional
	size basic.Size, //pode ser nil/zero
	label string,
	color color.Color,
	textColor color.Color,
	cb func(*Button),

) *Button {
	bt := &Button{
		pos:             pos,
		size:            size,
		scale:           1.0,
		label:           label,
		backgroundColor: color,
		textColor:       textColor,
		CallBack:        cb,
		hoverColor:      colors.Lighten(color, 0.25),
	}

	bt.makeBody() //cria body com container e variaveis de button

	return bt
}
func (b *Button) GetPos() basic.Point {
	return b.pos
}

func (b *Button) SetPos(point basic.Point) {
	b.pos = point
}

func (b *Button) GetSize() basic.Size {
	return b.size
}

func (b *Button) Update(point basic.Point) {
	mouseX, mouseY := ebiten.CursorPosition() //ver como fazer com disabled

	b.currentPos = b.pos.Add(point)

	//TODO: colocar som de hovered
	b.hoverVerify(mouseX, mouseY)

	b.clicked = inputhelper.IsClicked(mouseX, mouseY, b.currentPos, b.size)

	b.body.Update(b.currentPos)

	if b.clicked {
		if b.CallBack != nil {
			b.CallBack(b)
		}
	}

}

func (b *Button) SetSize(sz basic.Size) {
	b.body.SetSize(sz)
	b.size = sz
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.body.Draw(screen)
}

// makeBody cria container com tamanho texto e cores designadas
func (b *Button) makeBody() {

	if b.textColor == nil {
		b.textColor = colors.White
	}
	//corpo do botão com container
	b.body = NewContainer(
		b.pos,
		b.size, //tamanho original fica guardado
		16.0,
		b.backgroundColor,
		basic.Center,
		basic.Center,
		NewText(
			basic.Point{},
			strings.ToUpper(b.label),
			b.textColor,
			18, //VER SE ESSA FONTE DA
		),
		func(c *Container) {
			//fazer aqui relação com callback
		},
	)
}

// Hover verifica se o mouse está sob o botão
func (b *Button) hoverVerify(mouseX, mouseY int) {
	b.hovered = inputhelper.IsHovered(mouseX, mouseY, b.currentPos, b.size)

	if b.hovered {
		b.body.SetColor(b.hoverColor)
	} else {
		b.body.SetColor(b.backgroundColor)
	}
}
