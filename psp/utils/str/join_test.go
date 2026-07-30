package str_test

import (
	"strings"
	"testing"

	"github.com/pspdev-go/pspsdk-go/psp/utils/str"
	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		sep      string
		elements []string
		want     string
	}{
		{
			name:     "empty elements",
			sep:      ",",
			elements: []string{},
			want:     "",
		},
		{
			name:     "single element",
			sep:      ",",
			elements: []string{"only"},
			want:     "only",
		},
		{
			name:     "multiple elements",
			sep:      ",",
			elements: []string{"a", "b", "c"},
			want:     "a,b,c",
		},
		{
			name:     "empty separator",
			sep:      "",
			elements: []string{"go", "lang"},
			want:     "golang",
		},
		{
			name:     "empty strings in elements",
			sep:      "-",
			elements: []string{"a", "", "c"},
			want:     "a--c",
		},
		{
			name:     "multi character separator",
			sep:      " / ",
			elements: []string{"usr", "local", "bin"},
			want:     "usr / local / bin",
		},
		{
			name:     "long input",
			sep:      ":",
			elements: []string{strings.Repeat("x", 32), strings.Repeat("y", 32), strings.Repeat("z", 32)},
			want:     strings.Repeat("x", 32) + ":" + strings.Repeat("y", 32) + ":" + strings.Repeat("z", 32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, str.Join(tt.sep, tt.elements...))
		})
	}
}
