package gopancairo // import "go.vktec.org.uk/gopan/vtkcairo"

/*
#cgo pkg-config: pangocairo
#include <pango/pangocairo.h>
*/
import "C"

import (
	"go.vktec.org.uk/gopan"
	"go.vktec.org.uk/vtk/cairo"
	"unsafe"
)

type CairoLayout struct {
	gopan.Layout
	Cr cairo.Cairo
}

func CreateLayout(cr cairo.Cairo) CairoLayout {
	cl := C.pango_cairo_create_layout(cr.Cr)
	l := CairoLayout{ gopan.CreateLayoutFromC(unsafe.Pointer(cl)), cr }
	return l
}

func (l CairoLayout) Show() {
	C.pango_cairo_show_layout(l.Cr.Cr, l.Layout.L)
}

func (l CairoLayout) Update() {
	C.pango_cairo_update_layout(l.Cr.Cr, l.Layout.L)
}

func DefaultFontMap() gopan.FontMap {
	return gopan.CreateFontMapFromC(unsafe.Pointer(C.pango_cairo_font_map_get_default()))
}
