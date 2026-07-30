# PSPSDK Go

Write PSP Homebrew in Golang!

## Requirements

- [Golang](https://golang.org/dl/)
- [CMake](https://cmake.org/download/)
- [PSPSDK](https://github.com/pspdev/pspsdk)
- [Forked TinyGo](https://github.com/pspdev-go/tinygo)

## Building sample

```sh
./build-sample.sh
```

## PSPSDK header mappings

The packages call PSPSDK symbols directly. The only remaining C bridge is
`bridge/main.c`, which defines the PSP module metadata and calls the TinyGo
entry point.

| Go package | PSPSDK header |
| --- | --- |
| `psp/ctrl` | `pspctrl.h` |
| `psp/debugscreen` | `pspdebug.h` |
| `psp/display` | `pspdisplay.h` |
| `psp/kernel` | the currently used API from aggregate header `pspkernel.h` |

Functions, constants, and structures use Go-friendly names. Constants and
debug structure types also retain PSPSDK-style aliases where useful when
porting C code.

The variadic C functions `pspDebugScreenPrintf` and
`pspDebugScreenKprintf` cannot be called directly using the Go ABI. Use
`debugscreen.PutString`, `debugscreen.PrintData`, `debugscreen.PutInt`, or
`debugscreen.PutHex32` instead. Handler installation functions accept an
`unsafe.Pointer` to a C-ABI callback; an ordinary Go function is not such a
pointer.

### Generated user-mode bindings

Every user-mode `psp*.h` header in the installed PSPSDK has a corresponding
package below `psp/`. The `psp` prefix is removed from package names; for
example:

| Header | Go package |
| --- | --- |
| `pspaudio.h` | `psp/audio` |
| `pspnet_adhoc.h` | `psp/net_adhoc` |
| `psputility_savedata.h` | `psp/utility_savedata` |

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
