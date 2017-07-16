package docx

import (
	"encoding/xml"
	"errors"
)

// TableItem - элемент таблици
type TableItem struct {
	Params TableParams `xml:"tblPr"`
	Grid   TableGrid   `xml:"tblGrid"`
	Rows   []*TableRow `xml:"tr,omitempty"`
}

// TableGrid - Grid table
type TableGrid struct {
	Cols []*WidthValue `xml:"gridCol,omitempty"`
}

type WTableGrid struct {
	Cols []*WWidthValue `xml:"w:gridCol,omitempty"`
}

func (g *TableGrid) ToWGirdParams() *WTableGrid {
	gw := WTableGrid{}
	for _, col := range g.Cols {
		gw.Cols = append(gw.Cols, (*WWidthValue)(col))
	}
	return &gw
}

// TableParamsEx - Other params table
type TableParamsEx struct {
	Shadow ShadowValue `xml:"shd"`
}

type WTableParamsEx struct {
	Shadow ShadowValue `xml:"w:shd"`
}

// Tag - имя тега элемента
func (item *TableItem) Tag() string {
	return "tbl"
}

// PlainText - текст
func (item *TableItem) PlainText() string {
	return ""
}

// Type - тип элемента
func (item *TableItem) Type() DocItemType {
	return Table
}

// Clone - клонирование
func (item *TableItem) Clone() DocItem {
	result := new(TableItem)
	result.Grid.Cols = make([]*WidthValue, 0)
	for _, col := range item.Grid.Cols {
		if col != nil {
			w := new(WidthValue)
			w.Type = col.Type
			w.Value = col.Value
			result.Grid.Cols = append(result.Grid.Cols, w)
		}
	}
	if item.Params.DocGrid != nil {
		result.Params.DocGrid = new(IntValue)
		result.Params.DocGrid.Value = item.Params.DocGrid.Value
	}
	if item.Params.Ind != nil {
		result.Params.Ind = new(WidthValue)
		result.Params.Ind.Type = item.Params.Ind.Type
		result.Params.Ind.Value = item.Params.Ind.Value
	}
	if item.Params.Jc != nil {
		result.Params.Jc = new(StringValue)
		result.Params.Jc.Value = item.Params.Jc.Value
	}
	if item.Params.Layout != nil {
		result.Params.Layout = new(TableLayout)
		result.Params.Layout.Type = item.Params.Layout.Type
	}
	if item.Params.Shadow != nil {
		result.Params.Shadow = new(ShadowValue)
		result.Params.Shadow.From(item.Params.Shadow)
	}
	if item.Params.Width != nil {
		result.Params.Width = new(WidthValue)
		result.Params.Width.From(item.Params.Width)
	}
	if item.Params.Borders != nil {
		result.Params.Borders = new(TableBorders)
		result.Params.Borders.From(item.Params.Borders)
	}
	if item.Params.Style != nil {
		result.Params.Style = new(StyleValue)
		result.Params.Style.From(item.Params.Style)
	}
	if item.Params.Look != nil {
		result.Params.Look = new(LookValue)
		result.Params.Look.From(item.Params.Look)
	}
	// Клонирование строк
	result.Rows = make([]*TableRow, 0)
	for _, row := range item.Rows {
		if row != nil {
			result.Rows = append(result.Rows, row.Clone())
		}
	}
	return result
}

// TableParams - Params table
type TableParams struct {
	Style   *StyleValue   `xml:"tblStyle,omitempty"`
	Width   *WidthValue   `xml:"tblW,omitempty"`
	Jc      *StringValue  `xml:"jc,omitempty"`
	Ind     *WidthValue   `xml:"tblInd,omitempty"`
	Borders *TableBorders `xml:"tblBorders,omitempty"`
	Shadow  *ShadowValue  `xml:"shd,omitempty"`
	Layout  *TableLayout  `xml:"tblLayout,omitempty"`
	DocGrid *IntValue     `xml:"docGrid,omitempty"`
	Look    *LookValue    `xml:"tblLook,omitempty"`
}

