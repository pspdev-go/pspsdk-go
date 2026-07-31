// Package format implements the small printf subset used by PSP output APIs.
package format

import (
	"unsafe"

	pspstr "github.com/pspdev-go/pspsdk-go/psp/utils/str"
)

// Sprintf formats common integer, string, character, boolean, and pointer
// values without pulling in Go's comparatively large fmt package.
func Sprintf(format string, args ...any) string {
	out := make([]byte, 0, len(format)+16)
	argIndex := 0
	for i := 0; i < len(format); i++ {
		if format[i] != '%' {
			out = append(out, format[i])
			continue
		}
		i++
		if i >= len(format) {
			out = append(out, '%')
			break
		}
		if format[i] == '%' {
			out = append(out, '%')
			continue
		}

		pad, left, width := byte(' '), false, 0
		if format[i] == '-' {
			left = true
			i++
		} else if format[i] == '0' {
			pad = '0'
			i++
		}
		for i < len(format) && format[i] >= '0' && format[i] <= '9' {
			width = width*10 + int(format[i]-'0')
			i++
		}
		if i >= len(format) {
			out = append(out, '%')
			break
		}
		verb := format[i]
		if argIndex >= len(args) {
			out = append(out, "%!"+string(verb)+"(MISSING)"...)
			continue
		}
		value := args[argIndex]
		argIndex++
		text, ok := formatValue(verb, value)
		if !ok {
			out = append(out, "%!"+string(verb)+"(INVALID)"...)
			continue
		}
		if width > len(text) && !left {
			out = appendPad(out, pad, width-len(text), &text)
		}
		out = append(out, text...)
		if width > len(text) && left {
			for n := len(text); n < width; n++ {
				out = append(out, ' ')
			}
		}
	}
	return string(out)
}

func appendPad(out []byte, pad byte, count int, text *string) []byte {
	if pad == '0' && len(*text) > 0 && ((*text)[0] == '-' || (*text)[0] == '+') {
		out = append(out, (*text)[0])
		*text = (*text)[1:]
	}
	for i := 0; i < count; i++ {
		out = append(out, pad)
	}
	return out
}

func formatValue(verb byte, value any) (string, bool) {
	switch verb {
	case 's':
		switch v := value.(type) {
		case string:
			return v, true
		case []byte:
			return string(v), true
		}
	case 't':
		if v, ok := value.(bool); ok {
			if v {
				return "true", true
			}
			return "false", true
		}
	case 'c':
		if v, ok := signed(value); ok {
			return string(rune(v)), true
		}
		if v, ok := unsigned(value); ok {
			return string(rune(v)), true
		}
	case 'd', 'i':
		if v, ok := signed(value); ok {
			return pspstr.Itoa(v), true
		}
		if v, ok := unsigned(value); ok {
			return unsignedBase(v, 10, false), true
		}
	case 'u':
		if v, ok := unsigned(value); ok {
			return unsignedBase(v, 10, false), true
		}
		if v, ok := signed(value); ok {
			return unsignedBase(uint64(v), 10, false), true
		}
	case 'x', 'X', 'o', 'b':
		v, ok := unsigned(value)
		if !ok {
			if signedValue, signedOK := signed(value); signedOK {
				v, ok = uint64(signedValue), true
			}
		}
		if ok {
			base := uint64(16)
			if verb == 'o' {
				base = 8
			} else if verb == 'b' {
				base = 2
			}
			return unsignedBase(v, base, verb == 'X'), true
		}
	case 'p':
		switch v := value.(type) {
		case unsafe.Pointer:
			return "0x" + unsignedBase(uint64(uintptr(v)), 16, false), true
		case uintptr:
			return "0x" + unsignedBase(uint64(v), 16, false), true
		}
	}
	return "", false
}

func signed(value any) (int64, bool) {
	switch v := value.(type) {
	case int:
		return int64(v), true
	case int8:
		return int64(v), true
	case int16:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return v, true
	}
	return 0, false
}

func unsigned(value any) (uint64, bool) {
	switch v := value.(type) {
	case uint:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case uint64:
		return v, true
	case uintptr:
		return uint64(v), true
	}
	return 0, false
}

func unsignedBase(value, base uint64, upper bool) string {
	if value == 0 {
		return "0"
	}
	digits := "0123456789abcdef"
	if upper {
		digits = "0123456789ABCDEF"
	}
	var buf [64]byte
	i := len(buf)
	for value != 0 {
		i--
		buf[i] = digits[value%base]
		value /= base
	}
	return string(buf[i:])
}
