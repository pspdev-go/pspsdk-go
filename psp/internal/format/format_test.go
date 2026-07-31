package format

import "testing"

func TestSprintf(t *testing.T) {
	tests := []struct {
		format string
		args   []any
		want   string
	}{
		{"plain", nil, "plain"},
		{"100%%", nil, "100%"},
		{"%s %d %u", []any{"value", -12, uint32(34)}, "value -12 34"},
		{"%08X", []any{uint32(0x12ab)}, "000012AB"},
		{"%-5s!", []any{"go"}, "go   !"},
		{"%04d", []any{-7}, "-007"},
		{"%x %o %b", []any{255, 8, 5}, "ff 10 101"},
		{"%c %t", []any{'A', true}, "A true"},
		{"%d", nil, "%!d(MISSING)"},
		{"%f", []any{1}, "%!f(INVALID)"},
	}
	for _, test := range tests {
		if got := Sprintf(test.format, test.args...); got != test.want {
			t.Errorf("Sprintf(%q) = %q, want %q", test.format, got, test.want)
		}
	}
}
