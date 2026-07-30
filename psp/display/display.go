// Package display provides Go bindings for PSPSDK's pspdisplay.h.
package display

import "unsafe"

type PixelFormat int32

const (
	PixelFormat565  PixelFormat = 0
	PixelFormat5551 PixelFormat = 1
	PixelFormat4444 PixelFormat = 2
	PixelFormat8888 PixelFormat = 3
)

const (
	PSP_DISPLAY_PIXEL_FORMAT_565  = PixelFormat565
	PSP_DISPLAY_PIXEL_FORMAT_5551 = PixelFormat5551
	PSP_DISPLAY_PIXEL_FORMAT_4444 = PixelFormat4444
	PSP_DISPLAY_PIXEL_FORMAT_8888 = PixelFormat8888
)

type SetBufSync int32

const (
	SetBufNextHSync SetBufSync = 0
	SetBufNextVSync SetBufSync = 1

	SetBufImmediate = SetBufNextHSync
	SetBufNextFrame = SetBufNextVSync
)

const (
	PSP_DISPLAY_SETBUF_NEXTHSYNC = SetBufNextHSync
	PSP_DISPLAY_SETBUF_NEXTVSYNC = SetBufNextVSync
	PSP_DISPLAY_SETBUF_IMMEDIATE = SetBufImmediate
	PSP_DISPLAY_SETBUF_NEXTFRAME = SetBufNextFrame
)

type Mode int32

const (
	ModeLCD       Mode = 0
	ModeVESA1A    Mode = 0x1a
	ModePseudoVGA Mode = 0x60
)

const (
	PSP_DISPLAY_MODE_LCD        = ModeLCD
	PSP_DISPLAY_MODE_VESA1A     = ModeVESA1A
	PSP_DISPLAY_MODE_PSEUDO_VGA = ModePseudoVGA
)

const (
	ErrorOK       int32 = 0
	ErrorPointer  int32 = -2147483389 // 0x80000103
	ErrorArgument int32 = -2147483385 // 0x80000107
)

const (
	SCE_DISPLAY_ERROR_OK       = ErrorOK
	SCE_DISPLAY_ERROR_POINTER  = ErrorPointer
	SCE_DISPLAY_ERROR_ARGUMENT = ErrorArgument
)

func SetMode(mode Mode, width, height int32) int32 {
	return sceDisplaySetMode(int32(mode), width, height)
}

func GetMode(mode *Mode, width, height *int32) int32 {
	return sceDisplayGetMode((*int32)(unsafe.Pointer(mode)), width, height)
}

func SetFrameBuf(topAddr unsafe.Pointer, bufferWidth int32, pixelFormat PixelFormat, sync SetBufSync) int32 {
	return sceDisplaySetFrameBuf(topAddr, bufferWidth, int32(pixelFormat), int32(sync))
}

func GetFrameBuf(topAddr *unsafe.Pointer, bufferWidth *int32, pixelFormat *PixelFormat, sync SetBufSync) int32 {
	return sceDisplayGetFrameBuf(topAddr, bufferWidth, (*int32)(unsafe.Pointer(pixelFormat)), int32(sync))
}

func GetVcount() uint32           { return sceDisplayGetVcount() }
func WaitVblankStart() int32      { return sceDisplayWaitVblankStart() }
func WaitVblankStartCB() int32    { return sceDisplayWaitVblankStartCB() }
func WaitVblank() int32           { return sceDisplayWaitVblank() }
func WaitVblankCB() int32         { return sceDisplayWaitVblankCB() }
func GetAccumulatedHcount() int32 { return sceDisplayGetAccumulatedHcount() }
func GetCurrentHcount() int32     { return sceDisplayGetCurrentHcount() }
func GetFramePerSec() float32     { return sceDisplayGetFramePerSec() }
func IsForeground() bool          { return sceDisplayIsForeground() != 0 }
func IsVblank() bool              { return sceDisplayIsVblank() != 0 }

//go:linkname sceDisplaySetMode sceDisplaySetMode
func sceDisplaySetMode(mode, width, height int32) int32

//go:linkname sceDisplayGetMode sceDisplayGetMode
func sceDisplayGetMode(mode, width, height *int32) int32

//go:linkname sceDisplaySetFrameBuf sceDisplaySetFrameBuf
func sceDisplaySetFrameBuf(topAddr unsafe.Pointer, bufferWidth, pixelFormat, sync int32) int32

//go:linkname sceDisplayGetFrameBuf sceDisplayGetFrameBuf
func sceDisplayGetFrameBuf(topAddr *unsafe.Pointer, bufferWidth, pixelFormat *int32, sync int32) int32

//go:linkname sceDisplayGetVcount sceDisplayGetVcount
func sceDisplayGetVcount() uint32

//go:linkname sceDisplayWaitVblankStart sceDisplayWaitVblankStart
func sceDisplayWaitVblankStart() int32

//go:linkname sceDisplayWaitVblankStartCB sceDisplayWaitVblankStartCB
func sceDisplayWaitVblankStartCB() int32

//go:linkname sceDisplayWaitVblank sceDisplayWaitVblank
func sceDisplayWaitVblank() int32

//go:linkname sceDisplayWaitVblankCB sceDisplayWaitVblankCB
func sceDisplayWaitVblankCB() int32

//go:linkname sceDisplayGetAccumulatedHcount sceDisplayGetAccumulatedHcount
func sceDisplayGetAccumulatedHcount() int32

//go:linkname sceDisplayGetCurrentHcount sceDisplayGetCurrentHcount
func sceDisplayGetCurrentHcount() int32

//go:linkname sceDisplayGetFramePerSec sceDisplayGetFramePerSec
func sceDisplayGetFramePerSec() float32

//go:linkname sceDisplayIsForeground sceDisplayIsForeground
func sceDisplayIsForeground() int32

//go:linkname sceDisplayIsVblank sceDisplayIsVblank
func sceDisplayIsVblank() int32
