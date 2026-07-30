package str

func Join(sep string, elements ...string) string {
	if len(elements) == 0 {
		return ""
	}
	if len(elements) == 1 {
		return elements[0]
	}

	n := len(sep) * (len(elements) - 1)
	for _, s := range elements {
		n += len(s)
	}

	b := make([]byte, n)
	bp := copy(b, elements[0])
	for _, s := range elements[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
