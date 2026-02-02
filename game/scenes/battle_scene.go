package scenes

//controla o que acontece na batalha

import (
	"github.com/allanjose001/go-battleship/game/state"

	"github.com/hajimehoshi/ebiten/v2"
)

type BattleScene struct {
	state *state.GameState
}

func NewBattleScene(s *state.GameState) *BattleScene {
	return &BattleScene{state: s}
}

func (s *BattleScene) Update() error {
	return nil
}

func (s *BattleScene) Draw(screen *ebiten.Image) {
	s.state.PlayerBoard.Draw(screen)
	s.state.AIBoard.Draw(screen)
}
