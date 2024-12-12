package game

import (
	"github.com/brunoquindeler/golang-galaga-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite
	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfWidth,
		Y: 500,
	}

	return &Player{
		image:             image,
		position:          position,
		game:              game,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0
	bounds := p.image.Bounds()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		if p.position.X > 0 {
			p.position.X -= speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		if p.position.X < float64(screenWidth-bounds.Dx()) {
			p.position.X += speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		if p.position.Y < float64(screenHeight-bounds.Dy()) {
			p.position.Y += speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		if p.position.Y > 0 {
			p.position.Y -= speed
		}
	}

	p.laserLoadingTimer.Update()
	if (ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()

		halfWidth := float64(bounds.Dx()) / 2
		halfHeight := float64(bounds.Dy()) / 2

		spawnPosition := Vector{
			X: p.position.X + halfWidth,
			Y: p.position.Y - halfHeight/2,
		}

		laser := NewLaser(spawnPosition)

		p.game.AddLaser(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(p.position.X, p.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
