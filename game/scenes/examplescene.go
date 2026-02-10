package scenes

import (
	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/hajimehoshi/ebiten/v2"
)

type ExampleScene struct {
	container components.Widget
	col       components.Column
	row       components.Row
}

func (e *ExampleScene) OnEnter(_ Scene, size basic.Size) {

	//primeiro vou criar um alista de botões

	buttons := []components.Widget{
		components.NewButton(
			basic.Point{},
			basic.Size{W: 400, H: 70},
			"botão x",
			colors.Green,
			nil,
			func(b *components.Button) {},
		),
		components.NewButton(
			basic.Point{},
			basic.Size{W: 400, H: 70},
			"botão x",
			colors.Green,
			nil,
			func(b *components.Button) {},
		),
		components.NewButton(
			basic.Point{},
			basic.Size{W: 400, H: 70},
			"botão x",
			colors.Green,
			nil,
			func(b *components.Button) {},
		),
	}

	e.container = components.NewContainer(
		basic.Point{50, 50},
		basic.Size{W: 0.8 * size.W, H: 0.8 * size.H},
		10.0,
		colors.White,
		basic.Center,
		basic.Center,
		&e.col,
		func(c *components.Container) {},
	)
	e.col = *components.NewColumn(
		basic.Point{},
		10.0,
		basic.Size{W: 0.8 * size.W, H: 0.8 * size.H},
		basic.Start,
		basic.Start,
		buttons,
	)

	/*e.row = *components.NewRow(
		basic.Point{},
		10.0,
		size,
		basic.Center,
		basic.Center,
		[]components.Widget{
			e.container,
		},
	)*/

}

func (e *ExampleScene) OnExit(_ Scene) {
}

func (e *ExampleScene) Update() error {
	e.container.Update(basic.Point{})
	return nil
}

func (e *ExampleScene) Draw(screen *ebiten.Image) {
	e.container.Draw(screen)
}
