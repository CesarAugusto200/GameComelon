package scenes

import (
    "fyne.io/fyne/v2"
    "math"
    "math/rand"
    "time"
    "Comelon/Models"
)



func (s *GameScene) RunBots() {
    for {
        for _, bot := range s.bots {
            // Encuentra la celula mas cercana al bot
            var closestCell *Models.Cell
            closestDistance := float64(s.window.Canvas().Size().Width) * float64(s.window.Canvas().Size().Height)

            for _, cell := range s.cells {
                dx := bot.X - cell.X
                dy := bot.Y - cell.Y
                distance := math.Sqrt(float64(dx*dx + dy*dy))

                if distance < closestDistance {
                    closestDistance = distance
                    closestCell = cell
                }
            }

            if closestCell != nil {
                dx := closestCell.X - bot.X
                dy := closestCell.Y - bot.Y
                distance := math.Sqrt(float64(dx*dx + dy*dy))

                if distance > 0 {
                    dirX := float32(float64(dx) / distance)
                    dirY := float32(float64(dy) / distance)
                    bot.X += int(dirX)
                    bot.Y += int(dirY)
                    bot.X = clamp(bot.X, 0, gameWidth)
                    bot.Y = clamp(bot.Y, 0, gameHeight)
                    bot.Circle.Move(fyne.NewPos(float32(bot.X), float32(bot.Y)))

                    if bot.Eat(closestCell) {
                        for i, cell := range s.cells {
                            if cell == closestCell {
                                s.cells[i] = Models.NewCell(int(s.window.Canvas().Size().Width), int(s.window.Canvas().Size().Height))
                            }
                        }

                        bot.Circle.Refresh()
                    } else {
                        bot.Circle.Refresh()
                    }
                }
            }

            if s.c.Size > bot.Size {
                dx := s.c.X - bot.X
                dy := s.c.Y - bot.Y
                distance := math.Sqrt(float64(dx*dx + dy*dy))

                if distance < float64(s.c.Size)+float64(bot.Size) {
                    for i, b := range s.bots {
                        if b == bot {
                            s.bots = append(s.bots[:i], s.bots[i+1:]...)
                            bot.Circle.Hide()
                            break
                        }
                    }
                }
            }

            // Mueve el bot aleatoriamente
            randomDirX := rand.Intn(3) - 1
            randomDirY := rand.Intn(3) - 1
            bot.X += randomDirX
            bot.Y += randomDirY
            bot.X = clamp(bot.X, 0, gameWidth)
            bot.Y = clamp(bot.Y, 0, gameHeight)
            bot.Circle.Move(fyne.NewPos(float32(bot.X), float32(bot.Y)))
            bot.Circle.Refresh()
        }

        time.Sleep(time.Millisecond * 10)
    }
}


func (s *GameScene) RunCells() {
    for {
        // Logica para el comportamiento de las células
        for _, cell := range s.cells {
            // Genera direcciones aleatorias
            randomDirX := rand.Intn(3) - 1
            randomDirY := rand.Intn(3) - 1

            cell.X += randomDirX
            cell.Y += randomDirY

            
            cell.X = clamp(cell.X, 0, gameWidth)
            cell.Y = clamp(cell.Y, 0, gameHeight)

            // Actualiza la posición del círculo de la celula
            cell.Circle.Move(fyne.NewPos(float32(cell.X), float32(cell.Y)))

            cell.Circle.Refresh() 
        }

        time.Sleep(time.Millisecond * 100)
    }
}


func (s *GameScene) CheckGameOver() {
    for {
        
        time.Sleep(time.Second * 1)
    }
}


