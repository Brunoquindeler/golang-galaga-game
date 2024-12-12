package game

import (
	"math/rand"

	"github.com/brunoquindeler/golang-galaga-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	image    *ebiten.Image
	position Vector
	speed    float64
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	maxLeftX := float64(image.Bounds().Max.X)
	maxRightX := screenHeight - float64(image.Bounds().Max.X)
	position := Vector{
		X: (maxLeftX + rand.Float64()*(maxRightX-maxLeftX)),
		Y: -100,
	}

	speed := (rand.Float64() * 10)

	return &Meteor{
		image:    image,
		position: position,
		speed:    speed,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.image, op)
}

func (m *Meteor) Collider() Rect {
	bounds := m.image.Bounds()

	return NewRect(m.position.X, m.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
