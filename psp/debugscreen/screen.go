// Package debugscreen provides Go bindings for PSPSDK's pspdebug.h.
package debugscreen

import (
	"unsafe"

	pspstr "github.com/pspdev-go/pspsdk-go/psp/utils/str"
)

// RegBlock mirrors PspDebugRegBlock.
type RegBlock struct {
	Frame    [6]uint32
	R        [32]uint32
	Status   uint32
	Lo       uint32
	Hi       uint32
	BadVAddr uint32
	Cause    uint32
	EPC      uint32
	FPR      [32]float32
	FSR      uint32
	FIR      uint32
	FramePtr uint32
	Unused   uint32
	Index    uint32
	Random   uint32
	EntryLo0 uint32
	EntryLo1 uint32
	Context  uint32
	PageMask uint32
	Wired    uint32
	Cop0_7   uint32
	Cop0_8   uint32
	Cop0_9   uint32
	EntryHi  uint32
	Cop0_11  uint32
	Cop0_12  uint32
	Cop0_13  uint32
	Cop0_14  uint32
	PRId     uint32
	Padding  [100]uint32
}

type PspDebugRegBlock = RegBlock

// StackTrace mirrors PspDebugStackTrace.
type StackTrace struct {
	CallAddr uint32
	FuncAddr uint32
}

type PspDebugStackTrace = StackTrace

// ProfilerRegs mirrors PspDebugProfilerRegs.
type ProfilerRegs struct {
	Enable        uint32
	SystemClock   uint32
	CPUClock      uint32
	Internal      uint32
	Memory        uint32
	Cop0          uint32
	VFPU          uint32
	Sleep         uint32
	BusAccess     uint32
	UncachedLoad  uint32
	UncachedStore uint32
	CachedLoad    uint32
	CachedStore   uint32
	IMiss         uint32
	DMiss         uint32
	DWriteback    uint32
	Cop0Inst      uint32
	FPUInst       uint32
	VFPUInst      uint32
	LocalBus      uint32
}

type PspDebugProfilerRegs = ProfilerRegs

func Init() { pspDebugScreenInit() }
func InitEx(vramBase unsafe.Pointer, mode int32, setup bool) {
	var setupValue int32
	if setup {
		setupValue = 1
	}
	pspDebugScreenInitEx(vramBase, mode, setupValue)
}
func EnableBackColor(enable bool) {
	var value int32
	if enable {
		value = 1
	}
	pspDebugScreenEnableBackColor(value)
}
func SetBackColor(color uint32) { pspDebugScreenSetBackColor(color) }
func SetTextColor(color uint32) { pspDebugScreenSetTextColor(color) }
func SetColorMode(mode int32)   { pspDebugScreenSetColorMode(mode) }
func PutChar(x, y int32, color uint32, ch byte) {
	pspDebugScreenPutChar(x, y, color, ch)
}
func SetXY(x, y int32)       { pspDebugScreenSetXY(x, y) }
func Home()                  { SetXY(0, 0) }
func SetOffset(offset int32) { pspDebugScreenSetOffset(offset) }
func SetBase(base *uint32)   { pspDebugScreenSetBase(base) }
func GetX() int32            { return pspDebugScreenGetX() }
func GetY() int32            { return pspDebugScreenGetY() }
func Clear()                 { pspDebugScreenClear() }

// PrintData prints data without requiring a trailing NUL byte.
func PrintData(data []byte) int32 {
	if len(data) == 0 {
		return 0
	}
	return pspDebugScreenPrintData(&data[0], int32(len(data)))
}

// PutString prints a Go string. Embedded NUL bytes terminate the output.
func PutString(s string) int32 {
	data := append([]byte(s), 0)
	return pspDebugScreenPuts(&data[0])
}

func PutHex32(value uint32) int32 {
	const digits = "0123456789ABCDEF"
	var data [8]byte
	for i := len(data) - 1; i >= 0; i-- {
		data[i] = digits[value&0xf]
		value >>= 4
	}
	return PrintData(data[:])
}

func PutInt(value int) int32 {
	return PutString(pspstr.Itoa(value))
}

func GetStackTrace(results []uint32) int32 {
	if len(results) == 0 {
		return 0
	}
	return pspDebugGetStackTrace(&results[0], int32(len(results)))
}

func ClearLineEnable()  { pspDebugScreenClearLineEnable() }
func ClearLineDisable() { pspDebugScreenClearLineDisable() }

