package stdio_kernel

import (
	"unsafe"

	pspformat "github.com/pspdev-go/pspsdk-go/psp/internal/format"
)

// Printf formats values in Go and writes them to a kernel file descriptor.
// It supports the same verbs as debugscreen.Printf.
func Printf(fd int32, format string, args ...any) int32 {
	data := append([]byte(pspformat.Sprintf(format, args...)), 0)
	return fdputs(fd, &data[0])
}

//go:linkname fdputs pspsdk_go_fdputs
func fdputs(fd int32, text *byte) int32

var _ unsafe.Pointer
