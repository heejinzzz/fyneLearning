package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
	"sync"
)

var (
	a  fyne.App
	w  fyne.Window
	wg sync.WaitGroup
)

func imageItem(name string, image fyne.CanvasObject) fyne.CanvasObject {
	imageName := canvas.NewText(name, color.Gray{Y: 128})
	imageName.Alignment = fyne.TextAlignCenter
	imageName.TextSize = 12
	ctn := container.New(layout.NewBorderLayout(nil, imageName, nil, nil), imageName, image)
	return ctn
}

func imageContent(imageItems ...fyne.CanvasObject) *fyne.Container {
	ctn := container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 90)), imageItems...)
	return ctn
}

func imageContentScroll(imageContent fyne.CanvasObject) *container.Scroll {
	scroll := container.NewScroll(imageContent)
	return scroll
}

func imageBoardFoot(folderName string) fyne.CanvasObject {
	foot := canvas.NewText(folderName, color.Gray{Y: 128})
	return foot
}

func imageBoardHead() fyne.CanvasObject {
	background := canvas.NewRectangle(color.Gray{Y: 128})
	title := canvas.NewText("Image Browser", color.White)
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter
	head := container.New(layout.NewPaddedLayout(), background, title)
	return head
}

func imageBoard(head, foot, left, right fyne.CanvasObject, content ...fyne.CanvasObject) fyne.CanvasObject {
	board := container.New(layout.NewBorderLayout(head, foot, left, right))
	if head != nil {
		board.Add(head)
	}
	if foot != nil {
		board.Add(foot)
	}
	if left != nil {
		board.Add(left)
	}
	if right != nil {
		board.Add(right)
	}
	for _, object := range content {
		board.Add(object)
	}
	return board
}

func selectImgBtn() fyne.CanvasObject {
	btn := widget.NewButtonWithIcon("select folder", theme.FolderIcon(), func() {
		showImages()
	})
	ctn := container.New(layout.NewCenterLayout(), btn)
	return ctn
}

func showImages() {
	dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(errors.New("Fail to open folder: "+uri.String()+"\nError: "+err.Error()), w)
			return
		}
		head := imageBoardHead()
		foot := imageBoardFoot("Directory: " + uri.Path())
		content := imageContent()
		scroll := imageContentScroll(content)
		board := imageBoard(head, foot, nil, nil, scroll)
		w.SetContent(board)
		URIs, err := uri.List()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		for _, URI := range URIs {
			if isFile(URI) {
				wg.Add(1)
				go loadImage(content, URI)
			}
		}
		wg.Wait()
	}, w).Show()
}

func isFile(uri fyne.URI) bool {
	extension := strings.ToLower(uri.Extension())
	return extension == ".png" || extension == ".jpg" || extension == ".jpeg" || extension == ".gif"
}

func loadImage(boardContent *fyne.Container, imgURI fyne.URI) {
	defer wg.Done()
	imgResource, err := storage.LoadResourceFromURI(imgURI)
	if err != nil {
		dialog.ShowError(errors.New("Fail to open file: "+imgURI.String()+"\nError: "+err.Error()), w)
		return
	}
	img := canvas.NewImageFromResource(imgResource)
	imgName := imgURI.Name()
	if len(imgName) > 12 {
		imgName = imgName[:12] + "..."
	}
	imgItem := imageItem(imgName, img)
	boardContent.Add(imgItem)
	boardContent.Refresh()
}

func mainMenu() *fyne.MainMenu {
	menu := fyne.NewMainMenu(fyne.NewMenu("File", fyne.NewMenuItem("Open Directory...", func() {
		showImages()
	})))
	return menu
}

func main() {
	a = app.New()
	a.Settings().SetTheme(myTheme{})
	w = a.NewWindow("ImageBrowser")

	btn := selectImgBtn()
	head := imageBoardHead()
	board := imageBoard(head, nil, nil, nil, btn)
	w.SetContent(board)
	w.SetMainMenu(mainMenu())

	w.Resize(fyne.NewSize(300, 500))
	w.ShowAndRun()
}
