package docx

// HeightValue - значение высоты
type HeightValue struct {
	Value      int64  `xml:"val,attr"`
	HeightRule string `xml:"hRule,attr,omitempty"`
}

// From (HeightValue)
func (h *HeightValue) From(h1 *HeightValue) {
	if h1 != nil {
		h.HeightRule = h1.HeightRule
		h.Value = h1.Value
	}
}

// WidthValue - значение длины
type WidthValue struct {
	Value int64  `xml:"w,attr"`
	Type  string `xml:"type,attr,omitempty"`
}

// From (WidthValue)
func (w *WidthValue) From(w1 *WidthValue) {
	if w1 != nil {
		w.Type = w1.Type
		w.Value = w1.Value
	}
}

// SizeValue - значение размера
type SizeValue struct {
	Width       int64  `xml:"w,attr"`
	Height      int64  `xml:"h,attr"`
	Orientation string `xml:"orient,attr,omitempty"`
}

// From (SizeValue)
func (s *SizeValue) From(s1 *SizeValue) {
	if s1 != nil {
		s.Height = s1.Height
		s.Orientation = s1.Orientation
		s.Width = s1.Width
	}
}

// EmptyValue - пустое значение
type EmptyValue struct {
}

// StringValue - одиночное string значение
type StringValue struct {
	Value string `xml:"val,attr,omitempty"`
}

// From (StringValue)
func (s *StringValue) From(s1 *StringValue) {
	if s1 != nil {
		s.Value = s1.Value
	}
}

// BoolValue - одиночное bool значение
type BoolValue struct {
	Value bool `xml:"val,attr"`
}

// IntValue - одиночное int значение
type IntValue struct {
	Value int64 `xml:"val,attr"`
}

// From (IntValue)
func (i *IntValue) From(i1 *IntValue) {
	if i1 != nil {
		i.Value = i1.Value
	}
}

// FloatValue - одиночное float значение
type FloatValue struct {
	Value float64 `xml:"val,attr"`
}

// ReferenceValue - reference value
type ReferenceValue struct {
	Type string `xml:"type,attr"`
	ID   string `xml:"id,attr"`
}

// SpacingValue - spacing value
type SpacingValue struct {
	After    int64  `xml:"after,attr"`
	Before   int64  `xml:"before,attr"`
	Line     int64  `xml:"line,attr"`
	LineRule string `xml:"lineRule,attr"`
}

// From (SpacingValue)
func (s *SpacingValue) From(s1 *SpacingValue) {
	if s1 != nil {
		s.After = s1.After
		s.Before = s1.Before
		s.Line = s1.Line
		s.LineRule = s1.LineRule
	}
}

// MarginValue - margin значение
type MarginValue struct {
	Top    int64 `xml:"top,attr"`
	Left   int64 `xml:"left,attr"`
	Bottom int64 `xml:"bottom,attr"`
	Right  int64 `xml:"right,attr"`
	Header int64 `xml:"header,attr,omitempty"`
	Footer int64 `xml:"footer,attr,omitempty"`
}

// From (MarginValue)
func (m *MarginValue) From(m1 *MarginValue) {
	if m1 != nil {
		m.Top = m1.Top
		m.Left = m1.Left
		m.Bottom = m1.Bottom
		m.Right = m1.Right
		m.Header = m1.Header
		m.Footer = m1.Footer
	}
}

// Margins - margins значение
type Margins struct {
	Top    WidthValue `xml:"top"`
	Left   WidthValue `xml:"left"`
	Bottom WidthValue `xml:"bottom"`
	Right  WidthValue `xml:"right"`
}

// From (Margins)
func (m *Margins) From(m1 *Margins) {
	if m1 != nil {
		m.Top.From(&m1.Top)
		m.Left.From(&m1.Left)
		m.Bottom.From(&m1.Bottom)
		m.Right.From(&m1.Right)
	}
}

// ShadowValue - значение тени
type ShadowValue struct {
	Value string `xml:"val,attr"`
	Color string `xml:"color,attr"`
	Fill  string `xml:"fill,attr"`
}

// From (ShadowValue)
func (s *ShadowValue) From(s1 *ShadowValue) {
	if s1 != nil {
		s.Value = s1.Value
		s.Color = s1.Color
		s.Fill = s1.Fill
	}
}

type StyleValue struct {
	Value string `xml:"val,attr,omitempty"`
}

func (s *StyleValue) From(s1 *StyleValue) {
	if s1 != nil {
		s.Value = s1.Value
	}
}

type LookValue struct {
	Value       string `xml:"val,attr,omitempty"`
	FirstRow    string `xml:"firstRow,attr,omitempty"`
	LastRow     string `xml:"lastRow,attr,omitempty"`
	FirstColumn string `xml:"firstColumn,attr,omitempty"`
	LastColumn  string `xml:"lastColumn,attr,omitempty"`
	NoHBand     string `xml:"noHBand,attr,omitempty"`
	NoVBand     string `xml:"noVBand,attr,omitempty"`
}

func (l *LookValue) From(l1 *LookValue) {
	if l1 != nil {
		l.FirstColumn = l1.FirstColumn
		l.FirstRow = l1.FirstRow
		l.LastColumn = l1.LastColumn
		l.LastRow = l1.LastRow
		l.NoHBand = l1.NoHBand
		l.NoVBand = l1.NoVBand
		l.Value = l1.Value
	}
}

type PBdrValue struct {
	Top     BdrValue `xml:"top,omitempty"`
	Left    BdrValue `xml:"left,omitempty"`
	Bottom  BdrValue `xml:"bottom,omitempty"`
	Right   BdrValue `xml:"right,omitempty"`
	Between BdrValue `xml:"between,omitempty"`
	Bar     BdrValue `xml:"bar,omitempty"`
}

func (pb *PBdrValue) From(pb1 *PBdrValue) {
	if pb1 != nil {
		pb.Top.From(&pb1.Top)
		pb.Left.From(&pb1.Left)
		pb.Bottom.From(&pb1.Bottom)
		pb.Right.From(&pb1.Right)
		pb.Between.From(&pb1.Between)
	}
}

type BdrValue struct {
	Value string `xml:"val,attr,omitempty"`
	Sz    string `xml:"sz,attr,omitempty"`
	Space string `xml:"space,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
}

func (b *BdrValue) From(b1 *BdrValue) {
	if b1 != nil {
		b.Value = b1.Value
		b.Sz = b1.Sz
		b.Space = b1.Space
		b.Color = b1.Color
	}
}
