package docx

import (
	"errors"
	//+1"fmt"
	"github.com/aymerick/raymond"
	"github.com/razer96/docxt/graph"
	"reflect"
	"regexp"
	"strings"
)

var (
	rxTemplateItem = regexp.MustCompile(`\{\{\s*([\w|\.|$]+)\s*\}\}`)
	rxMergeCellV   = regexp.MustCompile(`\[\s?v-merge\s?\]`)
	rxMergeIndex   = regexp.MustCompile(`\[\s?index\s?:\s?[\d|\.|\,|\$]+\s?\]`)
	rxBrCellV      = regexp.MustCompile(`\[\s?BR\s?\]`)
)

// Функционал шаблонизатора
func renderTemplateDocument(document *Document, v interface{}) error {
	if document != nil {
		// Проходимся по структуре документа
		for _, item := range document.Body.Items {
			if err := renderDocItem(item, v); err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("Not valid template document")
}

func renderTemplateHeader(header *Header, v interface{}) error {
	if header != nil {
		for _, item := range header.Items {
			if err := renderDocItem(item, v); err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("Not valid template document")
}

// Поиск элементов шаблона и спаивания текстовых элементов
func findTemplatePatternsInParagraph(p *ParagraphItem) {
	if p != nil {
		// Перебор элементов параграфа и поиск начал {{ и конца }}
		var startItem *RecordItem
		for index := 0; index < len(p.Items); index++ {
			i := p.Items[index]
			if i.Type() == Record {
				record := i.(*RecordItem)
				if record != nil {
					if startItem != nil {
						startItem.Text.Value += record.Text.Value
						// Удаляем элемент
						p.Items = append(p.Items[:index], p.Items[index+1:]...)
						// Проверка на конец
						closeIndex := strings.Index(startItem.Text.Value, "}}")
						if closeIndex < 0 {
							index--
							continue
						}
						anotherOpen := strings.Index(startItem.Text.Value[closeIndex:], "{{")
						if anotherOpen > 0 {
							//multiple template
							index--
							continue
						}

						startItem = nil

					} else {
						if strings.Index(record.Text.Value, "{{") >= 0 {
							startItem = record
							continue
						}
					}
				}
			}
			//startItem = nil
		}
	}
}

// Рендер элемента документа
func renderDocItem(item DocItem, v interface{}) error {
	switch elem := item.(type) {
	// Параграф
	case *ParagraphItem:
		{
			findTemplatePatternsInParagraph(elem)
			for _, i := range elem.Items {
				if err := renderDocItem(i, v); err != nil {
					return err
				}
			}
		}
	// Запись
	case *RecordItem:
		{
			if len(elem.Text.Value) > 0 {
				if rxTemplateItem.MatchString(elem.Text.Value) {
					text := modeTemplateText(elem.Text.Value)
					switch v.(type) {
					case *map[string]interface{}:
						qoute := strings.Index(text, "{{")
						first_ := strings.Index(text, "_")
						if first_ > 0 {
							text = text[:qoute+3] + text[first_+1:]
						}
					}

					out, err := raymond.Render(text, v)
					if err != nil {
						return err
					}
					elem.Text.Value = out
				}
			}
		}
	// Таблица
	case *TableItem:
		{
			for rowIndex := 0; rowIndex < len(elem.Rows); rowIndex++ {
				row := elem.Rows[rowIndex]
				if row != nil {
					// Если массив
					if obj, name, ok := haveArrayInRow(row, v); ok {
						lines := objToLines(obj, name)
						template := row.Clone()
						currentRow := row
						for _, line := range lines {
							if currentRow == nil {
								currentRow = template.Clone()
								// Insert Row
								elem.Rows = append(elem.Rows[:rowIndex], append([]*TableRow{currentRow}, elem.Rows[rowIndex:]...)...)
							}
							if err := renderRow(currentRow, &line); err != nil {
								return err
							}
							currentRow = nil
							rowIndex++
						}
						template = nil
						// Балансируем индекс
						rowIndex--
						continue
					}
					// Если нет
					if err := renderRow(row, v); err != nil {
						return err
					}
				}
			}
			// После обхода таблицы проходимся по ячейкам и проверяем merge флаги
			// С конца таблицы, проверяем по ячейкам
			for rowIndex := len(elem.Rows) - 1; rowIndex >= 0; rowIndex-- {
				setBoldRight := false
				// Обходим ячейки
				for cellIndex, cell := range elem.Rows[rowIndex].Cells {
					if len(cell.Items) > 0 {
						plainText := plainTextFromTableCell(cell)

						// Если найден флаг соединения
						if rxMergeCellV.MatchString(plainText) {
							if rowIndex > 0 {
								topCell := elem.Rows[rowIndex-1].Cells[cellIndex]
								if topCell != nil {
									if plainText == plainTextFromTableCell(topCell) {
										cell.Params.VerticalMerge = new(StringValue)
										for _, i := range cell.Items {
											clearTextFromDocItem(i)
										}
										continue
									}
								}
							}
							cell.Params.VerticalMerge = new(StringValue)
							cell.Params.VerticalMerge.Value = "restart"
							removeTemplateFromCell(rxMergeCellV, cell)
							removeTemplateFromCell(rxMergeIndex, cell)
						}
						// Проверка на флаг усановки жирного шрифта во всех ячейках справа
						if rxBrCellV.MatchString(plainText) {
							setBoldRight = !setBoldRight
							removeTemplateFromCell(rxBrCellV, cell)
						}
						// Если флаг выставлен, применяем жирный стиль у шрифта
						if setBoldRight {
							setBoldToCell(true, cell)
						}
					}
				}
			}
		}
	}
	return nil
}

func clearTextFromDocItem(item DocItem) {
	if item != nil {
		switch elem := item.(type) {
		case *ParagraphItem:
			{
				for _, i := range elem.Items {
					clearTextFromDocItem(i)
				}
			}
		case *RecordItem:
			{
				elem.Text.Value = ""
			}
		}
	}
}

func setBoldToCell(bold bool, cell *TableCell) {
	if cell != nil {
		for _, item := range cell.Items {
			setBoldToDocItem(bold, item)
		}
	}
}

func setBoldToDocItem(bold bool, item DocItem) {
	if item != nil {
		switch elem := item.(type) {
		case *ParagraphItem:
			{
				for _, i := range elem.Items {
					setBoldToDocItem(bold, i)
				}
			}
		case *RecordItem:
			{
				if bold {
					if elem.Params.Bold == nil {
						elem.Params.Bold = new(EmptyValue)
					}
					if elem.Params.BoldCS == nil {
						elem.Params.BoldCS = new(EmptyValue)
					}
				} else {
					if elem.Params.Bold != nil {
						elem.Params.Bold = nil
					}
					if elem.Params.BoldCS != nil {
						elem.Params.BoldCS = nil
					}
				}
			}
		}
	}
}

// removeTemplateFromCell - очищаяем контент ячейки от шаблона
func removeTemplateFromCell(template *regexp.Regexp, cell *TableCell) {
	if cell != nil && template != nil {
		for _, item := range cell.Items {
			removeTemplateFromDocItem(template, item)
		}
	}
}

// removeTemplateFromDocItem - очищаяем контент элемента документа от шаблона
func removeTemplateFromDocItem(template *regexp.Regexp, item DocItem) {
	if item != nil && template != nil {
		switch elem := item.(type) {
		case *ParagraphItem:
			{
				for _, i := range elem.Items {
					removeTemplateFromDocItem(template, i)
				}
			}
		case *RecordItem:
			{
				elem.Text.Value = template.ReplaceAllString(elem.Text.Value, "")
			}
		}
	}
}

// objToLines - раскладываем объект на строки
func objToLines(v interface{}, name string) []map[string]interface{} {
	node := new(graph.Node)
	node.FromObject(v)
	return node.ListMap()
}

// renderRow - вывод строки таблицы
func renderRow(row *TableRow, v interface{}) error {
	if row != nil {
		for _, cell := range row.Cells {
			if cell != nil {
				for _, item := range cell.Items {
					if err := renderDocItem(item, v); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// Модифицируем текст шаблона
func modeTemplateText(tpl string) string {
	//fmt.Println("Mode: ", tpl)
	tpl = strings.Replace(tpl, "{{", "{{{", -1)
	tpl = strings.Replace(tpl, "}}", "}}}", -1)
	tpl = strings.Replace(tpl, ".", "_", -1)
	return strings.Replace(tpl, ":length", "_length", -1)
}

// haveArrayInRow - содержится ли массив в строке
func haveArrayInRow(row *TableRow, v interface{}) (interface{}, string, bool) {
	if row != nil {
		for _, cell := range row.Cells {
			if match := rxTemplateItem.FindStringSubmatch(plainTextFromTableCell(cell)); match != nil && len(match) > 1 {
				names := strings.Split(match[1], ".")
				if len(names) > 0 {
					t := reflect.TypeOf(v)
					val := reflect.ValueOf(v)
					var lastVal reflect.Value
					for _, name := range names {
						if t.Kind() == reflect.Map {
							val = val.MapIndex(reflect.ValueOf(name))
							t = reflect.TypeOf(val.Interface())
						} else {
							t = findType(t, name)
							val, _ = findValue(val, name)
						}

						if t != nil {
							if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
								if lastVal.IsValid() {
									return lastVal.Interface(), name, true
								}
								return val.Interface(), name, true
							}
						} else {
							break
						}
						lastVal = val
					}
				}
			}
		}
	}
	return nil, "", false
}

// Простой текс у ячейки
func plainTextFromTableCell(cell *TableCell) string {
	var result string
	if cell != nil {
		for _, item := range cell.Items {
			result += item.PlainText()
		}
	}
	return result
}

// findType - получаем тип по имени
func findType(t reflect.Type, name string) reflect.Type {
	kind := t.Kind()
	// Если это ссылка, то получаем истенный тип
	if kind == reflect.Ptr || kind == reflect.Interface {
		t = t.Elem()
	}
	kind = t.Kind()
	if kind == reflect.Struct {
		if field, ok := t.FieldByName(name); ok {
			return field.Type
		}
	}

	return nil
}

// findValue - получаем значение по имени
func findValue(v reflect.Value, name string) (reflect.Value, bool) {
	kind := v.Type().Kind()
	// Если это ссылка, то получаем истенный тип
	if kind == reflect.Ptr || kind == reflect.Interface {
		v = v.Elem()
	}
	kind = v.Type().Kind()
	if kind == reflect.Struct {
		v := v.FieldByName(name)
		if v.IsValid() {
			return v, true
		}
	}
	return v, false
}