type WTableParams struct {
	Style   *WStyleValue   `xml:"w:tblStyle,omitempty"`
	Width   *WWidthValue   `xml:"w:tblW,omitempty"`
	Jc      *WStringValue  `xml:"w:jc,omitempty"`
	Ind     *WWidthValue   `xml:"w:tblInd,omitempty"`
	Borders *WTableBorders `xml:"w:tblBorders,omitempty"`
	Shadow  *WShadowValue  `xml:"w:shd,omitempty"`
	Layout  *WTableLayout  `xml:"w:tblLayout,omitempty"`
	DocGrid *WIntValue     `xml:"w:docGrid,omitempty"`
	Look    *WLookValue    `xml:"w:tblLook,omitempty"`
}

func (tp *TableParams) ToWTableParams() *WTableParams {
	wtp := WTableParams{}
	if tp.Style != nil {
		wtp.Style = (*WStyleValue)(tp.Style)
	}
	if tp.Width != nil {
		wtp.Width = (*WWidthValue)(tp.Width)
	}
	if tp.Jc != nil {
		wtp.Jc = (*WStringValue)(tp.Jc)
	}
	if tp.Ind != nil {
		wtp.Ind = (*WWidthValue)(tp.Ind)
	}
	if tp.Borders != nil {
		wtp.Borders = tp.Borders.ToWTableBorders()
	}
	if tp.Shadow != nil {
		wtp.Shadow = (*WShadowValue)(tp.Shadow)
	}
	if tp.Layout != nil {
		wtp.Layout = (*WTableLayout)(tp.Layout)
	}
	if tp.DocGrid != nil {
		wtp.DocGrid = (*WIntValue)(tp.DocGrid)
	}
	if tp.Look != nil {
		wtp.Look = (*WLookValue)(tp.Look)
	}
	return &wtp
}

// TableLayout - layout params
type TableLayout struct {
	Type string `xml:"type,attr"`
}
type WTableLayout struct {
	Type string `xml:"w:type,attr"`
}

// TableBorders in table
type TableBorders struct {
	Top     TableBorder  `xml:"top"`
	Left    TableBorder  `xml:"left"`
	Bottom  TableBorder  `xml:"bottom"`
	Right   TableBorder  `xml:"right"`
	InsideH *TableBorder `xml:"insideH,omitempty"`
	InsideV *TableBorder `xml:"insideV,omitempty"`
}

func (tb *TableBorders) ToWTableBorders() *WTableBorders {
	wtb := WTableBorders{Top: WTableBorder(tb.Top),
		Left:   WTableBorder(tb.Left),
		Bottom: WTableBorder(tb.Bottom),
		Right:  WTableBorder(tb.Right)}
	if tb.InsideH != nil {
		wtb.InsideH = (*WTableBorder)(tb.InsideH)
	}
	if tb.InsideV != nil {
		wtb.InsideV = (*WTableBorder)(tb.InsideV)
	}
	return &wtb
}

type WTableBorders struct {
	Top     WTableBorder  `xml:"w:top"`
	Left    WTableBorder  `xml:"w:left"`
	Bottom  WTableBorder  `xml:"w:bottom"`
	Right   WTableBorder  `xml:"w:right"`
	InsideH *WTableBorder `xml:"w:insideH,omitempty"`
	InsideV *WTableBorder `xml:"w:insideV,omitempty"`
}

// From (TableBorders)
func (b *TableBorders) From(b1 *TableBorders) {
	if b1 != nil {
		b.Top.From(&b1.Top)
		b.Left.From(&b1.Left)
		b.Bottom.From(&b1.Bottom)
		b.Right.From(&b1.Right)
		if b1.InsideH != nil {
			b.InsideH = new(TableBorder)
			b.InsideH.From(b1.InsideH)
		}
		if b1.InsideV != nil {
			b.InsideV = new(TableBorder)
			b.InsideV.From(b1.InsideV)
		}
	}
}

// TableBorder in borders
type TableBorder struct {
	Value  string `xml:"val,attr"`
	Color  string `xml:"color,attr"`
	Size   int64  `xml:"sz,attr"`
	Space  int64  `xml:"space,attr"`
	Shadow int64  `xml:"shadow,attr"`
	Frame  int64  `xml:"frame,attr"`
}

