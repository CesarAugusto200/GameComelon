package scenes

import (
	"Comelon/Models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"time"
)

const (
	gameWidth  = 800
	gameHeight = 600
)

type GameScene struct {
	window fyne.Window
	bots   []*Models.Bot
	cells  []*Models.Cell
	c      *Models.Player
}

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{
		window: window,
		bots:   make([]*Models.Bot, 3),
		cells:  make([]*Models.Cell, 10),
	}

	for i := range scene.bots {
		scene.bots[i] = Models.NewBot()
	}
	for i := range scene.cells {
		scene.cells[i] = Models.NewCell(int(scene.window.Canvas().Size().Width), int(scene.window.Canvas().Size().Height))
	}

	scene.c = Models.NewPlayer()

	botWidgets := make([]fyne.CanvasObject, len(scene.bots))
	cellWidgets := make([]fyne.CanvasObject, len(scene.cells))
	playerWidget := scene.c.GetWidget()

	for i, bot := range scene.bots {
		botWidgets[i] = bot.GetWidget()
	}

	for i, cell := range scene.cells {
		cellWidgets[i] = cell.GetWidget()
	}

	content := container.NewWithoutLayout(append(append(botWidgets, cellWidgets...), playerWidget)...)

	if content != nil {
		scene.window.SetContent(content)
	} else {
		println("error")
	}

	scene.window.Canvas().SetOnTypedKey(scene.KeyTyped)

	go scene.RunBots()
	go scene.RunPlayer()
	go scene.generateNewCells() 

	return scene
}

func (s *GameScene) KeyTyped(keyEvent *fyne.KeyEvent) {
	switch keyEvent.Name {
	case fyne.KeyUp:
		s.c.Y -= 5
	case fyne.KeyDown:
		s.c.Y += 5
	case fyne.KeyLeft:
		s.c.X -= 5
	case fyne.KeyRight:
		s.c.X += 5
	}

	s.c.X = clamp(s.c.X, 0, gameWidth)
	s.c.Y = clamp(s.c.Y, 0, gameHeight)

	s.c.UpdateCirclePosition()

	s.c.Circle.Refresh()
}

func (s *GameScene) RunPlayer() {
	for {
		for i, cell := range s.cells {
			if s.c.Eat(cell) {
				
				cell.GetWidget().Hide()

				// Reemplaza la celula comida con una nueva en una posición aleatoria
				s.cells[i] = Models.NewCell(int(s.window.Canvas().Size().Width), int(s.window.Canvas().Size().Height))
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func (s *GameScene) generateNewCells() {
    for {
        // Genera nuevas celulas cada 10 segundos
        time.Sleep(time.Second * 10)

        for i := 0; i < 10; i++ {
            newCell := Models.NewCell(int(s.window.Canvas().Size().Width), int(s.window.Canvas().Size().Height))
            s.cells = append(s.cells, newCell)
        }
    }
}


func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}