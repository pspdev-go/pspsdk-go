package vlf

import (
	_ "unsafe"

	pspformat "github.com/pspdev-go/pspsdk-go/psp/internal/format"
)

// AddText formats values in Go and adds the resulting text to the VLF GUI.
func AddText(x, y int32, format string, args ...any) int32 {
	data := append([]byte(pspformat.Sprintf(format, args...)), 0)
	return addText(x, y, &data[0])
}

// SetText formats values in Go and replaces the contents of a VLF text item.
func SetText(text int32, format string, args ...any) int32 {
	data := append([]byte(pspformat.Sprintf(format, args...)), 0)
	return setText(text, &data[0])
}

//go:linkname addText pspsdk_go_vlf_add_text
func addText(x, y int32, text *byte) int32

//go:linkname setText pspsdk_go_vlf_set_text
func setText(text int32, value *byte) int32
