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
)

const Scale = C.PANGO_SCALE

type Layout struct{ L *C.PangoLayout }

func CreateLayoutFromC(cl unsafe.Pointer) Layout {
	l := Layout{(*C.PangoLayout)(cl)}
	runtime.SetFinalizer(l, func(l Layout) {
		C.g_object_unref(l.L)
	})
	return l
}

func (l Layout) SetFontDescription(d FontDescription) {
	C.pango_layout_set_font_description(l.L, d.d)
}

func (l Layout) SetText(text string) {
	ctext := C.CString(text)
	C.pango_layout_set_text(l.L, ctext, C.int(len(text)))
	C.free(unsafe.Pointer(ctext))
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
	runtime.SetFinalizer(d, func(d FontDescription) {
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
	runtime.SetFinalizer(f, func(f Font) {
		C.g_object_unref(f.f)
	})
	return f
}

type Font struct{ f *C.PangoFont }

func (f Font) Metrics() FontMetrics {
	cm := C.pango_font_get_metrics(f.f, nil) // TODO: support fontsets
	m := FontMetrics{cm}
	runtime.SetFinalizer(m, func(m FontMetrics) {
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
