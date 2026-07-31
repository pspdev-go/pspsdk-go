# PSPSDK Go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/pspdev-go/pspsdk-go.svg)](https://pkg.go.dev/github.com/pspdev-go/pspsdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/pspdev-go/pspsdk-go)](https://goreportcard.com/report/github.com/pspdev-go/pspsdk-go)
[![CI](https://github.com/pspdev-go/pspsdk-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/pspdev-go/pspsdk-go/actions/workflows/ci.yaml)

`pspsdk-go` provides Go bindings and ABI adapters for writing PSP homebrew
with PSPSDK. Applications import packages below
`github.com/pspdev-go/pspsdk-go/psp` and are compiled with
[`pspgo`](https://github.com/pspdev-go/pspgo).

Use `pspgo` to build `pspsdk-go` applications. A regular `go build` produces a
host executable, while invoking TinyGo directly skips the bridge selection,
PSPSDK library resolution, PSP linking, and `EBOOT.PBP` packaging performed by
`pspgo`.

## Requirements

- [Go](https://go.dev/)
- [pspgo](https://github.com/pspdev-go/pspgo)
- [TinyGo with PSP support](https://github.com/pspdev-go/tinygo)
- CMake

`pspgo` uses `cmake --build`; it does not invoke Make directly. CMake can use
Make, Ninja, or another supported build backend.

Set `PSPDEV` to the PSPSDK installation and make its tools available:

```sh
export PSPDEV="$HOME/pspdev"
export PATH="$PSPDEV/bin:$PATH"
```

Install `pspgo` by following its
[installation guide](https://github.com/pspdev-go/pspgo#installation). It can
be built from source, installed with `go install`, or downloaded from the
GitHub Releases page.

## Example projects

You can reference the [pspsdk-example](https://github.com/pspdev-go/pspsdk-example) repository.

## Use in your own project

Add `pspsdk-go` to the Go module containing your application:

```sh
go mod init example.com/my-psp-app
go get github.com/pspdev-go/pspsdk-go@latest
```

`pspgo` resolves this module's directory from `go.mod`, including local paths
configured with a `replace` directive:

```sh
pspgo doctor
pspgo build .
```

The SDK path can also be stored in `pspgo.toml`:

```toml
title = "My PSP Game"
output = "my-game"
build_dir = "build/pspgo"
kernel_mode = false
```

The PSP target is built into TinyGo and selected by `pspgo`; applications do
not need to provide a `psp.json`.

Minimal application code can import the bindings directly:

```go
package main

import (
	"github.com/pspdev-go/pspsdk-go/psp/debugscreen"
	"github.com/pspdev-go/pspsdk-go/psp/kernel"
	"github.com/pspdev-go/pspsdk-go/psp/threadman"
)

func main() {
	debugscreen.Init()
	debugscreen.PutString("Hello from Go!")
	threadman.SceKernelDelayThread(5_000_000)
	kernel.ExitGame()
}
```

See [`example/main.go`](example/main.go) for a complete rotating-cube example.

## PSPSDK header mappings

The packages call PSPSDK symbols directly. The only remaining C bridge is
`bridge/main.c`, which defines the PSP module metadata and calls the TinyGo
entry point.

| Go package        | PSPSDK header                                              |
| ----------------- | ---------------------------------------------------------- |
| `psp/ctrl`        | `pspctrl.h`                                                |
| `psp/debugscreen` | `pspdebug.h`                                               |
| `psp/display`     | `pspdisplay.h`                                             |
| `psp/kernel`      | the currently used API from aggregate header `pspkernel.h` |

Functions, constants, and structures use Go-friendly names. Constants and
debug structure types also retain PSPSDK-style aliases where useful when
porting C code.

The variadic C functions `pspDebugScreenPrintf` and
`pspDebugScreenKprintf` cannot be called directly using the Go ABI. Use
`debugscreen.PutString`, `debugscreen.PrintData`, `debugscreen.PutInt`, or
`debugscreen.PutHex32` instead. `debugscreen.Printf` provides a lightweight
Go implementation of `%%`, `%s`, `%d`, `%i`, `%u`, `%x`, `%X`, `%o`, `%b`,
`%c`, `%t`, and `%p`, including field width, left alignment, and zero padding.
`debugscreen.Kprintf` and `kdebug.Kprintf` provide the same formatting support
for their corresponding kernel debug outputs.
Handler installation functions accept an
`unsafe.Pointer` to a C-ABI callback; an ordinary Go function is not such a
pointer.

### Generated bindings

All 140 top-level headers in the installed PSPSDK have a corresponding package
below `psp/`. This includes user-mode and kernel/driver `psp*.h` headers as
well as ARK, CFW, Vita POPS, VLF, and helper-library headers. The `psp` prefix
is removed from package names; for example:

| Header                  | Go package             |
| ----------------------- | ---------------------- |
| `pspaudio.h`            | `psp/audio`            |
| `pspnet_adhoc.h`        | `psp/net_adhoc`        |
| `psputility_savedata.h` | `psp/utility_savedata` |
| `pspctrl_kernel.h`      | `psp/ctrl_kernel`      |
| `pspnand_driver.h`      | `psp/nand_driver`      |
| `systemctrl_ark.h`      | `psp/systemctrl_ark`   |
| `vitapops.h`            | `psp/vitapops`         |
| `vlf.h`                 | `psp/vlf`              |

The generated bindings expose C functions with an exported first letter and
retain the rest of the PSPSDK name. For example,
`sceAudioChReserve` becomes `audio.SceAudioChReserve`. Enum constants retain
their original PSPSDK names. C pointers, structures, unions, arrays, and
callbacks that cannot be represented safely without a hand-maintained ABI
type are exposed as `unsafe.Pointer`.

Regenerate the bindings against the currently installed PSPSDK with:

```sh
python3 tools/genbindings.py
gofmt -w psp/*/bindings_gen.go
```

See [bindings-report.md](bindings-report.md) for the package generated for
each header, declaration counts, and unsupported variadic functions.

Kernel and driver packages add a library-requirement marker so automatic
resolution selects the kernel/driver archive even when a user-mode archive
exports an identically named symbol. Importing these packages does not grant
kernel privileges. Build a kernel-mode module when the API requires it by
setting `kernel_mode = true` in `pspgo.toml`:

```toml
kernel_mode = true
```

Kernel/driver APIs can modify firmware state or raw devices and must only be
used with hardware- and firmware-appropriate arguments.

ARK, CFW, Adrenaline, Vita POPS, and VLF packages require the corresponding
runtime environment and modules; successful static linking does not make
those APIs available on unsupported firmware. Some combinations of custom
firmware archives cause `psp-fixup-imports` to warn about import-stub order.
Treat that warning as a deployment compatibility issue and test the resulting
binary on the exact target firmware.

The variadic VLF text functions are available as `vlf.AddText` and
`vlf.SetText`, using the same lightweight Go formatter as
`debugscreen.Printf`.
