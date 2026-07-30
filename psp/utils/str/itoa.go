package str

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

// Itoa converts an integer value to its decimal string representation.
func Itoa[T Integer](v T) string {
	if v == 0 {
		return "0"
	}

	switch x := any(v).(type) {
	case int:
		return formatSigned(int64(x))
	case int8:
		return formatSigned(int64(x))
	case int16:
		return formatSigned(int64(x))
	case int32:
		return formatSigned(int64(x))
	case int64:
		return formatSigned(x)

	case uint:
		return formatUnsigned(uint64(x))
	case uint8:
		return formatUnsigned(uint64(x))
	case uint16:
		return formatUnsigned(uint64(x))
	case uint32:
		return formatUnsigned(uint64(x))
	case uint64:
		return formatUnsigned(x)
	case uintptr:
		return formatUnsigned(uint64(x))
	}

	return ""
}

// formatUnsigned formats an unsigned integer as a base-10 string.
func formatUnsigned(v uint64) string {
	if v == 0 {
		return "0"
	}

	var buf [20]byte
	i := len(buf)

	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}

	return string(buf[i:])
}

// formatSigned formats a signed integer as a base-10 string.
func formatSigned(v int64) string {
	if v >= 0 {
		return formatUnsigned(uint64(v))
	}

	u := uint64(-(v + 1))
	u += 1

	var buf [20]byte
	i := len(buf)

	for u > 0 {
		i--
		buf[i] = byte('0' + u%10)
		u /= 10
	}

	i--
	buf[i] = '-'
	return string(buf[i:])
}
