package game

//estado global do jogo

import (
	"image/color"

	"github.com/allanjose001/go-battleship/game/scenes"
	"github.com/allanjose001/go-battleship/game/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	scene *scenes.BattleScene
}

func NewGame() *Game {
	st := state.NewGameState()
	return &Game{scene: scenes.NewBattleScene(st)}
}

func (g *Game) Update() error {
	if g.scene != nil {
		return g.scene.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 122, G: 133, B: 180, A: 255})
	if g.scene != nil {
		g.scene.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 900, 650
}
