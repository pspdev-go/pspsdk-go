// Package ctrl provides Go bindings for PSPSDK's pspctrl.h.
package ctrl

import _ "unsafe"

type Mode int32

const (
	ModeDigital Mode = 0
	ModeAnalog  Mode = 1
)

const (
	PSP_CTRL_MODE_DIGITAL = ModeDigital
	PSP_CTRL_MODE_ANALOG  = ModeAnalog
)

type Buttons uint32

const (
	ButtonSelect      Buttons = 0x000001
	ButtonL3          Buttons = 0x000002
	ButtonR3          Buttons = 0x000004
	ButtonStart       Buttons = 0x000008
	ButtonUp          Buttons = 0x000010
	ButtonRight       Buttons = 0x000020
	ButtonDown        Buttons = 0x000040
	ButtonLeft        Buttons = 0x000080
	ButtonLTrigger    Buttons = 0x000100
	ButtonRTrigger    Buttons = 0x000200
	ButtonL2          Buttons = 0x000100
	ButtonR2          Buttons = 0x000200
	ButtonL1          Buttons = 0x000400
	ButtonR1          Buttons = 0x000800
	ButtonTriangle    Buttons = 0x001000
	ButtonCircle      Buttons = 0x002000
	ButtonCross       Buttons = 0x004000
	ButtonSquare      Buttons = 0x008000
	ButtonHome        Buttons = 0x010000
	ButtonHold        Buttons = 0x020000
	ButtonWLANUp      Buttons = 0x040000
	ButtonRemote      Buttons = 0x080000
	ButtonVolumeUp    Buttons = 0x100000
	ButtonVolumeDown  Buttons = 0x200000
	ButtonScreen      Buttons = 0x400000
	ButtonNote        Buttons = 0x800000
	ButtonDisc        Buttons = 0x1000000
	ButtonMemoryStick Buttons = 0x2000000
)

// Legacy PSPSDK-style names.
const (
	PSP_CTRL_SELECT   = ButtonSelect
	PSP_CTRL_L3       = ButtonL3
	PSP_CTRL_R3       = ButtonR3
	PSP_CTRL_START    = ButtonStart
	PSP_CTRL_UP       = ButtonUp
	PSP_CTRL_RIGHT    = ButtonRight
	PSP_CTRL_DOWN     = ButtonDown
	PSP_CTRL_LEFT     = ButtonLeft
	PSP_CTRL_LTRIGGER = ButtonLTrigger
	PSP_CTRL_RTRIGGER = ButtonRTrigger
	PSP_CTRL_L2       = ButtonL2
	PSP_CTRL_R2       = ButtonR2
	PSP_CTRL_L1       = ButtonL1
	PSP_CTRL_R1       = ButtonR1
	PSP_CTRL_TRIANGLE = ButtonTriangle
	PSP_CTRL_CIRCLE   = ButtonCircle
	PSP_CTRL_CROSS    = ButtonCross
	PSP_CTRL_SQUARE   = ButtonSquare
	PSP_CTRL_HOME     = ButtonHome
	PSP_CTRL_HOLD     = ButtonHold
	PSP_CTRL_WLAN_UP  = ButtonWLANUp
	PSP_CTRL_REMOTE   = ButtonRemote
	PSP_CTRL_VOLUP    = ButtonVolumeUp
	PSP_CTRL_VOLDOWN  = ButtonVolumeDown
	PSP_CTRL_SCREEN   = ButtonScreen
	PSP_CTRL_NOTE     = ButtonNote
	PSP_CTRL_DISC     = ButtonDisc
	PSP_CTRL_MS       = ButtonMemoryStick
)

// SceCtrlData has the same 16-byte layout as the C structure.
type SceCtrlData struct {
	TimeStamp uint32
	Buttons   Buttons
	Lx        uint8
	Ly        uint8
	Rx        uint8
	Ry        uint8
	Reserved  [4]uint8
}

type SceCtrlLatch struct {
	Make    Buttons
	Break   Buttons
	Press   Buttons
	Release Buttons
}

func SetSamplingCycle(cycle int32) int32  { return sceCtrlSetSamplingCycle(cycle) }
func GetSamplingCycle(cycle *int32) int32 { return sceCtrlGetSamplingCycle(cycle) }
func SetSamplingMode(mode Mode) int32     { return sceCtrlSetSamplingMode(int32(mode)) }
func GetSamplingMode(mode *Mode) int32    { return sceCtrlGetSamplingMode(mode) }
func PeekBufferPositive(data *SceCtrlData, count int32) int32 {
	return sceCtrlPeekBufferPositive(data, count)
}
func PeekBufferNegative(data *SceCtrlData, count int32) int32 {
	return sceCtrlPeekBufferNegative(data, count)
}
func ReadBufferPositive(data *SceCtrlData, count int32) int32 {
	return sceCtrlReadBufferPositive(data, count)
}
func ReadBufferNegative(data *SceCtrlData, count int32) int32 {
	return sceCtrlReadBufferNegative(data, count)
}
func PeekLatch(data *SceCtrlLatch) int32 { return sceCtrlPeekLatch(data) }
func ReadLatch(data *SceCtrlLatch) int32 { return sceCtrlReadLatch(data) }
func SetIdleCancelThreshold(idleReset, idleBack int32) int32 {
	return sceCtrlSetIdleCancelThreshold(idleReset, idleBack)
}
func GetIdleCancelThreshold(idleReset, idleBack *int32) int32 {
	return sceCtrlGetIdleCancelThreshold(idleReset, idleBack)
}

//go:linkname sceCtrlSetSamplingCycle sceCtrlSetSamplingCycle
func sceCtrlSetSamplingCycle(cycle int32) int32

//go:linkname sceCtrlGetSamplingCycle sceCtrlGetSamplingCycle
func sceCtrlGetSamplingCycle(cycle *int32) int32

//go:linkname sceCtrlSetSamplingMode sceCtrlSetSamplingMode
func sceCtrlSetSamplingMode(mode int32) int32

//go:linkname sceCtrlGetSamplingMode sceCtrlGetSamplingMode
func sceCtrlGetSamplingMode(mode *Mode) int32

//go:linkname sceCtrlPeekBufferPositive sceCtrlPeekBufferPositive
func sceCtrlPeekBufferPositive(data *SceCtrlData, count int32) int32

//go:linkname sceCtrlPeekBufferNegative sceCtrlPeekBufferNegative
func sceCtrlPeekBufferNegative(data *SceCtrlData, count int32) int32

//go:linkname sceCtrlReadBufferPositive sceCtrlReadBufferPositive
func sceCtrlReadBufferPositive(data *SceCtrlData, count int32) int32

//go:linkname sceCtrlReadBufferNegative sceCtrlReadBufferNegative
func sceCtrlReadBufferNegative(data *SceCtrlData, count int32) int32

//go:linkname sceCtrlPeekLatch sceCtrlPeekLatch
func sceCtrlPeekLatch(data *SceCtrlLatch) int32

//go:linkname sceCtrlReadLatch sceCtrlReadLatch
func sceCtrlReadLatch(data *SceCtrlLatch) int32

//go:linkname sceCtrlSetIdleCancelThreshold sceCtrlSetIdleCancelThreshold
func sceCtrlSetIdleCancelThreshold(idleReset, idleBack int32) int32

//go:linkname sceCtrlGetIdleCancelThreshold sceCtrlGetIdleCancelThreshold
func sceCtrlGetIdleCancelThreshold(idleReset, idleBack *int32) int32
