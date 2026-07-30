package main

import (
	"github.com/pspdev-go/pspsdk-go/psp/ctrl"
	"github.com/pspdev-go/pspsdk-go/psp/debugscreen"
	"github.com/pspdev-go/pspsdk-go/psp/display"
	"github.com/pspdev-go/pspsdk-go/psp/kernel"
)

func main() {
	debugscreen.Init()
	ctrl.SetSamplingCycle(0)
	ctrl.SetSamplingMode(ctrl.ModeDigital)

	var count int
	var pad ctrl.SceCtrlData

	for {
		display.WaitVblankStart()
		debugscreen.SetXY(0, 0)
		debugscreen.PutString("TinyGo PSP future app layer\n")
		debugscreen.PutString("Press START to exit\n")
		debugscreen.PutString("\n")
		debugscreen.PutString("Count: ")
		debugscreen.PutInt(count)
		debugscreen.PutString("\n")

		ctrl.ReadBufferPositive(&pad, 1)
		if pad.Buttons&ctrl.PSP_CTRL_START != 0 {
			break
		}

		count++
	}
	kernel.ExitGame()
}
