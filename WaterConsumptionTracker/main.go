package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"time"
)

var (
	currentDate               binding.String
	currentDayConsumptionData binding.Int
	addConsumptionData        binding.Int
	mondayData                binding.Int
	tuesdayData               binding.Int
	wednesdayData             binding.Int
	thursdayData              binding.Int
	fridayData                binding.Int
	saturdayData              binding.Int
	sundayData                binding.Int
)

func newCurrentDayConsumption() *canvas.Text {
	data := binding.IntToStringWithFormat(currentDayConsumptionData, "%dml")
	consumption, err := data.Get()
	if err != nil {
		log.Println(err)
	}
	label := canvas.NewText(consumption, theme.PrimaryColor())
	label.TextSize = 40
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	currentDayConsumptionData.AddListener(binding.NewDataListener(func() {
		newData := binding.IntToStringWithFormat(currentDayConsumptionData, "%dml")
		newConsumption, err := newData.Get()
		if err != nil {
			log.Println(err)
		}
		label.Text = newConsumption
		label.Refresh()
	}))
	return label
}

func newCurrentDate() *widget.Label {
	date := widget.NewLabelWithData(currentDate)
	date.Alignment = fyne.TextAlignCenter
	return date
}

func newAddEntry() *widget.Entry {
	entry := widget.NewEntryWithData(binding.IntToString(addConsumptionData))
	entry.Validator = func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		return nil
	}
	return entry
}

func newAddButton() *widget.Button {
	button := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		addConsumption, err := addConsumptionData.Get()
		if err != nil {
			log.Println(err)
		}
		currentDayConsumption, err := currentDayConsumptionData.Get()
		if err != nil {
			log.Println(err)
		}
		err = currentDayConsumptionData.Set(currentDayConsumption + addConsumption)
		if err != nil {
			log.Println(err)
		}
	})
	return button
}

func newAddContainer() *fyne.Container {
	entry := newAddEntry()
	unit := widget.NewLabel("ml")
	button := newAddButton()
	ctn := container.New(layout.NewHBoxLayout(), entry, unit, button)
	return ctn
}

func newHistory() *widget.Form {
	MondayData := widget.NewLabelWithData(binding.IntToStringWithFormat(mondayData, "%dml"))
	MondayData.Alignment = fyne.TextAlignTrailing
	TuesdayData := widget.NewLabelWithData(binding.IntToStringWithFormat(tuesdayData, "%dml"))
	TuesdayData.Alignment = fyne.TextAlignTrailing
	WednesdayData := widget.NewLabelWithData(binding.IntToStringWithFormat(wednesdayData, "%dml"))
	WednesdayData.Alignment = fyne.TextAlignTrailing
	ThursdayData := widget.NewLabelWithData(binding.IntToStringWithFormat(thursdayData, "%dml"))
	ThursdayData.Alignment = fyne.TextAlignTrailing
	FridayData := widget.NewLabelWithData(binding.IntToStringWithFormat(fridayData, "%dml"))
	FridayData.Alignment = fyne.TextAlignTrailing
	SaturdayData := widget.NewLabelWithData(binding.IntToStringWithFormat(saturdayData, "%dml"))
	SaturdayData.Alignment = fyne.TextAlignTrailing
	SundayData := widget.NewLabelWithData(binding.IntToStringWithFormat(sundayData, "%dml"))
	SundayData.Alignment = fyne.TextAlignTrailing
	history := widget.NewForm(
		widget.NewFormItem("Monday", MondayData),
		widget.NewFormItem("Tuesday", TuesdayData),
		widget.NewFormItem("Wednesday", WednesdayData),
		widget.NewFormItem("Thursday", ThursdayData),
		widget.NewFormItem("Friday", FridayData),
		widget.NewFormItem("Saturday", SaturdayData),
		widget.NewFormItem("Sunday", SundayData),
	)
	return history
}

func newHistoryCard() *widget.Card {
	history := newHistory()
	card := widget.NewCard("History", "Totals this week", history)
	return card
}

func newMainPage() *fyne.Container {
	title := newCurrentDayConsumption()
	date := newCurrentDate()
	addContainer := newAddContainer()
	card := newHistoryCard()
	ctn := container.New(layout.NewVBoxLayout(), title, date, addContainer, card)
	return ctn
}

func setCurrentDate() {
	date := time.Now().Format("Mon _2 Jan 2006")
	currentDate = binding.BindString(&date)
}

func setCurrentDayOfWeek() {
	switch time.Now().Format("Monday") {
	case "Monday":
		currentDayConsumptionData = mondayData
	case "Tuesday":
		currentDayConsumptionData = tuesdayData
	case "Wednesday":
		currentDayConsumptionData = wednesdayData
	case "Thursday":
		currentDayConsumptionData = thursdayData
	case "Friday":
		currentDayConsumptionData = fridayData
	case "Saturday":
		currentDayConsumptionData = saturdayData
	case "Sunday":
		currentDayConsumptionData = sundayData
	default:
		log.Println("Unknown day: ", time.Now().Format("Monday"))
	}
}

func initData() {
	mondayData = binding.NewInt()
	tuesdayData = binding.NewInt()
	wednesdayData = binding.NewInt()
	thursdayData = binding.NewInt()
	fridayData = binding.NewInt()
	saturdayData = binding.NewInt()
	sundayData = binding.NewInt()
	addConsumptionData = binding.NewInt()
}

func loadDataFromStorage(preferences fyne.Preferences) {
	mondayData = binding.BindPreferenceInt("mondayData", preferences)
	tuesdayData = binding.BindPreferenceInt("tuesdayData", preferences)
	wednesdayData = binding.BindPreferenceInt("wednesdayData", preferences)
	thursdayData = binding.BindPreferenceInt("thursdayData", preferences)
	fridayData = binding.BindPreferenceInt("fridayData", preferences)
	saturdayData = binding.BindPreferenceInt("saturdayData", preferences)
	sundayData = binding.BindPreferenceInt("sundayData", preferences)
	addConsumptionData = binding.NewInt()
}

func checkWeekChange(preferences fyne.Preferences) {
	year, week := time.Now().ISOWeek()
	lastUseYear := preferences.IntWithFallback("lastUseYear", 0)
	lastUseWeek := preferences.IntWithFallback("lastUseWeek", 0)
	if year != lastUseYear || week != lastUseWeek {
		preferences.SetInt("mondayData", 0)
		preferences.SetInt("tuesdayData", 0)
		preferences.SetInt("wednesdayData", 0)
		preferences.SetInt("thursdayData", 0)
		preferences.SetInt("fridayData", 0)
		preferences.SetInt("saturdayData", 0)
		preferences.SetInt("sundayData", 0)
	}
	preferences.SetInt("lastUseYear", year)
	preferences.SetInt("lastUseWeek", week)
}

func main() {
	a := app.NewWithID("com.heejinzzz.WaterConsumptionTracker")
	w := a.NewWindow("WaterConsumptionTracker")

	//initData()
	checkWeekChange(a.Preferences())
	loadDataFromStorage(a.Preferences())
	setCurrentDate()
	setCurrentDayOfWeek()

	mainPage := newMainPage()
	w.SetContent(container.NewScroll(mainPage))

	w.Resize(fyne.NewSize(500, 900))
	w.ShowAndRun()
}