type WTableBorder struct {
	Value  string `xml:"w:val,attr"`
	Color  string `xml:"w:color,attr"`
	Size   int64  `xml:"w:sz,attr"`
	Space  int64  `xml:"w:space,attr"`
	Shadow int64  `xml:"w:shadow,attr"`
	Frame  int64  `xml:"w:frame,attr"`
}

// From (TableBorder)
func (b *TableBorder) From(b1 *TableBorder) {
	if b1 != nil {
		b.Value = b1.Value
		b.Color = b1.Color
		b.Frame = b1.Frame
		b.Shadow = b1.Shadow
		b.Size = b1.Size
		b.Space = b1.Space
	}
}

// TableRow - row in table
type TableRow struct {
	OtherParams *TableParamsEx `xml:"tblPrEx,omitempty"`
	Params      TableRowParams `xml:"trPr"`
	Cells       []*TableCell   `xml:"tc,omitempty"`
	RsidR       string         `xml:"rsidR,attr,omitempty"`
	RsidTr      string         `xml:"rsidTr,attr,omitempty"`
}

// TableRowParams - row params
type TableRowParams struct {
	Height   HeightValue `xml:"trHeight"`
	IsHeader bool
}

// TableCell - table cell
type TableCell struct {
	Params TableCellParams `xml:"tcPr"`
	Items  []DocItem
}

// TableCellParams - cell params
type TableCellParams struct {
	Width         *WidthValue   `xml:"tcW,omitempty"`
	Borders       *TableBorders `xml:"tcBorders,omitempty"`
	Shadow        *ShadowValue  `xml:"shd,omitempty"`
	Margins       *Margins      `xml:"tcMar,omitempty"`
	VerticalAlign *StringValue  `xml:"vAlign,omitempty"`
	VerticalMerge *StringValue  `xml:"vMerge,omitempty"`
	GridSpan      *IntValue     `xml:"gridSpan,omitempty"`
	HideMark      *EmptyValue   `xml:"hideMark,omitempty"`
	NoWrap        *EmptyValue   `xml:"noWrap,omitempty"`
}

func (tcp *TableCellParams) toWTableCellParams() *WTableCellParams {
	wtcp := WTableCellParams{}
	// Width         *WidthValue   `xml:"tcW,omitempty"`
	if tcp.Width != nil {
		wtcp.Width = (*WWidthValue)(tcp.Width)
	}
	// Borders       *TableBorders `xml:"tcBorders,omitempty"`
	if tcp.Borders != nil {
		wtcp.Borders = tcp.Borders.ToWTableBorders()
	}
	// Shadow        *ShadowValue  `xml:"shd,omitempty"`
	if tcp.Shadow != nil {
		wtcp.Shadow = (*WShadowValue)(tcp.Shadow)
	}
	// Margins       *Margins      `xml:"tcMar,omitempty"`
	if tcp.Margins != nil {
		wtcp.Margins = tcp.Margins.ToWMargins()
	}
	// VerticalAlign *StringValue  `xml:"vAlign,omitempty"`
	if tcp.VerticalAlign != nil {
		wtcp.VerticalAlign = (*WStringValue)(tcp.VerticalAlign)
	}
	// VerticalMerge *StringValue  `xml:"vMerge,omitempty"`
	if tcp.GridSpan != nil {
		wtcp.GridSpan = (*WIntValue)(tcp.GridSpan)
	}
	// GridSpan      *IntValue     `xml:"gridSpan,omitempty"`
	if tcp.HideMark != nil {
		wtcp.HideMark = (*WEmptyValue)(tcp.HideMark)
	}
	// HideMark      *EmptyValue   `xml:"hideMark,omitempty"`
	if tcp.NoWrap != nil {
		wtcp.NoWrap = (*WEmptyValue)(tcp.NoWrap)
	}
	// NoWrap        *EmptyValue   `xml:"noWrap,omitempty"`
	return &wtcp
}

type WTableCellParams struct {
	Width         *WWidthValue   `xml:"w:tcW,omitempty"`
	Borders       *WTableBorders `xml:"w:tcBorders,omitempty"`
	Shadow        *WShadowValue  `xml:"w:shd,omitempty"`
	Margins       *WMargins      `xml:"w:tcMar,omitempty"`
	VerticalAlign *WStringValue  `xml:"w:vAlign,omitempty"`
	VerticalMerge *WStringValue  `xml:"w:vMerge,omitempty"`
	GridSpan      *WIntValue     `xml:"w:gridSpan,omitempty"`
	HideMark      *WEmptyValue   `xml:"w:hideMark,omitempty"`
	NoWrap        *WEmptyValue   `xml:"w:noWrap,omitempty"`
}

