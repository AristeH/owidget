// Package owidget описание структур для таблицы и внутренние функции
package owidget

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/widget"
)

// TableStyle - стиль таблицы
type TableStyle struct {
	RowAlterColor string // Цвет строки четной
	HeaderColor   string // Цвет шапки
	RowColor      string // Цвет строки нечетной
}

// CellColor - цвета для ячейки
type CellColor struct {
	fgColor color.RGBA // Цвет текста
	bgColor color.RGBA // Цвет фона
}

// OTable - описание таблицы
type OTable struct {
	widget.BaseWidget
	ID          string                  // Имя таблицы уникальное в пределах формы
	Form        FormData                // Формa владелец таблицы
	ColumnStyle map[string]*ColumnStyle // Описание колонок
	TabStyle    TableStyle              // Цвета фонов таблицы(шапка, строка
	Data        map[string][]string     // Исходные данные таблицы
	Enum        map[string][]string     // Данные для колонки
	DataV       [][]string
	// Отображаемые данные(сортировка, фильтр) 1 столбец ID записи, 1 строка шапка
	Table *widget.Table // Таблица fyne
	// Header      *widget.Table           // Шапка таблицы пока не реализована
	// Footer      *widget.Table // когда удастся сделать скроллинг
	// left      *widget.Table
	Properties *OTable               // Таблица для редактирования описания колонок
	Tool       *widget.Toolbar       // Командная панель  таблицы
	Selected   widget.TableCellID    // Выделенная ячейка таблицы
	Edit       bool                  // Редактируемость таблицы
	CellColor  map[string]*CellColor // Индивидуальный массив отображения ячеек
	// wb         map[*widget.Button]int
}

// MakeTableData - функция заполняющая структуру OTable из входных данных
func (t *OTable) fill(d GetData) {
	colColumns := len(d.DataDescription[0])
	t.fillColumns(d)

	colV := 0 //количество видимых столбцов для пользователя
	for i := 0; i < colColumns; i++ {
		b := strings.HasPrefix(d.Data[0][i], "id_") //исключим столбцы с типом ID
		if !b {
			colV++
		}
	}
	t.Data = make(map[string][]string)
	t.DataV = make([][]string, len(d.Data))

	for i := 0; i < len(d.Data); i++ {
		datad := make([]string, colColumns)
		datav := make([]string, colV)
		v := 0
		for j := 0; j < colColumns; j++ {
			//спрячем id ссылки на другие таблицы
			b := strings.HasPrefix(d.DataDescription[0][j], "id_")
			if !b {
				datav[v] = d.Data[i][j]
				v++
			}
			datad[j] = d.Data[i][j]
		}
		t.Data[d.Data[i][0]] = datad //запишем в map
		t.DataV[i] = datav
		// Log.WithFields(logrus.Fields{"v": v}).Info("mt")
	}
	t.TabStyle.RowAlterColor = "RowAlterColor"
	t.TabStyle.HeaderColor = "HeaderColor"
	t.TabStyle.RowColor = "RowColor"
	t.Selected = widget.TableCellID{}
	t.MakeTableLabel()
}

