package str_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/pspdev-go/pspsdk-go/psp/utils/str"
	"github.com/stretchr/testify/assert"
)

func TestItoa(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		run  func() string
		want string
	}{
		{
			name: "int zero",
			run:  func() string { return str.Itoa(int(0)) },
			want: "0",
		},
		{
			name: "int negative",
			run:  func() string { return str.Itoa(int(-12345)) },
			want: "-12345",
		},
		{
			name: "int8 min",
			run:  func() string { return str.Itoa(int8(math.MinInt8)) },
			want: strconv.FormatInt(math.MinInt8, 10),
		},
		{
			name: "int16 positive",
			run:  func() string { return str.Itoa(int16(32767)) },
			want: "32767",
		},
		{
			name: "int32 min",
			run:  func() string { return str.Itoa(int32(math.MinInt32)) },
			want: strconv.FormatInt(math.MinInt32, 10),
		},
		{
			name: "int64 max",
			run:  func() string { return str.Itoa(int64(math.MaxInt64)) },
			want: strconv.FormatInt(math.MaxInt64, 10),
		},
		{
			name: "uint zero",
			run:  func() string { return str.Itoa(uint(0)) },
			want: "0",
		},
		{
			name: "uint8 max",
			run:  func() string { return str.Itoa(uint8(math.MaxUint8)) },
			want: strconv.FormatUint(math.MaxUint8, 10),
		},
		{
			name: "uint16 max",
			run:  func() string { return str.Itoa(uint16(math.MaxUint16)) },
			want: strconv.FormatUint(math.MaxUint16, 10),
		},
		{
			name: "uint32 max",
			run:  func() string { return str.Itoa(uint32(math.MaxUint32)) },
			want: strconv.FormatUint(math.MaxUint32, 10),
		},
		{
			name: "uint64 max",
			run:  func() string { return str.Itoa(uint64(math.MaxUint64)) },
			want: strconv.FormatUint(math.MaxUint64, 10),
		},
		{
			name: "uintptr value",
			run:  func() string { return str.Itoa(uintptr(123456)) },
			want: "123456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.run())
		})
	}
}

func TestFormatUnsigned(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input uint64
		want  string
	}{
		{
			name:  "zero",
			input: 0,
			want:  "0",
		},
		{
			name:  "single digit",
			input: 7,
			want:  "7",
		},
		{
			name:  "multiple digits",
			input: 1234567890,
			want:  "1234567890",
		},
		{
			name:  "max value",
			input: math.MaxUint64,
			want:  strconv.FormatUint(math.MaxUint64, 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, str.FormatUnsigned(tt.input))
		})
	}
}

func TestFormatSigned(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input int64
		want  string
	}{
		{
			name:  "zero",
			input: 0,
			want:  "0",
		},
		{
			name:  "positive value",
			input: 987654321,
			want:  "987654321",
		},
		{
			name:  "negative value",
			input: -987654321,
			want:  "-987654321",
		},
		{
			name:  "min value",
			input: math.MinInt64,
			want:  strconv.FormatInt(math.MinInt64, 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, str.FormatSigned(tt.input))
		})
	}
}