// Handler installation functions accept a C-ABI function pointer. Passing nil
// restores or disables the handler as documented by PSPSDK. A Go function is
// not itself a C-ABI function pointer.
func InstallErrorHandler(handler unsafe.Pointer) int32 {
	return pspDebugInstallErrorHandler(handler)
}
func DumpException(regs *RegBlock) { pspDebugDumpException(regs) }
func InstallKprintfHandler(handler unsafe.Pointer) int32 {
	return pspDebugInstallKprintfHandler(handler)
}
func GetStackTrace2(regs *RegBlock, trace []StackTrace) int32 {
	if len(trace) == 0 {
		return 0
	}
	return pspDebugGetStackTrace2(regs, &trace[0], int32(len(trace)))
}

func ProfilerEnable()                    { pspDebugProfilerEnable() }
func ProfilerDisable()                   { pspDebugProfilerDisable() }
func ProfilerClear()                     { pspDebugProfilerClear() }
func ProfilerGetRegs(regs *ProfilerRegs) { pspDebugProfilerGetRegs(regs) }
func ProfilerPrint()                     { pspDebugProfilerPrint() }

func InstallStdinHandler(handler unsafe.Pointer) int32 {
	return pspDebugInstallStdinHandler(handler)
}
func InstallStdoutHandler(handler unsafe.Pointer) int32 {
	return pspDebugInstallStdoutHandler(handler)
}
func InstallStderrHandler(handler unsafe.Pointer) int32 {
	return pspDebugInstallStderrHandler(handler)
}

func SioPutchar(ch int32) { pspDebugSioPutchar(ch) }
func SioGetchar() int32   { return pspDebugSioGetchar() }
func SioPuts(s string) {
	data := append([]byte(s), 0)
	pspDebugSioPuts(&data[0])
}
func SioPutData(data []byte) int32 {
	if len(data) == 0 {
		return 0
	}
	return pspDebugSioPutData(&data[0], int32(len(data)))
}
func SioPutText(data []byte) int32 {
	if len(data) == 0 {
		return 0
	}
	return pspDebugSioPutText(&data[0], int32(len(data)))
}
func SioInit()              { pspDebugSioInit() }
func SioSetBaud(baud int32) { pspDebugSioSetBaud(baud) }
func EnablePutchar()        { pspDebugEnablePutchar() }
func SioInstallKprintf()    { pspDebugSioInstallKprintf() }
func GDBStubInit()          { pspDebugGdbStubInit() }
func Breakpoint()           { pspDebugBreakpoint() }
func SioEnableKprintf()     { pspDebugSioEnableKprintf() }
func SioDisableKprintf()    { pspDebugSioDisableKprintf() }
func ScreenshotSave(filename string) int32 {
	data := append([]byte(filename), 0)
	return pspScreenshotSave(&data[0])
}

// pspDebugScreenPrintf and pspDebugScreenKprintf are deliberately not bound:
// Go cannot directly call C variadic functions. Use PutString or PrintData.

//go:linkname pspDebugScreenInit pspDebugScreenInit
func pspDebugScreenInit()

//go:linkname pspDebugScreenInitEx pspDebugScreenInitEx
func pspDebugScreenInitEx(vramBase unsafe.Pointer, mode, setup int32)

//go:linkname pspDebugScreenEnableBackColor pspDebugScreenEnableBackColor
func pspDebugScreenEnableBackColor(enable int32)

//go:linkname pspDebugScreenSetBackColor pspDebugScreenSetBackColor
func pspDebugScreenSetBackColor(color uint32)

//go:linkname pspDebugScreenSetTextColor pspDebugScreenSetTextColor
func pspDebugScreenSetTextColor(color uint32)

//go:linkname pspDebugScreenSetColorMode pspDebugScreenSetColorMode
func pspDebugScreenSetColorMode(mode int32)

//go:linkname pspDebugScreenPutChar pspDebugScreenPutChar
func pspDebugScreenPutChar(x, y int32, color uint32, ch byte)

//go:linkname pspDebugScreenSetXY pspDebugScreenSetXY
func pspDebugScreenSetXY(x, y int32)

//go:linkname pspDebugScreenSetOffset pspDebugScreenSetOffset
func pspDebugScreenSetOffset(offset int32)

//go:linkname pspDebugScreenSetBase pspDebugScreenSetBase
func pspDebugScreenSetBase(base *uint32)

