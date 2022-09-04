package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	taskList    []*task
	todoList    *widget.List
	taskDetail  *widget.Form
	toolBar     *widget.Toolbar
	currentTask *task
	board       *fyne.Container
	w           fyne.Window
)

type task struct {
	title       string
	description string
	done        bool
	category    taskCategory
	priority    taskPriority
	due         string
	completion  float64
}

type taskCategory string

const (
	taskCategoryDaily taskCategory = "Daily"
	taskCategoryStudy taskCategory = "Study"
	taskCategoryWork  taskCategory = "Work"
	taskCategoryOther taskCategory = "Other"
)

type taskPriority string

const (
	taskPriorityLow  taskPriority = "Low"
	taskPriorityMid  taskPriority = "Mid"
	taskPriorityHigh taskPriority = "High"
)

func newEmptyTask() *task {
	return &task{
		title:       "New Task",
		description: "Write your description of this task here.",
		done:        false,
	}
}

func newTodoList() fyne.CanvasObject {
	todos := widget.NewList(
		func() int {
			return len(taskList)
		},
		func() fyne.CanvasObject {
			title := widget.NewLabel("Task Title")
			//title.Wrapping = fyne.TextWrapWord
			check := widget.NewCheck("", func(b bool) {})
			ctn := container.New(layout.NewHBoxLayout(), title, check)
			return ctn
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			ctn := object.(*fyne.Container)
			title := ctn.Objects[0].(*widget.Label)
			title.SetText(taskList[id].title)
			check := ctn.Objects[1].(*widget.Check)
			check.OnChanged = func(b bool) {
				taskList[id].done = b
			}
			check.SetChecked(taskList[id].done)
			ctn.Refresh()
		},
	)
	todos.OnSelected = func(id widget.ListItemID) {
		currentTask = taskList[id]
		taskDetail = newTaskDetail().(*widget.Form)
		board = newBoard(toolBar, todoList, taskDetail).(*fyne.Container)
		w.SetContent(board)
	}
	return todos
}

func newTaskDetail() fyne.CanvasObject {
	title := widget.NewEntry()
	title.SetText(currentTask.title)
	title.OnChanged = func(s string) {
		currentTask.title = s
		todoList.Refresh()
	}
	description := widget.NewMultiLineEntry()
	description.Wrapping = fyne.TextWrapWord
	description.SetMinRowsVisible(6)
	description.SetText(currentTask.description)
	description.OnChanged = func(s string) {
		currentTask.description = s
	}
	category := widget.NewSelect([]string{"Daily", "Study", "Work", "Other"}, func(s string) {
		currentTask.category = taskCategory(s)
	})
	if currentTask.category != "" {
		category.SetSelected(string(currentTask.category))
	}
	priority := widget.NewRadioGroup([]string{"Low", "Mid", "High"}, func(s string) {
		currentTask.priority = taskPriority(s)
	})
	if currentTask.priority != "" {
		priority.SetSelected(string(currentTask.priority))
	}
	due := widget.NewEntry()
	due.SetText(currentTask.due)
	due.OnChanged = func(s string) {
		currentTask.due = s
	}
	completion := widget.NewSlider(0, 100)
	completion.SetValue(currentTask.completion)
	completion.OnChanged = func(f float64) {
		currentTask.completion = f
	}
	detail := widget.NewForm(
		widget.NewFormItem("Title", title),
		widget.NewFormItem("Description", description),
		widget.NewFormItem("Category", category),
		widget.NewFormItem("Priority", priority),
		widget.NewFormItem("Due", due),
		widget.NewFormItem("Completion", completion),
	)
	return detail
}

func newToolBar() fyne.CanvasObject {
	tools := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			taskList = append(taskList, newEmptyTask())
			todoList.Refresh()
			todoList.Select(len(taskList) - 1)
		}),
	)
	return tools
}

func newBoard(toolBar fyne.CanvasObject, todoList fyne.CanvasObject, taskDetail fyne.CanvasObject) fyne.CanvasObject {
	return container.New(layout.NewBorderLayout(toolBar, nil, todoList, nil), toolBar, todoList, taskDetail)
}

func main() {
	a := app.New()
	w = a.NewWindow("TaskList")

	taskList = []*task{newEmptyTask()}
	currentTask = taskList[0]
	toolBar = newToolBar().(*widget.Toolbar)
	taskDetail = newTaskDetail().(*widget.Form)
	todoList = newTodoList().(*widget.List)
	todoList.Select(0)
	board = newBoard(toolBar, todoList, container.NewScroll(taskDetail)).(*fyne.Container)
	w.SetContent(board)

	w.Resize(fyne.NewSize(500, 900))
	w.ShowAndRun()
}
