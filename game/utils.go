package game

const (
	screenWidth  = 800
	screenHeight = 800
)

type Vector struct {
	X float64
	Y float64
}

type Rect struct {
	X, Y, Width, Height float64
}

func NewRect(x, y, w, h float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.maxX() &&
		other.X <= r.maxX() &&
		r.Y <= other.maxY() &&
		other.Y <= r.maxY()
}

func (r Rect) maxX() float64 {
	return r.X + r.Width
}

func (r Rect) maxY() float64 {
	return r.Y + r.Height
}
