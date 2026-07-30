// Package kernel contains the small set of bindings currently used from
// pspkernel.h. That header is an aggregate of many PSPSDK subsystem headers;
// those subsystems should be exposed as their own Go packages.
package kernel

import _ "unsafe"

func ExitGame() { sceKernelExitGame() }

//go:linkname sceKernelExitGame sceKernelExitGame
func sceKernelExitGame()
