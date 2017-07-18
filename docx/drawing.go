package docx

type Drawing struct {
	Inline *Inline `xml:"inline,omitempty"`
	Anchor *Anchor `xml:"anchor,omitempty"`
}

type WDrawing struct {
	Inline *WInline `xml:"wp:inline,omitempty"`
	Anchor *WAnchor `xml:"wp:anchor,omitempty"`
}

func (d *Drawing) ToWDrawing() *WDrawing {
	wd := WDrawing{}
	if d.Inline != nil {
		wd.Inline = d.Inline.ToWInline()
	}
	if d.Anchor != nil {
		wd.Anchor = d.Anchor.ToWAnchor()
	}
	return &wd
}

type Position struct {
	RelativeFrom string `xml:"relativeFrom,attr,omitempty"`
	PosOffset    string `xml:"posOffset,omitempty"`
}

type WPosition struct {
	RelativeFrom string `xml:"relativeFrom,attr,omitempty"`
	PosOffset    string `xml:"wp:posOffset,omitempty"`
}

type SizeRel struct {
	RelativeFrom string `xml:"relativeFrom,attr,omitempty"`
	PctWidth     string `xml:"pctWidth,omitempty"`
	PctHeight    string `xml:"pctHeight,omitempty"`
}

type WSizeRel struct {
	RelativeFrom string `xml:"relativeFrom,attr,omitempty"`
	PctWidth     string `xml:"wp14:pctWidth,omitempty"`
	PctHeight    string `xml:"wp14:pctHeight,omitempty"`
}

type Anchor struct {
	DistT             string             `xml:"distT,attr,omitempty"`
	DistB             string             `xml:"distB,attr,omitempty"`
	DistL             string             `xml:"distL,attr,omitempty"`
	DistR             string             `xml:"distR,attr,omitempty"`
	SimplePosAttr     string             `xml:"simplePos,attr,omitempty"`
	SimplePos         *XyValue           `xml:"simplePos,omitempty"`
	RelativeHeight    string             `xml:"relativeHeight,attr,omitempty"`
	BehindDoc         string             `xml:"behindDoc,attr,omitempty"`
	Locked            string             `xml:"locked,attr,omitempty"`
	LayoutInCell      string             `xml:"layoutInCell,attr,omitempty"`
	AllowOverlap      string             `xml:"allowOverlap,attr,omitempty"`
	PositionH         *Position          `xml:"positionH,omitempty"`
	PositionV         *Position          `xml:"positionV,omitempty"`
	Extent            *CxCyValue         `xml:"extent,omitempty"`
	EffectExtent      *LtrbValue         `xml:"effectExtent,omitempty"`
	WrapNone          *EmptyValue        `xml:"wrapNone,omitempty"`
	DocPr             *IdNameValue       `xml:"docPr,omitempty"`
	CNvGraphicFramePr *CNvGraphicFramePr `xml:"cNvGraphicFramePr,omitempty"`
	Graphic           *Graphic           `xml:"graphic,omitempty"`
	SizeRelH          *SizeRel           `xml:"sizeRelH,omitempty"`
	SizeRelV          *SizeRel           `xml:"sizeRelV,omitempty"`
}

type WAnchor struct {
	DistT             string              `xml:"distT,attr,omitempty"`
	DistB             string              `xml:"distB,attr,omitempty"`
	DistL             string              `xml:"distL,attr,omitempty"`
	DistR             string              `xml:"distR,attr,omitempty"`
	SimplePosAttr     string              `xml:"simplePos,attr,omitempty"`
	RelativeHeight    string              `xml:"relativeHeight,attr,omitempty"`
	BehindDoc         string              `xml:"behindDoc,attr,omitempty"`
	Locked            string              `xml:"locked,attr,omitempty"`
	LayoutInCell      string              `xml:"layoutInCell,attr,omitempty"`
	AllowOverlap      string              `xml:"allowOverlap,attr,omitempty"`
	SimplePos         *XyValue            `xml:"wp:simplePos,omitempty"`
	PositionH         *WPosition          `xml:"wp:positionH,omitempty"`
	PositionV         *WPosition          `xml:"wp:positionV,omitempty"`
	Extent            *CxCyValue          `xml:"wp:extent,omitempty"`
	EffectExtent      *LtrbValue          `xml:"wp:effectExtent,omitempty"`
	WrapNone          *EmptyValue         `xml:"wp:wrapNone,omitempty"`
	DocPr             *IdNameValue        `xml:"wp:docPr,omitempty"`
	CNvGraphicFramePr *ACNvGraphicFramePr `xml:"wp:cNvGraphicFramePr,omitempty"`
	Graphic           *AGraphic           `xml:"a:graphic,omitempty"`
	SizeRelH          *WSizeRel           `xml:"wp14:sizeRelH,omitempty"`
	SizeRelV          *WSizeRel           `xml:"wp14:sizeRelV,omitempty"`
}