// properties - свойства таблицы, описание колонок
func (t *OTable) properties() *GetData {
	colColumns := 10
	colRows := len(t.ColumnStyle)
	datag := make([][]string, colRows)
	cs := t.ColumnStyle
	i := 0
	for _, v := range cs {
		datag[i] = make([]string, colColumns)
		datag[i][0] = v.id
		datag[i][1] = v.name                     // заголовок
		datag[i][2] = v.tip                      // тип столбца
		datag[i][3] = v.formula                  // Формула
		datag[i][4] = v.fgColor                  // цвет теста столбца
		datag[i][5] = v.bgColor                  // цвет фона столбца
		datag[i][6] = fmt.Sprintf("%v", v.width) // ширина столбца в символах
		if v.visible {                           // видимость столбца
			datag[i][7] = "1"
		} else {
			datag[i][7] = "0"
		}
		if v.edit { // видимость столбца
			datag[i][8] = "1"
		} else {
			datag[i][8] = "0"
		}
		datag[i][9] = strconv.FormatInt(int64(v.order), 2) // порядок вывода
		i++
	}

	datag[0][0] = "id"
	datag[0][1] = "Header"
	datag[0][2] = "Type"
	datag[0][3] = "formula"
	datag[0][4] = "fgColor"
	datag[0][5] = "bgColor"
	datag[0][6] = "Width"
	datag[0][7] = "visible"
	datag[0][8] = "edit"
	datag[0][9] = "order"

	// инициализация описания данных таблицы
	datadescription := make([][]string, 4)
	for i := 0; i < 4; i++ {
		datadescription[i] = make([]string, colColumns)
	}

	// Name column
	datadescription[0][0] = "id"
	datadescription[0][1] = "Header"
	datadescription[0][2] = "Type"
	datadescription[0][3] = "Formula"
	datadescription[0][4] = "Color"
	datadescription[0][5] = "BGColor"
	datadescription[0][6] = "Width"
	datadescription[0][7] = "visible"
	datadescription[0][8] = "edit"
	datadescription[0][9] = "order"

	//  Type column
	datadescription[1][0] = "string"
	datadescription[1][1] = "string"
	datadescription[1][2] = "string"
	datadescription[1][3] = "string"
	datadescription[1][4] = "enum"
	datadescription[1][5] = "enum"
	datadescription[1][6] = "int"
	datadescription[1][7] = "bool"
	datadescription[1][8] = "bool"
	datadescription[1][9] = "int"

	// Width column
	datadescription[2][0] = "15"
	datadescription[2][1] = "15"
	datadescription[2][2] = "20"
	datadescription[2][3] = "10"
	datadescription[2][4] = "15"
	datadescription[2][5] = "15"
	datadescription[2][6] = "6"
	datadescription[2][7] = "4"
	datadescription[2][8] = "4"
	datadescription[2][9] = "3"

	//Formula column
	datadescription[3][0] = ""
	datadescription[3][1] = ""
	datadescription[3][2] = ""
	datadescription[3][3] = ""
	datadescription[3][4] = ""
	datadescription[3][5] = ""
	datadescription[3][6] = ""
	datadescription[3][7] = ""
	datadescription[3][8] = ""
	datadescription[3][9] = ""

	f := GetData{}
	f.Data = datag
	f.DataDescription = datadescription
	f.Enum = map[string][]string{
		"BGColor": Names,
		"Color":   Names,
	}
	return &f

}

// getColorCell - получим цвет фона и текста отображаемой ячейки
func (t *OTable) getColorCell(i widget.TableCellID) *CellColor {
	c := CellColor{}
	col := t.ColumnStyle[t.DataV[0][i.Col]]
	if col.fgColor != "" {
		c.fgColor = MapColor[col.fgColor]
	} else {
		c.fgColor = MapColor["black"]
	}
	//цвет фона строки
	if i.Row == 0 {
		c.bgColor = MapColor[t.TabStyle.HeaderColor]
	} else if i.Row%2 == 0 {
		c.bgColor = MapColor[t.TabStyle.RowAlterColor]
	} else {
		c.bgColor = MapColor[t.TabStyle.RowColor]
	}
	// цвет фона столбца

	if val, ok := MapColor[col.bgColor]; ok {
		c.bgColor = mix(val, c.bgColor)
	}
	// цвет ячейки
	id, ok := t.CellColor[strconv.Itoa(i.Row)+";"+strconv.Itoa(i.Col)]
	if ok {
		c = *id
	}

	// цвет выделенной ячейки
	if i == t.Selected {
		c.bgColor = MapColor["Selected"]
	}
	return &c
}

func (t *OTable) sortDown() {
	var temp []string
	x := t.DataV
	k := t.Selected.Col
	tip := t.ColumnStyle[t.DataV[0][t.Selected.Col]].tip
	n := len(x)
	usl := false
	if strings.HasPrefix(tip, "float") {
		tip = "float"
	}
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			switch tip {
			case "string", "bool", "id_string":
				usl = strings.ToUpper(x[i][k]) < strings.ToUpper(x[j][k])
			case "int":
				i1, _ := strconv.Atoi(x[i][k])
				i2, _ := strconv.Atoi(x[j][k])
				usl = i1 < i2
			case "float":
				i1, _ := strconv.ParseFloat(x[i][k], 32)
				i2, _ := strconv.ParseFloat(x[j][k], 32)
				usl = i1 < i2
			}
			if usl {
				temp = x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
}

func (t *OTable) sortUp() {
	var temp []string
	x := t.DataV
	k := t.Selected.Col
	tip := t.ColumnStyle[t.DataV[0][t.Selected.Col]].tip
	n := len(x)
	usl := false
	if strings.HasPrefix(tip, "float") {
		tip = "float"
	}
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			switch tip {
			case "string", "bool", "id_string":
				usl = strings.ToUpper(x[i][k]) > strings.ToUpper(x[j][k])
			case "int":
				i1, _ := strconv.Atoi(x[i][k])
				i2, _ := strconv.Atoi(x[j][k])
				usl = i1 > i2
			case "float":
				i1, _ := strconv.ParseFloat(x[i][k], 32)
				i2, _ := strconv.ParseFloat(x[j][k], 32)
				usl = i1 > i2
			}
			if usl {
				temp = x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
}