// Clone (TableCell) - клонирование ячейки
func (cell *TableCell) Clone() *TableCell {
	result := new(TableCell)
	if cell.Params.GridSpan != nil {
		result.Params.GridSpan = new(IntValue)
		result.Params.GridSpan.Value = cell.Params.GridSpan.Value
	}
	if cell.Params.HideMark != nil {
		result.Params.HideMark = new(EmptyValue)
	}
	if cell.Params.NoWrap != nil {
		result.Params.NoWrap = new(EmptyValue)
	}
	if cell.Params.Shadow != nil {
		result.Params.Shadow = new(ShadowValue)
		result.Params.Shadow.From(cell.Params.Shadow)
	}
	if cell.Params.VerticalAlign != nil {
		result.Params.VerticalAlign = new(StringValue)
		result.Params.VerticalAlign.Value = cell.Params.VerticalAlign.Value
	}
	if cell.Params.VerticalMerge != nil {
		result.Params.VerticalMerge = new(StringValue)
		result.Params.VerticalMerge.Value = cell.Params.VerticalMerge.Value
	}
	if cell.Params.Margins != nil {
		result.Params.Margins = new(Margins)
		result.Params.Margins.From(cell.Params.Margins)
	}
	if cell.Params.Width != nil {
		result.Params.Width = new(WidthValue)
		result.Params.Width.From(cell.Params.Width)
	}
	if cell.Params.Borders != nil {
		result.Params.Borders = new(TableBorders)
		result.Params.Borders.From(cell.Params.Borders)
	}
	result.Items = make([]DocItem, 0)
	for _, item := range cell.Items {
		if item != nil {
			result.Items = append(result.Items, item.Clone())
		}
	}
	return result
}

// Clone (TableRow) - клонирование строки таблицы
func (row *TableRow) Clone() *TableRow {
	result := new(TableRow)
	result.Params = row.Params
	result.OtherParams = new(TableParamsEx)
	result.OtherParams.Shadow = row.OtherParams.Shadow
	// Клонируем ячейки
	result.Cells = make([]*TableCell, 0)
	for _, cell := range row.Cells {
		if cell != nil {
			result.Cells = append(result.Cells, cell.Clone())
		}
	}
	result.RsidR = row.RsidR
	result.RsidTr = row.RsidTr
	return result
}

/* ДЕКОДИРОВАНИЕ */

