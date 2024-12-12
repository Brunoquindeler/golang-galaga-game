package main

import (
	"github.com/brunoquindeler/golang-galaga-game/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
