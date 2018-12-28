package binaryext

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutUint64(t *testing.T) {
	tests := []struct {
		in  uint64
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MaxUint64, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, test := range tests {
		out := PutUint64(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestPutUint32(t *testing.T) {
	tests := []struct {
		in  uint32
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MaxUint32, []byte{0xff, 0xff, 0xff, 0xff}},
	}
	for _, test := range tests {
		out := PutUint32(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestPutUint16(t *testing.T) {
	tests := []struct {
		in  uint16
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MaxUint16, []byte{0xff, 0xff}},
	}
	for _, test := range tests {
		out := PutUint16(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestPutUint8(t *testing.T) {
	tests := []struct {
		in  uint8
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MaxUint8, []byte{0xff}},
	}
	for _, test := range tests {
		out := PutUint8(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestUint64(t *testing.T) {
	tests := []struct {
		in  []byte
		out uint64
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, math.MaxUint64},
	}
	for _, test := range tests {
		out, err := Uint64(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Uint64(test.in)
		assert.NotNil(t, err)
	}
}
func TestUint32(t *testing.T) {
	tests := []struct {
		in  []byte
		out uint32
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff, 0xff, 0xff, 0xff}, math.MaxUint32},
	}
	for _, test := range tests {
		out, err := Uint32(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Uint32(test.in)
		assert.NotNil(t, err)
	}
}
func TestUint16(t *testing.T) {
	tests := []struct {
		in  []byte
		out uint16
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff, 0xff}, math.MaxUint16},
	}
	for _, test := range tests {
		out, err := Uint16(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Uint16(test.in)
		assert.NotNil(t, err)
	}
}
func TestUint8(t *testing.T) {
	tests := []struct {
		in  []byte
		out uint8
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff}, math.MaxUint8},
	}
	for _, test := range tests {
		out, err := Uint8(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Uint8(test.in)
		assert.NotNil(t, err)
	}
}
func TestPutInt64(t *testing.T) {
	tests := []struct {
		in  int64
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MinInt64, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{math.MaxInt64, []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, test := range tests {
		out := PutInt64(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestPutInt32(t *testing.T) {
	tests := []struct {
		in  int32
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MinInt32, []byte{0xff, 0xff, 0xff, 0xff}},
		{math.MaxInt32, []byte{0xfe, 0xff, 0xff, 0xff}},
	}
	for _, test := range tests {
		out := PutInt32(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestPutInt16(t *testing.T) {
	tests := []struct {
		in  int16
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MinInt16, []byte{0xff, 0xff}},
		{math.MaxInt16, []byte{0xfe, 0xff}},
	}
	for _, test := range tests {
		out := PutInt16(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestPutInt8(t *testing.T) {
	tests := []struct {
		in  int8
		out []byte
	}{
		{0, []byte{0x0}},
		{math.MinInt8, []byte{0xff}},
		{math.MaxInt8, []byte{0xfe}},
	}
	for _, test := range tests {
		out := PutInt8(test.in)
		assert.Equal(t, test.out, out)
	}
}
func TestInt64(t *testing.T) {
	tests := []struct {
		in  []byte
		out int64
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, math.MinInt64},
		{[]byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, math.MaxInt64},
	}
	for _, test := range tests {
		out, err := Int64(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Int64(test.in)
		assert.NotNil(t, err)
	}
}
func TestInt32(t *testing.T) {
	tests := []struct {
		in  []byte
		out int32
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff, 0xff, 0xff, 0xff}, math.MinInt32},
		{[]byte{0xfe, 0xff, 0xff, 0xff}, math.MaxInt32},
	}
	for _, test := range tests {
		out, err := Int32(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Int32(test.in)
		assert.NotNil(t, err)
	}
}
func TestInt16(t *testing.T) {
	tests := []struct {
		in  []byte
		out int16
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff, 0xff}, math.MinInt16},
		{[]byte{0xfe, 0xff}, math.MaxInt16},
	}
	for _, test := range tests {
		out, err := Int16(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Int16(test.in)
		assert.NotNil(t, err)
	}
}
func TestInt8(t *testing.T) {
	tests := []struct {
		in  []byte
		out int8
	}{
		{[]byte{0x0}, 0},
		{[]byte{0xff}, math.MinInt8},
		{[]byte{0xfe}, math.MaxInt8},
	}
	for _, test := range tests {
		out, err := Int8(test.in)
		assert.Nil(t, err)
		assert.Equal(t, test.out, out)
	}
	testErrors := []struct {
		in []byte
	}{
		{[]byte{}},
		{[]byte{0x1, 0xff}},
	}
	for _, test := range testErrors {
		_, err := Int8(test.in)
		assert.NotNil(t, err)
	}
}
