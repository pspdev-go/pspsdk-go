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
		debugscreen.Printf(
			"TinyGo PSP future app layer\nPress START to exit\n\nCount: %d\n",
			count,
		)

		ctrl.ReadBufferPositive(&pad, 1)
		if pad.Buttons&ctrl.PSP_CTRL_START != 0 {
			break
		}

		count++
	}
	kernel.ExitGame()
}
