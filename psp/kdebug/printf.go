package kdebug

import (
	"unsafe"

	pspformat "github.com/pspdev-go/pspsdk-go/psp/internal/format"
)

// Kprintf formats values in Go and sends the resulting string to the kernel
// debug output. It supports the same verbs as debugscreen.Printf.
func Kprintf(format string, args ...any) {
	data := append([]byte(pspformat.Sprintf(format, args...)), 0)
	pspKdebugPuts(&data[0])
}

//go:linkname pspKdebugPuts psp_kdebug_puts
func pspKdebugPuts(text *byte)

var _ unsafe.Pointer