func (a *Anchor) ToWAnchor() *WAnchor {
	wa := WAnchor{DistT: a.DistT, DistB: a.DistB, DistL: a.DistL, DistR: a.DistR,
		SimplePosAttr:  a.SimplePosAttr,
		RelativeHeight: a.RelativeHeight,
		BehindDoc:      a.BehindDoc,
		Locked:         a.Locked,
		LayoutInCell:   a.LayoutInCell,
		AllowOverlap:   a.AllowOverlap,
		SimplePos:      a.SimplePos,
		WrapNone:       a.WrapNone,
		Extent:         a.Extent,
		EffectExtent:   a.EffectExtent,
		DocPr:          a.DocPr}
	if a.PositionH != nil {
		wa.PositionH = (*WPosition)(a.PositionH)
	}
	if a.PositionV != nil {
		wa.PositionV = (*WPosition)(a.PositionV)
	}
	if a.CNvGraphicFramePr != nil {
		wa.CNvGraphicFramePr = a.CNvGraphicFramePr.ToACNvGraphicFramePr()
	}
	if a.Graphic != nil {
		wa.Graphic = a.Graphic.ToAGraphic()
	}
	if a.SizeRelH != nil {
		wa.SizeRelH = (*WSizeRel)(a.SizeRelH)
	}
	if a.SizeRelH != nil {
		wa.SizeRelV = (*WSizeRel)(a.SizeRelV)
	}
	return &wa
}

type Inline struct {
	DistT             string             `xml:"distT,attr,omitempty"`
	DistB             string             `xml:"distB,attr,omitempty"`
	DistL             string             `xml:"distL,attr,omitempty"`
	DistR             string             `xml:"distR,attr,omitempty"`
	Extent            *CxCyValue         `xml:"extent,omitempty"`
	EffectExtent      *LtrbValue         `xml:"effectExtent,omitempty"`
	DocPr             *IdNameValue       `xml:"docPr,omitempty"`
	CNvGraphicFramePr *CNvGraphicFramePr `xml:"cNvGraphicFramePr,omitempty"`
	Graphic           *Graphic           `xml:"graphic,omitempty"`
}

type WInline struct {
	DistT             string              `xml:"distT,attr,omitempty"`
	DistB             string              `xml:"distB,attr,omitempty"`
	DistL             string              `xml:"distL,attr,omitempty"`
	DistR             string              `xml:"distR,attr,omitempty"`
	Extent            *CxCyValue          `xml:"wp:extent,omitempty"`
	EffectExtent      *LtrbValue          `xml:"wp:effectExtent,omitempty"`
	DocPr             *IdNameValue        `xml:"wp:docPr,omitempty"`
	CNvGraphicFramePr *ACNvGraphicFramePr `xml:"wp:cNvGraphicFramePr,omitempty"`
	Graphic           *AGraphic           `xml:"a:graphic,omitempty"`
}

func (i *Inline) ToWInline() *WInline {
	wi := WInline{DistT: i.DistT, DistB: i.DistB, DistL: i.DistL, DistR: i.DistR}
	if i.Extent != nil {
		wi.Extent = i.Extent
	}
	if i.EffectExtent != nil {
		wi.EffectExtent = i.EffectExtent
	}
	if i.DocPr != nil {
		wi.DocPr = i.DocPr
	}
	if i.CNvGraphicFramePr != nil {
		wi.CNvGraphicFramePr = i.CNvGraphicFramePr.ToACNvGraphicFramePr()
	}
	if i.Graphic != nil {
		wi.Graphic = i.Graphic.ToAGraphic()
	}
	return &wi
}

type CNvGraphicFramePr struct {
	GraphicFrameLocks *XmlnSValue `xml:"graphicFrameLocks,omitempty"`
}

type ACNvGraphicFramePr struct {
	GraphicFrameLocks *AXmlnSValue `xml:"a:graphicFrameLocks,omitempty"`
}

