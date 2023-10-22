package Models

import (
    "image/color"
    "math"
    "math/rand"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
)

type Bot struct {
    X, Y, Size int
    Circle     *canvas.Circle
}

func (b *Bot) GetWidget() fyne.CanvasObject {
    return b.Circle
}

// aui es lo que hace que el bot pueda comer una celula
func (b *Bot) Eat(c *Cell) bool {
    dx := b.X - c.X
    dy := b.Y - c.Y
    distance := math.Sqrt(float64(dx*dx + dy*dy))

    if distance < float64(b.Size+c.Size) {
        b.Size += 1
        b.Circle.Resize(fyne.NewSize(float32(10+b.Size), float32(10+b.Size))) // Incrementa el tamaño del circulo cuando el bot come una celula
        b.Circle.Refresh()                                                   
        return true
    }
    return false
}




type Cell struct {
    X, Y, Size int
    Circle     *canvas.Circle 
}

func (c *Cell) GetWidget() fyne.CanvasObject {
    return c.Circle
}

func NewBot() *Bot {
    rand.Seed(time.Now().UnixNano())
    bot := &Bot{
        X:    rand.Intn(100),
        Y:    rand.Intn(100),
        Size: 10,
    }
    bot.Circle = canvas.NewCircle(color.RGBA{R: 0, G: 0, B: 255, A: 255}) // Crea un circulo azul para el bot
    bot.Circle.StrokeWidth = 2    //agreggo un borde para el circulo                                         
    bot.Circle.Resize(fyne.NewSize(float32(10+bot.Size), float32(10+bot.Size))) 
    bot.Circle.Move(fyne.NewPos(float32(bot.X), float32(bot.Y)))
    return bot
}

func NewCell(canvasWidth, canvasHeight int) *Cell {
	rand.Seed(time.Now().UnixNano())
	cell := &Cell{
		X:    rand.Intn(canvasWidth),
		Y:    rand.Intn(canvasHeight),
		Size: 10,
	}
    cell.Circle = canvas.NewCircle(color.RGBA{R: 0, G: 255, B: 0, A: 255}) // Crea un circulo verde para la celula
    cell.Circle.StrokeWidth = 2                                             
    cell.Circle.Resize(fyne.NewSize(float32(10+cell.Size), float32(10+cell.Size))) 
    cell.Circle.Move(fyne.NewPos(float32(cell.X), float32(cell.Y)))
    return cell
}
