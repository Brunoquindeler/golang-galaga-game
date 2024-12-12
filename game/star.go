package game

import (
	"math/rand"

	"github.com/brunoquindeler/golang-galaga-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct {
	image    *ebiten.Image
	position Vector
	speed    float64
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]

	minLeftX := float64(image.Bounds().Max.X)
	maxRightX := screenHeight - float64(image.Bounds().Max.X)
	position := Vector{
		X: (minLeftX + rand.Float64()*(maxRightX-minLeftX)),
		Y: -100,
	}

	speed := (rand.Float64() * 5)

	return &Star{
		image:    image,
		position: position,
		speed:    speed,
	}
}

func (s *Star) Update() {
	s.position.Y += s.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(s.position.X, s.position.Y)

	screen.DrawImage(s.image, op)
}

func (s *Star) Collider() Rect {
	bounds := s.image.Bounds()

	return NewRect(s.position.X, s.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