func (c *CNvGraphicFramePr) ToACNvGraphicFramePr() *ACNvGraphicFramePr {
	ac := ACNvGraphicFramePr{}
	if c.GraphicFrameLocks != nil {
		ac.GraphicFrameLocks = (*AXmlnSValue)(c.GraphicFrameLocks)
	}
	return &ac
}

type Graphic struct {
	A              string       `xml:"a,attr,omitempty"`
	NoChangeAspect string       `xml:"noChangeAspect,attr,omitempty"`
	GraphicData    *GraphicData `xml:"graphicData,omitempty"`
}

type AGraphic struct {
	A              string          `xml:"xmlns:a,attr,omitempty"`
	NoChangeAspect string          `xml:"noChangeAspect,attr,omitempty"`
	GraphicData    *PicGraphicData `xml:"a:graphicData,omitempty"`
}

func (g *Graphic) ToAGraphic() *AGraphic {
	ag := AGraphic{A: g.A, NoChangeAspect: g.NoChangeAspect}
	if g.GraphicData != nil {
		ag.GraphicData = g.GraphicData.ToPicGraphicData()
	}
	return &ag
}

type GraphicData struct {
	Uri string `xml:"uri,attr,omitempty"`
	Pic *Pic   `xml:"pic,omitempty"`
}

type PicGraphicData struct {
	Uri string  `xml:"uri,attr,omitempty"`
	Pic *PicPic `xml:"pic:pic,omitempty"`
}

func (g *GraphicData) ToPicGraphicData() *PicGraphicData {
	pg := PicGraphicData{Uri: g.Uri}
	if g.Pic != nil {
		pg.Pic = g.Pic.ToPicPic()
	}
	return &pg
}

type Pic struct {
	Pic      string    `xml:"pic,attr,omitempty"`
	NvPicPr  *NvPicPr  `xml:"nvPicPr,omitempty"`
	BlipFill *BlipFill `xml:"blipFill,omitempty"`
	SpPr     *SpPr     `xml:"spPr,omitempty"`
}

type PicPic struct {
	Pic      string      `xml:"xmlns:pic,attr,omitempty"`
	NvPicPr  *PicNvPicPr `xml:"pic:nvPicPr,omitempty"`
	BlipFill *ABlipFill  `xml:"pic:blipFill,omitempty"`
	SpPr     *ASpPr      `xml:"pic:spPr,omitempty"`
}

func (p *Pic) ToPicPic() *PicPic {
	pp := PicPic{Pic: p.Pic}
	if p.NvPicPr != nil {
		pp.NvPicPr = (*PicNvPicPr)(p.NvPicPr)
	}
	if p.BlipFill != nil {
		pp.BlipFill = p.BlipFill.ToABlipFill()
	}
	if p.SpPr != nil {
		pp.SpPr = p.SpPr.ToASpPr()
	}
	return &pp
}

type SpPr struct {
	Xfrm     *Xfrm     `xml:"xfrm,omitempty"`
	PrstGeom *PrstGeom `xml:"prstGeom,omitempty"`
}

type ASpPr struct {
	Xfrm     *AXfrm     `xml:"a:xfrm,omitempty"`
	PrstGeom *APrstGeom `xml:"a:prstGeom,omitempty"`
}

func (s *SpPr) ToASpPr() *ASpPr {
	as := ASpPr{}
	if s.Xfrm != nil {
		as.Xfrm = (*AXfrm)(s.Xfrm)
	}
	if s.PrstGeom != nil {
		as.PrstGeom = (*APrstGeom)(s.PrstGeom)
	}
	return &as
}

type PrstGeom struct {
	Prst  string `xml:"prst,attr,omitempty"`
	AvLst AvLst  `xml:"avLst,omitempty"`
}

type APrstGeom struct {
	Prst  string `xml:"prst,attr,omitempty"`
	AvLst AvLst  `xml:"a:avLst,omitempty"`
}

type AvLst struct {
}

type Xfrm struct {
	Rot string    `xml:"rot,attr,omitempty"`
	Off XyValue   `xml:"off,omitempty"`
	Ext CxCyValue `xml:"ext,omitempty"`
}
type AXfrm struct {
	Rot string    `xml:"rot,attr,omitempty"`
	Off XyValue   `xml:"a:off,omitempty"`
	Ext CxCyValue `xml:"a:ext,omitempty"`
}
type NvPicPr struct {
	CNvPr    *IdNameValue `xml:"cNvPr,omitempty"`
	CNvPicPr *EmptyValue  `xml:"cNvPicPr,omitempty"`
}

