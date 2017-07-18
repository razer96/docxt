package docx

import (
	"encoding/xml"
)

type IBookMark struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type BookMarkStart struct {
	IBookMark
}
type BookMarkEnd struct {
	IBookMark
}

func (b *IBookMark) Tag() string {
	return "bookMark"
}
func (b *BookMarkEnd) Tag() string {
	return "bookMarkEnd"
}
func (b *BookMarkEnd) Clone() DocItem {
	result := new(BookMarkEnd)
	return result
}
func (b *BookMarkStart) Tag() string {
	return "bookMarkStart"
}
func (b *BookMarkStart) Clone() DocItem {
	result := new(BookMarkStart)
	return result
}
func (b *IBookMark) Type() DocItemType {
	return BookMark
}
func (b *IBookMark) PlainText() string {
	return b.Name
}
func (b *IBookMark) decode(decoder *xml.Decoder) error {
	return nil
}
func (b *BookMarkStart) encode(encoder *xml.Encoder) error {
	var attrs []xml.Attr
	attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "w:" + "id"}, Value: b.ID})
	if b.Name != "" {
		attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "w:" + "name"}, Value: b.Name})
	}
	start := xml.StartElement{Name: xml.Name{Local: "w:" + b.Tag()},
		Attr: attrs}
	if err := encoder.EncodeToken(start); err != nil {
		return err
	}

	if err := encoder.EncodeToken(start.End()); err != nil {
		return err
	}
	return encoder.Flush()
}

func (b *BookMarkEnd) encode(encoder *xml.Encoder) error {
	var attrs []xml.Attr
	attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "w:" + "id"}, Value: b.ID})
	if b.Name != "" {
		attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "w:" + "name"}, Value: b.Name})
	}
	start := xml.StartElement{Name: xml.Name{Local: "w:" + b.Tag()},
		Attr: attrs}
	if err := encoder.EncodeToken(start); err != nil {
		return err
	}

	if err := encoder.EncodeToken(start.End()); err != nil {
		return err
	}
	return encoder.Flush()
}
