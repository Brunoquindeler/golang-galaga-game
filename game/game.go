package game

import (
	"fmt"

	"github.com/brunoquindeler/golang-galaga-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	stars            []*Star
	meteorSpawnTimer *Timer
	starSpawnTimer   *Timer
	score            uint
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(100),
		starSpawnTimer:   NewTimer(1000),
	}
	player := NewPlayer(g)
	g.player = player

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	// Lasers
	for i, laser := range g.lasers {
		if laser.position.Y < -100 {
			g.lasers = append(g.lasers[:i], g.lasers[i+1:]...)
		}
		laser.Update()
	}

	// Meteors
	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}

	// Stars
	g.starSpawnTimer.Update()
	if g.starSpawnTimer.IsReady() {
		g.starSpawnTimer.Reset()

		s := NewStar()
		g.stars = append(g.stars, s)
	}

	for i, s := range g.stars {
		s.Update()
		if s.Collider().Intersects(g.player.Collider()) {
			g.score += 5
			g.stars = append(g.stars[:i], g.stars[i+1:]...)
		}
	}

	// Collider Laser x Meteor
	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score++
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, laser := range g.lasers {
		laser.Draw(screen)
	}

	for _, meteor := range g.meteors {
		meteor.Draw(screen)
	}

	for _, star := range g.stars {
		star.Draw(screen)
	}

	// text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.FontUi, 20, 50, color.White)
	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), text.NewGoXFace(assets.FontUi), &text.DrawOptions{
		DrawImageOptions: ebiten.DrawImageOptions{},
		LayoutOptions:    text.LayoutOptions{},
	})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.lasers = nil
	g.meteors = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
}
