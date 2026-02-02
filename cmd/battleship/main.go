package main

import (
	"log"

	"github.com/allanjose001/go-battleship/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(900, 650)
	ebiten.SetWindowTitle("Fúria dos Mares")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
