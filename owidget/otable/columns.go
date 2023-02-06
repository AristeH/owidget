package otable

import (
	"fmt"
	"strconv"
)

// ColumnStyle - стиль колонки
type ColumnStyle struct {
	id      string  // id column
	name    string  // header
	formula string  // calculated expressions
	width   float32 // Ширина столбца
	bgColor string  // Цвет фона
	fgColor string  // Цвет текста
	tip     string  // Тип колонки(id, float, string, enum, date)
	visible bool    // Видимость
	edit    bool    // Редактирование колонки
	order   int16   // column output order
	sort    int8    // 0 - нет сортировки, 1 - возрастание, 2-убывание
}

// fillColumns - filling in columns from incoming data
func (t *owidget.OTable) fillColumns(d owidget.GetData) {
	columns := len(d.DataDescription[0]) // количество колонок таблицы
	//инициализация стиля
	t.ColumnStyle = make(map[string]*ColumnStyle)
	//ширина символа
	for i := 0; i < columns; i++ {
		cs := ColumnStyle{}
		cs.name = d.Data[0][i]
		cs.id = d.DataDescription[0][i]
		cs.bgColor = "" // индивидуальный цвет столбца фон
		cs.fgColor = ""
		cs.formula = d.DataDescription[3][i] // индивидуальный цвет текста столбца
		cs.tip = d.DataDescription[1][i]
		p, _ := strconv.Atoi(d.DataDescription[2][i]) // ширина столбца в символах
		cs.width = float32(p)                         // ширина колонки
		cs.visible = true                             // видимость столбца
		cs.edit = true                                // редактируемость столбца
		t.ColumnStyle[cs.name] = &cs
	}
	defer fmt.Println("Columns", t.ColumnStyle, "Finish")
}
