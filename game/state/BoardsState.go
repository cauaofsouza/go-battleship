package state

import "github.com/allanjose001/go-battleship/game/shared/board"

type GameState struct {
	PlayerBoard *board.Board
	AIBoard     *board.Board
}

func NewGameState() *GameState {
	return &GameState{
		PlayerBoard: board.NewBoard(80, 150, 320),
		AIBoard:     board.NewBoard(500, 150, 320),
	}
}
