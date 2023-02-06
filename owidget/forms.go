package owidget

import (
	"fmt"

	"fyne.io/fyne/v2"
)

type GetData struct {
	Container       string
	Data            [][]string
	DataDescription [][]string
	Enum            map[string][]string
}

type ActiveWidget struct {
	tip string //bool, ce
	ce  *owidget.oEntry
	ti  *owidget.tappableIcon
	sel *owidget.oSelect
	t   *otable.OTable
}

// FormData - данные формы
type FormData struct {
	ID           string                    // ID - ГУИД формы
	Table        map[string]*otable.OTable // Table - список таблиц формы
	W            fyne.Window
	ActiveWidget *ActiveWidget
}

var AppValues = make(map[string]*FormData)

func GetApp() map[string]*FormData {
	return AppValues
}

func GetW(name string) fyne.Window {
	return AppValues[name].W
}

func PutListForm(name, header string) *FormData {
	f := FormData{
		ID:    name,
		Table: make(map[string]*otable.OTable),
		//	ActiveContainer: &OTable{},
		ActiveWidget: &ActiveWidget{},
	}
	f.W = fyne.CurrentApp().NewWindow(header)
	AppValues[name] = &f
	fmt.Println("form ", name, "event ", "InitFormData()")
	return &f
}

func (f *FormData) NewOTable(name string, d GetData) *otable.OTable {
	table := otable.OTable{}
	table.CellColor = make(map[string]*otable.CellColor)
	table.Enum = d.Enum
	table.Form = *f
	table.Edit = true
	fmt.Println("1table.Form ", len(d.Data), "NewOTable")
	f.Table[name] = &table
	table.fill(d)
	return &table
}
