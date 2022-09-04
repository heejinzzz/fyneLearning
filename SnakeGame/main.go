package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"time"
)

var (
	snakeParts    *[]fyne.CanvasObject
	gameContainer *fyne.Container
)

type direction int

const (
	ToLeft  direction = 1
	ToRight direction = 2
	ToUp    direction = 3
	ToDown  direction = 4
)

func init() {
	gameContainer = container.NewWithoutLayout()
	snakeParts = &[]fyne.CanvasObject{}
}

func setupGame() {
	for i := 0; i < 10; i++ {
		rect := canvas.NewRectangle(color.NRGBA{G: 0x66, A: 0xff})
		rect.Resize(fyne.NewSize(10, 10))
		rect.Move(fyne.NewPos(90, float32(50+i*10)))
		*snakeParts = append(*snakeParts, rect)
	}
	gameContainer = container.NewWithoutLayout(*snakeParts...)
}

func snakeMove(drct direction) {
	*snakeParts = (*snakeParts)[1:]
	snakeHead := (*snakeParts)[len(*snakeParts)-1]
	snakeNewPart := canvas.NewRectangle(color.NRGBA{G: 0x66, A: 0xff})
	snakeNewPart.Resize(fyne.NewSize(10, 10))
	if drct == ToLeft {
		snakeNewPart.Move(fyne.NewPos(snakeHead.Position().X-10, snakeHead.Position().Y))
	} else if drct == ToRight {
		snakeNewPart.Move(fyne.NewPos(snakeHead.Position().X+10, snakeHead.Position().Y))
	} else if drct == ToUp {
		snakeNewPart.Move(fyne.NewPos(snakeHead.Position().X, snakeHead.Position().Y-10))
	} else {
		snakeNewPart.Move(fyne.NewPos(snakeHead.Position().X, snakeHead.Position().Y+10))
	}
	*snakeParts = append(*snakeParts, snakeNewPart)
	gameContainer.RemoveAll()
	for _, object := range *snakeParts {
		gameContainer.Add(object)
	}
	gameContainer.Refresh()
}

func main() {
	a := app.New()

	w := a.NewWindow("Snake Game")
	setupGame()
	w.SetContent(gameContainer)

	var currentDirection direction = ToDown
	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeyLeft {
			snakeMove(ToLeft)
			currentDirection = ToLeft
		} else if event.Name == fyne.KeyRight {
			snakeMove(ToRight)
			currentDirection = ToRight
		} else if event.Name == fyne.KeyUp {
			snakeMove(ToUp)
			currentDirection = ToUp
		} else if event.Name == fyne.KeyDown {
			snakeMove(ToDown)
			currentDirection = ToDown
		}
	})
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			snakeMove(currentDirection)
		}
	}()

	w.Resize(fyne.NewSize(200, 200))
	//w.SetPadded(false)
	w.ShowAndRun()
}
