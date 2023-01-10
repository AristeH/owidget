package owidget

import (
	"fmt"
	"fyne.io/fyne/v2"
)

type GetData struct {
	Data            [][]string
	DataDescription [][]string
	Enum            map[string][]string
}

type ActiveWidget struct {
	tip string //bool, ce
	ce  *oEntry
	ti  *tappableIcon
	sel *oSelect
	t   *OTable
}

// FormData - данные формы
type FormData struct {
	ID           string             // ID - ГУИД формы
	Table        map[string]*OTable // Table - список таблиц формы
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
		Table: make(map[string]*OTable),
		//	ActiveContainer: &OTable{},
		ActiveWidget: &ActiveWidget{},
	}
	f.W = fyne.CurrentApp().NewWindow(header)
	AppValues[name] = &f
	fmt.Println("form ", name, "event ", "InitFormData()")
	return &f
}

func (f *FormData) NewOTable(name string, d GetData) *OTable {
	table := OTable{}
	table.CellColor = make(map[string]*CellColor)
	table.Enum = d.Enum
	table.Form = *f
	table.Edit = true
	fmt.Println("1table.Form ", len(d.Data), "NewOTable")
	f.Table[name] = &table
	table.fill(d)
	return &table
}
