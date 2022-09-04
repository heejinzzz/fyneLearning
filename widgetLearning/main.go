package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"net/url"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Widget Learning")

	// accordion
	ctn := container.New(layout.NewVBoxLayout())
	acc := widget.NewAccordion(
		widget.NewAccordionItem("file", canvas.NewText("file content", color.Black)),
		widget.NewAccordionItem("song", canvas.NewText("song content", color.Black)),
		widget.NewAccordionItem("image", canvas.NewText("image content", color.Black)),
	)
	ctn.Add(acc)

	// button
	btn := widget.NewButtonWithIcon("cancel", theme.CancelIcon(), func() {

	})
	btn.Importance = widget.HighImportance
	ctn.Add(btn)

	// card
	jzxImage := canvas.NewImageFromResource(jzxImg)
	jzxImage.SetMinSize(fyne.NewSize(100, 300))
	card := widget.NewCard("Jisoo", "", jzxImage)
	ctn.Add(card)

	// check
	ck := widget.NewCheck("I have read the protocol.", func(b bool) {

	})
	ctn.Add(ck)

	// entry
	unameEntry := widget.NewEntry()
	unameEntry.SetPlaceHolder("user name")
	unameEntry.Validator = func(s string) error {
		if len(s) <= 3 {
			return errors.New("the user name is too short")
		}
		if len(s) > 24 {
			return errors.New("the user name is too long")
		}
		return nil
	}
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("password")
	ctn.Add(unameEntry)
	ctn.Add(passwordEntry)

	// fileIcon
	fileURI := storage.NewFileURI("images/myImage.png")
	fileIcon := widget.NewFileIcon(fileURI)
	ctn.Add(fileIcon)

	// form
	form := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)
	form.OnSubmit = func() {

	}
	form.OnCancel = func() {

	}
	ctn.Add(form)

	// hyperlink
	URL, err := url.Parse("www.baidu.com")
	if err != nil {
		panic(err)
	}
	link := widget.NewHyperlink("baidu", URL)
	link.Alignment = fyne.TextAlignCenter
	ctn.Add(link)

	// icon
	icon := widget.NewIcon(jzxImg)
	ctn.Add(icon)

	// label
	label := widget.NewLabel("this is a label. this is a label. this is a label. this is a label. this is a label.")
	label.Wrapping = fyne.TextWrapWord
	ctn.Add(label)

	// PopUpMenu
	menuItem1 := fyne.NewMenuItem("option1", func() {

	})
	menuItem2 := fyne.NewMenuItem("option2", func() {

	})
	menuItem3 := fyne.NewMenuItem("option3", func() {

	})
	menu := fyne.NewMenu("", menuItem1, menuItem2, menuItem3)
	menuPos := fyne.NewPos(20, 20)
	widget.ShowPopUpMenuAtPosition(menu, w.Canvas(), menuPos)

	// ProgressBar
	bar1 := widget.NewProgressBar()
	bar1.SetValue(0.72)
	bar2 := widget.NewProgressBarInfinite()
	ctn.Add(bar1)
	ctn.Add(bar2)

	// RadioGroup
	rg := widget.NewRadioGroup([]string{"option1", "option2"}, func(s string) {

	})
	ctn.Add(rg)

	// select
	selector := widget.NewSelect([]string{"option1", "option2"}, func(s string) {

	})
	selector.PlaceHolder = "select a option"
	ctn.Add(selector)

	// slider
	slider := widget.NewSlider(0, 100)
	ctn.Add(slider)

	// ToolBar
	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {

		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {

		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {

		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {

		}),
	)
	ctn.Add(toolBar)

	// Collection Widgets
	// list
	list := widget.NewList(
		func() int {
			return 3
		},
		func() fyne.CanvasObject {
			listItemIcon := widget.NewIcon(theme.ListIcon())
			listItemContent := widget.NewLabel("user:\tusername")
			return container.New(layout.NewHBoxLayout(), listItemIcon, listItemContent)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			users := []string{"ckl", "heejin", "test"}
			listItemCtn := object.(*fyne.Container)
			content := listItemCtn.Objects[1].(*widget.Label)
			content.SetText(users[id])
		},
	)
	list.OnSelected = func(id widget.ListItemID) {

	}
	ctn.Add(container.New(layout.NewPaddedLayout(), list))

	// table
	table := widget.NewTable(
		func() (int, int) {
			return 3, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("unit[x,y]")
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			TableItemLabel := object.(*widget.Label)
			TableItemLabel.SetText("unit[" + strconv.Itoa(id.Row) + "," + strconv.Itoa(id.Col) + "]")
		},
	)
	table.OnSelected = func(id widget.TableCellID) {

	}
	ctn.Add(container.New(layout.NewPaddedLayout(), table))

	// tree
	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			if id == "" {
				return []string{"cars", "trains"}
			}
			if id == "cars" {
				return []string{"ford", "tesla"}
			}
			if id == "trains" {
				return []string{"rocket", "tgv"}
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "cars" || id == "trains"
		},
		func(b bool) fyne.CanvasObject {
			return widget.NewLabel("object")
		},
		func(id widget.TreeNodeID, b bool, object fyne.CanvasObject) {
			treeItemLabel := object.(*widget.Label)
			treeItemLabel.SetText(id)
		},
	)
	tree.OnSelected = func(uid widget.TreeNodeID) {

	}
	ctn.Add(container.New(layout.NewPaddedLayout(), tree))

	// Container Widgets
	// AppTabs
	appTabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home Page")),
		container.NewTabItemWithIcon("More", theme.ContentAddIcon(), widget.NewLabel("More Page")),
	)
	ctn.Add(appTabs)

	// split
	vSplit := container.NewVSplit(widget.NewLabel("Top"), widget.NewLabel("Bottom"))
	ctn.Add(vSplit)
	hSplit := container.NewHSplit(widget.NewLabel("Top"), widget.NewLabel("Bottom"))
	ctn.Add(hSplit)

	scroll := container.NewScroll(ctn)
	w.SetContent(scroll)

	w.Resize(fyne.NewSize(500, 900))
	w.ShowAndRun()
}
