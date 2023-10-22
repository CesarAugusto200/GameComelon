package Models

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	playerInitialSize = 10
	playerColorRed    = 255
	playerColorGreen  = 0
	playerColorBlue   = 0
)

type Player struct {
	X, Y, Size int
	Circle     *canvas.Circle
}

func (p *Player) GetWidget() fyne.CanvasObject {
	return p.Circle
}

func NewPlayer() *Player {
    rand.Seed(time.Now().UnixNano())
    player := &Player{
        X:    rand.Intn(100),
        Y:    rand.Intn(100),
        Size: playerInitialSize,
    }
    player.initializeCircle() 
    return player
}

func (p *Player) initializeCircle() {
	p.Circle = canvas.NewCircle(color.RGBA{R: playerColorRed, G: playerColorGreen, B: playerColorBlue, A: 255})
	p.Circle.StrokeWidth = 1
	p.UpdateCirclePosition()
}

func (p *Player) UpdateCirclePosition() {
	p.Circle.Resize(fyne.NewSize(float32(10+p.Size), float32(10+p.Size)))
	p.Circle.Move(fyne.NewPos(float32(p.X), float32(p.Y)))
}

func (p *Player) MoveTowards(c *Cell) {
	dx := c.X - p.X
	dy := c.Y - p.Y

	distance := math.Sqrt(float64(dx*dx + dy*dy))

	if distance > 0 {
		p.X += int(float64(dx) / distance)
		p.Y += int(float64(dy) / distance)
	}

	p.UpdateCirclePosition()
	p.Circle.Refresh()
}

func (p *Player) Eat(c *Cell) bool {
	dx := p.X - c.X
	dy := p.Y - c.Y
	distance := math.Sqrt(float64(dx*dx + dy*dy))

	if distance < float64(p.Size+c.Size) {
		p.Size += 1
		p.UpdateCirclePosition()
		p.Circle.Refresh()
		return true
	}
	return false
}