//go:linkname pspDebugScreenGetX pspDebugScreenGetX
func pspDebugScreenGetX() int32

//go:linkname pspDebugScreenGetY pspDebugScreenGetY
func pspDebugScreenGetY() int32

//go:linkname pspDebugScreenClear pspDebugScreenClear
func pspDebugScreenClear()

//go:linkname pspDebugScreenPrintData pspDebugScreenPrintData
func pspDebugScreenPrintData(data *byte, size int32) int32

//go:linkname pspDebugScreenPuts pspDebugScreenPuts
func pspDebugScreenPuts(data *byte) int32

//go:linkname pspDebugGetStackTrace pspDebugGetStackTrace
func pspDebugGetStackTrace(results *uint32, max int32) int32

//go:linkname pspDebugScreenClearLineEnable pspDebugScreenClearLineEnable
func pspDebugScreenClearLineEnable()

//go:linkname pspDebugScreenClearLineDisable pspDebugScreenClearLineDisable
func pspDebugScreenClearLineDisable()

//go:linkname pspDebugInstallErrorHandler pspDebugInstallErrorHandler
func pspDebugInstallErrorHandler(handler unsafe.Pointer) int32

//go:linkname pspDebugDumpException pspDebugDumpException
func pspDebugDumpException(regs *RegBlock)

//go:linkname pspDebugInstallKprintfHandler pspDebugInstallKprintfHandler
func pspDebugInstallKprintfHandler(handler unsafe.Pointer) int32

//go:linkname pspDebugGetStackTrace2 pspDebugGetStackTrace2
func pspDebugGetStackTrace2(regs *RegBlock, trace *StackTrace, max int32) int32

//go:linkname pspDebugProfilerEnable pspDebugProfilerEnable
func pspDebugProfilerEnable()

//go:linkname pspDebugProfilerDisable pspDebugProfilerDisable
func pspDebugProfilerDisable()

//go:linkname pspDebugProfilerClear pspDebugProfilerClear
func pspDebugProfilerClear()

//go:linkname pspDebugProfilerGetRegs pspDebugProfilerGetRegs
func pspDebugProfilerGetRegs(regs *ProfilerRegs)

//go:linkname pspDebugProfilerPrint pspDebugProfilerPrint
func pspDebugProfilerPrint()

//go:linkname pspDebugInstallStdinHandler pspDebugInstallStdinHandler
func pspDebugInstallStdinHandler(handler unsafe.Pointer) int32

//go:linkname pspDebugInstallStdoutHandler pspDebugInstallStdoutHandler
func pspDebugInstallStdoutHandler(handler unsafe.Pointer) int32

//go:linkname pspDebugInstallStderrHandler pspDebugInstallStderrHandler
func pspDebugInstallStderrHandler(handler unsafe.Pointer) int32

//go:linkname pspDebugSioPutchar pspDebugSioPutchar
func pspDebugSioPutchar(ch int32)

//go:linkname pspDebugSioGetchar pspDebugSioGetchar
func pspDebugSioGetchar() int32

//go:linkname pspDebugSioPuts pspDebugSioPuts
func pspDebugSioPuts(data *byte)

//go:linkname pspDebugSioPutData pspDebugSioPutData
func pspDebugSioPutData(data *byte, length int32) int32

//go:linkname pspDebugSioPutText pspDebugSioPutText
func pspDebugSioPutText(data *byte, length int32) int32

//go:linkname pspDebugSioInit pspDebugSioInit
func pspDebugSioInit()

//go:linkname pspDebugSioSetBaud pspDebugSioSetBaud
func pspDebugSioSetBaud(baud int32)

//go:linkname pspDebugEnablePutchar pspDebugEnablePutchar
func pspDebugEnablePutchar()

//go:linkname pspDebugSioInstallKprintf pspDebugSioInstallKprintf
func pspDebugSioInstallKprintf()

//go:linkname pspDebugGdbStubInit pspDebugGdbStubInit
func pspDebugGdbStubInit()

//go:linkname pspDebugBreakpoint pspDebugBreakpoint
func pspDebugBreakpoint()

//go:linkname pspDebugSioEnableKprintf pspDebugSioEnableKprintf
func pspDebugSioEnableKprintf()

//go:linkname pspDebugSioDisableKprintf pspDebugSioDisableKprintf
func pspDebugSioDisableKprintf()

//go:linkname pspScreenshotSave pspScreenshotSave
func pspScreenshotSave(filename *byte) int32
