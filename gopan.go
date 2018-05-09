package gopan // import "go.vktec.org.uk/gopan"

/*
#cgo pkg-config: pango
#include <stdlib.h>
#include <pango/pango.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
	"image/color"
)

const Scale = C.PANGO_SCALE

type Layout struct{ L *C.PangoLayout }

func CreateLayoutFromC(cl unsafe.Pointer) Layout {
	l := Layout{(*C.PangoLayout)(cl)}
	runtime.SetFinalizer(&l, func(l *Layout) {
		C.g_object_unref(l.L)
	})
	return l
}

func (l Layout) Attributes() AttrList {
	p := C.pango_layout_get_attributes(l.L)
	C.pango_attr_list_ref(p)
	return newAttrListFromC(p)
}

func (l Layout) SetAttributes(attrs AttrList) {
	C.pango_layout_set_attributes(l.L, attrs.l)
}

func (l Layout) SetFontDescription(d FontDescription) {
	C.pango_layout_set_font_description(l.L, d.d)
}

func (l Layout) SetText(text string) {
	ctext := C.CString(text)
	C.pango_layout_set_text(l.L, ctext, C.int(len(text)))
	C.free(unsafe.Pointer(ctext))
}

func (l Layout) PixelSize() (int, int) {
	var w, h C.int
	C.pango_layout_get_pixel_size(l.L, &w, &h)
	return int(w), int(h)
}

type WrapMode C.PangoWrapMode

const (
	Word WrapMode = C.PANGO_WRAP_WORD
	Char WrapMode = C.PANGO_WRAP_CHAR
	WordChar WrapMode = C.PANGO_WRAP_WORD_CHAR
)

func (l Layout) SetWrap(wrap WrapMode) {
	C.pango_layout_set_wrap(l.L, C.PangoWrapMode(wrap))
}

func (l Layout) Width() int {
	return int(C.pango_layout_get_width(l.L))
}

func (l Layout) SetWidth(w int) {
	C.pango_layout_set_width(l.L, C.int(w))
}

func (l Layout) Height() int {
	return int(C.pango_layout_get_height(l.L))
}

func (l Layout) SetHeight(w int) {
	C.pango_layout_set_height(l.L, C.int(w))
}

type Context struct{ c *C.PangoContext }

func (l Layout) Context() Context {
	cc := C.pango_layout_get_context(l.L)
	return Context{cc}
}

type FontDescription struct{ d *C.PangoFontDescription }

func FontDescriptionFromString(desc string) FontDescription {
	cdesc := C.CString(desc)
	cd := C.pango_font_description_from_string(cdesc)
	C.free(unsafe.Pointer(cdesc))
	d := FontDescription{cd}
	runtime.SetFinalizer(&d, func(d *FontDescription) {
		C.pango_font_description_free(d.d)
	})
	return d
}

type FontMap struct{ m *C.PangoFontMap }

func CreateFontMapFromC(cm unsafe.Pointer) FontMap {
	return FontMap{(*C.PangoFontMap)(cm)}
}

func (m FontMap) LoadFont(c Context, d FontDescription) Font {
	cf := C.pango_font_map_load_font(m.m, c.c, d.d)
	f := Font{cf}
	runtime.SetFinalizer(&f, func(f *Font) {
		C.g_object_unref(f.f)
	})
	return f
}

type Font struct{ f *C.PangoFont }

func (f Font) Metrics() FontMetrics {
	cm := C.pango_font_get_metrics(f.f, nil) // TODO: support fontsets
	m := FontMetrics{cm}
	runtime.SetFinalizer(&m, func(m *FontMetrics) {
		C.pango_font_metrics_unref(m.m)
	})
	return m
}

type FontMetrics struct{ m *C.PangoFontMetrics }

func (m FontMetrics) Ascent() int {
	return int(C.pango_font_metrics_get_ascent(m.m))
}

func (m FontMetrics) Descent() int {
	return int(C.pango_font_metrics_get_descent(m.m))
}

func (m FontMetrics) ApproximateCharWidth() int {
	return int(C.pango_font_metrics_get_approximate_char_width(m.m))
}

func (m FontMetrics) ApproximateDigitWidth() int {
	return int(C.pango_font_metrics_get_approximate_digit_width(m.m))
}

type Attribute struct{ a *C.PangoAttribute }

type Style C.PangoStyle

const (
	NormalStyle Style = C.PANGO_STYLE_NORMAL
	Oblique     Style = C.PANGO_STYLE_OBLIQUE
	Italic      Style = C.PANGO_STYLE_ITALIC
)

type Variant C.PangoVariant

const (
	NormalVariant Variant = C.PANGO_VARIANT_NORMAL
	SmallCaps     Variant = C.PANGO_VARIANT_SMALL_CAPS
)

type Stretch C.PangoStretch

const (
	UltraCondensed Stretch = C.PANGO_STRETCH_ULTRA_CONDENSED
	ExtraCondensed Stretch = C.PANGO_STRETCH_EXTRA_CONDENSED
	Condensed      Stretch = C.PANGO_STRETCH_CONDENSED
	SemiCondensed  Stretch = C.PANGO_STRETCH_SEMI_CONDENSED
	NormalStretch  Stretch = C.PANGO_STRETCH_NORMAL
	SemiExpanded   Stretch = C.PANGO_STRETCH_SEMI_EXPANDED
	Expanded       Stretch = C.PANGO_STRETCH_EXPANDED
	ExtraExpanded  Stretch = C.PANGO_STRETCH_EXTRA_EXPANDED
	UltraExpanded  Stretch = C.PANGO_STRETCH_ULTRA_EXPANDED
)

type Weight C.PangoWeight

const (
	Thin       Weight = C.PANGO_WEIGHT_THIN
	UltraLight Weight = C.PANGO_WEIGHT_ULTRALIGHT
	Light      Weight = C.PANGO_WEIGHT_LIGHT
	SemiLight  Weight = C.PANGO_WEIGHT_SEMILIGHT
	Book       Weight = C.PANGO_WEIGHT_BOOK
	Normal     Weight = C.PANGO_WEIGHT_NORMAL
	Medium     Weight = C.PANGO_WEIGHT_MEDIUM
	SemiBold   Weight = C.PANGO_WEIGHT_SEMIBOLD
	Bold       Weight = C.PANGO_WEIGHT_BOLD
	UltraBold  Weight = C.PANGO_WEIGHT_ULTRABOLD
	Heavy      Weight = C.PANGO_WEIGHT_HEAVY
	UltraHeavy Weight = C.PANGO_WEIGHT_ULTRAHEAVY
)

type Underline C.PangoUnderline

const (
	NoUnderline Underline = C.PANGO_UNDERLINE_NONE
	Single      Underline = C.PANGO_UNDERLINE_SINGLE
	Double      Underline = C.PANGO_UNDERLINE_DOUBLE
	Low         Underline = C.PANGO_UNDERLINE_LOW
	Error       Underline = C.PANGO_UNDERLINE_ERROR
)

const (
	IndexFromTextBeginning = C.PANGO_ATTR_INDEX_FROM_TEXT_BEGINNING
	IndexToTextEnd = C.PANGO_ATTR_INDEX_TO_TEXT_END
)

func newAttributeFromC(p *C.PangoAttribute) Attribute {
	a := Attribute{p}
	runtime.SetFinalizer(&a, func(a *Attribute) {
		if a.a != nil {
			C.pango_attribute_destroy(a.a)
		}
	})
	return a
}

func NewFamilyAttr(family string) Attribute {
	cfamily := C.CString(family)
	defer C.free(unsafe.Pointer(cfamily))
	return newAttributeFromC(C.pango_attr_family_new(cfamily))
}

func NewStyleAttr(s Style) Attribute {
	return newAttributeFromC(C.pango_attr_style_new(C.PangoStyle(s)))
}

func NewVariantAttr(v Variant) Attribute {
	return newAttributeFromC(C.pango_attr_variant_new(C.PangoVariant(v)))
}

func NewStretchAttr(s Stretch) Attribute {
	return newAttributeFromC(C.pango_attr_stretch_new(C.PangoStretch(s)))
}

func NewWeightAttr(w Weight) Attribute {
	return newAttributeFromC(C.pango_attr_weight_new(C.PangoWeight(w)))
}

func NewSizeAttr(size int) Attribute {
	return newAttributeFromC(C.pango_attr_size_new(C.int(size)))
}

func NewAbsoluteSizeAttr(size int) Attribute {
	return newAttributeFromC(C.pango_attr_size_new_absolute(C.int(size)))
}

func NewFontDescAttr(desc FontDescription) Attribute {
	return newAttributeFromC(C.pango_attr_font_desc_new(desc.d))
}

func c2rgb(c color.Color) (r, g, b C.guint16) {
	gr, gg, gb, _ := c.RGBA()
	return C.guint16(gr), C.guint16(gg), C.guint16(gb)
}

func NewForegroundAttr(c color.Color) Attribute {
	return newAttributeFromC(C.pango_attr_foreground_new(c2rgb(c)))
}

func NewBackgroundAttr(c color.Color) Attribute {
	return newAttributeFromC(C.pango_attr_background_new(c2rgb(c)))
}

func NewStrikethroughAttr(strike bool) Attribute {
	istrike := 0
	if strike {
		istrike = 1
	}
	return newAttributeFromC(C.pango_attr_strikethrough_new(C.gboolean(istrike)))
}

func NewStrikethroughColorAttr(c color.Color) Attribute {
	return newAttributeFromC(C.pango_attr_strikethrough_color_new(c2rgb(c)))
}

func NewUnderlineAttr(uline Underline) Attribute {
	return newAttributeFromC(C.pango_attr_underline_new(C.PangoUnderline(uline)))
}

func NewUnderlineColorAttr(c color.Color) Attribute {
	return newAttributeFromC(C.pango_attr_underline_color_new(c2rgb(c)))
}

func (a Attribute) Copy() Attribute {
	return newAttributeFromC(C.pango_attribute_copy(a.a))
}

func (a Attribute) SetStart(start int) {
	a.a.start_index = C.guint(start)
}

func (a Attribute) SetEnd(start int) {
	a.a.end_index = C.guint(start)
}

type AttrList struct{ l *C.PangoAttrList }

func newAttrListFromC(p *C.PangoAttrList) AttrList {
	l := AttrList{p}
	runtime.SetFinalizer(&l, func(l *AttrList) {
		C.pango_attr_list_unref(l.l)
	})
	return l
}

func NewAttrList() AttrList {
	return newAttrListFromC(C.pango_attr_list_new())
}

func (l AttrList) IsNil() bool {
	return l.l == nil
}

func (l AttrList) Insert(a Attribute) {
	C.pango_attr_list_insert(l.l, C.pango_attribute_copy(a.a))
}

// Slightly more efficient than Insert, but `a` cannot be used afterwards
func (l AttrList) InsertInvalidate(a *Attribute) {
	C.pango_attr_list_insert(l.l, a.a)
	a.a = nil // So it's not destroyed when the finalizer runs
}