type PicNvPicPr struct {
	CNvPr    *IdNameValue `xml:"pic:cNvPr,omitempty"`
	CNvPicPr *EmptyValue  `xml:"pic:cNvPicPr,omitempty"`
}

type BlipFill struct {
	Blip    *Blip   `xml:"blip,omitempty"`
	Stretch Stretch `xml:"strecth,omitempty"`
}

type ABlipFill struct {
	Blip    *ABlip   `xml:"a:blip,omitempty"`
	Stretch AStretch `xml:"a:strecth,omitempty"`
}

func (b *BlipFill) ToABlipFill() *ABlipFill {
	ab := ABlipFill{}
	if b.Blip != nil {
		ab.Blip = b.Blip.ToABlip()
	}
	ab.Stretch = AStretch(b.Stretch)
	return &ab
}

type Stretch struct {
	FillRect FillRect `xml:"fillRect,omitempty"`
}

type AStretch struct {
	FillRect FillRect `xml:"a:fillRect,omitempty"`
}

type FillRect struct {
}

type Blip struct {
	Embed  string  `xml:"embed,attr,omitempty"`
	Cstate string  `xml:"cstate,attr,omitempty"`
	ExtLst *ExtLst `xml:"extLst,omitempty"`
}

type ABlip struct {
	Embed  string   `xml:"r:embed,attr,omitempty"`
	Cstate string   `xml:"cstate,attr,omitempty"`
	ExtLst *AExtLst `xml:"a:extLst,omitempty"`
}

func (b *Blip) ToABlip() *ABlip {
	ab := ABlip{Embed: b.Embed, Cstate: b.Cstate}
	if b.ExtLst != nil {
		ab.ExtLst = b.ExtLst.ToAExtLst()
	}
	return &ab
}

type ExtLst struct {
	Ext *Ext `xml:"ext,omitempty"`
}
type AExtLst struct {
	Ext *A14Ext `xml:"a:ext,omitempty"`
}

func (e *ExtLst) ToAExtLst() *AExtLst {
	ae := AExtLst{}
	if e.Ext != nil {
		ae.Ext = e.Ext.ToA14Ext()
	}
	return &ae
}

type Ext struct {
	Uri         string       `xml:"uri,attr,omitempty"`
	UseLocalDpi *UseLocalDpi `xml:"useLocalDpi,omitempty"`
}

type A14Ext struct {
	Uri         string          `xml:"uri,attr,omitempty"`
	UseLocalDpi *A14UseLocalDpi `xml:"a14:useLocalDpi,omitempty"`
}

func (e *Ext) ToA14Ext() *A14Ext {
	ae := A14Ext{Uri: e.Uri}
	if e.UseLocalDpi != nil {
		ae.UseLocalDpi = (*A14UseLocalDpi)(e.UseLocalDpi)
	}
	return &ae
}

type UseLocalDpi struct {
	A14 string `xml:"a14,attr,omitempty"`
	Val string `xml:"val,attr,omitempty"`
}

type A14UseLocalDpi struct {
	A14 string `xml:"xmlns:a14,attr,omitempty"`
	Val string `xml:"val,attr,omitempty"`
}

type CxCyValue struct {
	Cx string `xml:"cx,attr,omitempty"`
	Cy string `xml:"cy,attr,omitempty"`
}

type LtrbValue struct {
	L string `xml:"l,attr,omitempty"`
	T string `xml:"t,attr,omitempty"`
	R string `xml:"r,attr,omitempty"`
	B string `xml:"b,attr,omitempty"`
}

type IdNameValue struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type XmlnSValue struct {
	A              string `xml:"a,attr,omitempty"`
	NoChangeAspect string `xml:"noChangeAspect,attr,omitempty"`
	Pic            string `xml:"pic,attr,omitempty"`
	Uri            string `xml:"uri,attr,omitempty"`
	A14            string `xml:"a14,attr,omitempty"`
}

type AXmlnSValue struct {
	A              string `xml:"xmlns:a,attr,omitempty"`
	NoChangeAspect string `xml:"noChangeAspect,attr,omitempty"`
	Pic            string `xml:"pic,attr,omitempty"`
	Uri            string `xml:"uri,attr,omitempty"`
	A14            string `xml:"a14,attr,omitempty"`
}

type XyValue struct {
	X string `xml:"x,attr,omitempty"`
	Y string `xml:"y,attr,omitempty"`
}