// Декодирование таблицы
func (item *TableItem) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		item.Rows = make([]*TableRow, 0)
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "tblPr" {
						decoder.DecodeElement(&item.Params, &element)
					} else if element.Name.Local == "tblGrid" {
						decoder.DecodeElement(&item.Grid, &element)
					} else if element.Name.Local == "tr" {
						row := new(TableRow)
						for _, attr := range element.Attr {
							if attr.Name.Local == "rsidR" {
								row.RsidR = attr.Value
							}
							if attr.Name.Local == "rsidTr" {
								row.RsidTr = attr.Value
							}
						}
						if row.decode(decoder) == nil {
							item.Rows = append(item.Rows, row)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "tbl" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

// Декодирование строк таблицы
func (row *TableRow) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		row.Cells = make([]*TableCell, 0)
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}

			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "trHeight" {
						decoder.DecodeElement(&row.Params.Height, &element)
					} else if element.Name.Local == "tblHeader" {
						row.Params.IsHeader = true
					} else if element.Name.Local == "tblPrEx" {
						row.OtherParams = new(TableParamsEx)
						decoder.DecodeElement(row.OtherParams, &element)
					} else if element.Name.Local == "tc" {
						cell := new(TableCell)
						if cell.decode(decoder) == nil {
							row.Cells = append(row.Cells, cell)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "tr" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

// Декодирование ячеек таблицы
func (row *TableCell) decode(decoder *xml.Decoder) error {
	if decoder != nil {
		var end bool
		for !end {
			token, _ := decoder.Token()
			if token == nil {
				break
			}
			switch element := token.(type) {
			case xml.StartElement:
				{
					if element.Name.Local == "tcPr" {
						decoder.DecodeElement(&row.Params, &element)
					} else {
						i := decodeItem(&element, decoder)
						if i != nil {
							row.Items = append(row.Items, i)
						}
					}
				}
			case xml.EndElement:
				{
					if element.Name.Local == "tc" {
						end = true
					}
				}
			}
		}
		return nil
	}
	return errors.New("Not have decoder")
}

/* КОДИРОВАНИЕ */

// Кодирование таблицы
func (item *TableItem) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало таблицы
		start := xml.StartElement{Name: xml.Name{Local: "w:" + item.Tag()}}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры таблицы
		if err := encoder.EncodeElement(item.Params.ToWTableParams(), xml.StartElement{Name: xml.Name{Local: "w:" + "tblPr"}}); err != nil {
			return err
		}
		// Сетка таблицы
		if err := encoder.EncodeElement(item.Grid.ToWGirdParams(), xml.StartElement{Name: xml.Name{Local: "w:" + "tblGrid"}}); err != nil {
			return err
		}
		// Строки таблицы
		for _, row := range item.Rows {
			if row != nil {
				if err := row.encode(encoder); err != nil {
					return err
				}
			}
		}
		// Конец таблицы
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}

// Кодирование ячейки таблицы
func (cell *TableCell) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало ячейки таблицы
		start := xml.StartElement{Name: xml.Name{Local: "w:" + "tc"}}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры ячейки таблицы
		if err := encoder.EncodeElement(cell.Params.toWTableCellParams(), xml.StartElement{Name: xml.Name{Local: "w:" + "tcPr"}}); err != nil {
			return err
		}
		// Кодируем составные элементы
		for _, i := range cell.Items {
			if err := i.encode(encoder); err != nil {
				return err
			}
		}
		// Конец ячейки таблицы
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}

// Кодирование строки таблицы
func (row *TableRow) encode(encoder *xml.Encoder) error {
	if encoder != nil {
		// Начало строки таблицы
		var rsidR = xml.Attr{Name: xml.Name{Local: "w:" + "rsidR"}, Value: row.RsidR}
		var rsidTr = xml.Attr{Name: xml.Name{Local: "w:" + "rsidTr"}, Value: row.RsidTr}
		var attrs = []xml.Attr{rsidR, rsidTr}
		start := xml.StartElement{Name: xml.Name{Local: "w:" + "tr"}, Attr: attrs}
		if err := encoder.EncodeToken(start); err != nil {
			return err
		}
		// Параметры строки таблицы
		if row.OtherParams != nil {
			if err := encoder.EncodeElement((*WTableParamsEx)(row.OtherParams), xml.StartElement{Name: xml.Name{Local: "w:" + "tblPrEx"}}); err != nil {
				return err
			}
		}
		// Кодируем Параметры
		startPr := xml.StartElement{Name: xml.Name{Local: "w:" + "trPr"}}
		if err := encoder.EncodeToken(startPr); err != nil {
			return err
		}
		if err := encoder.EncodeElement(&row.Params.Height, xml.StartElement{Name: xml.Name{Local: "w:" + "trHeight"}}); err != nil {
			return err
		}
		if row.Params.IsHeader {
			startHeader := xml.StartElement{Name: xml.Name{Local: "w:" + "tblHeader"}}
			if err := encoder.EncodeToken(startHeader); err != nil {
				return err
			}
			if err := encoder.EncodeToken(startHeader.End()); err != nil {
				return err
			}
			if err := encoder.Flush(); err != nil {
				return err
			}
		}
		if err := encoder.EncodeToken(startPr.End()); err != nil {
			return err
		}
		if err := encoder.Flush(); err != nil {
			return err
		}
		// Кодируем ячейки
		for _, cell := range row.Cells {
			if cell != nil {
				if err := cell.encode(encoder); err != nil {
					return err
				}
			}
		}
		// Конец строки таблицы
		if err := encoder.EncodeToken(start.End()); err != nil {
			return err
		}
		return encoder.Flush()
	}
	return errors.New("Not have encoder")
}
