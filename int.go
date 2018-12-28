package binaryext

import errors "github.com/weathersource/go-errors"

var (
	ErrOverflow = errors.NewInvalidArgumentError("Byte count exceeds expectation.")
	ErrEmpty    = errors.NewInvalidArgumentError("No bytes provided.")
)

// putUint encodes a uint64 to []byte.
func putUint(x uint64, maxbytes int) []byte {

	buf := make([]byte, 0, maxbytes)

	for x >= 0x100 {
		buf = append(buf, byte(x))
		x >>= 8
	}
	buf = append(buf, byte(x))

	return buf
}

// getUint decodes a uint64 from []byte
func getUint(buf []byte, maxbytes int) (uint64, error) {
	var x uint64
	var s uint8
	for i, b := range buf {
		if i > maxbytes-1 && b > 0 {
			return 0, ErrOverflow
		}
		if i == len(buf)-1 {
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0xff) << s
		s += 8
	}
	return 0, ErrEmpty
}

// putInt encodes a int64 to []byte.
func putInt(x int64, maxbytes int) []byte {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return putUint(ux, maxbytes)
}

// getInt decodes a int64 from []byte
func getInt(buf []byte, maxbytes int) (int64, error) {
	ux, err := getUint(buf, maxbytes) // ok to continue in presence of error
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x, err
}

// PutUint64 encodes a uint64 to []byte.
func PutUint64(x uint64) []byte {
	return putUint(x, 8)
}

// PutUint32 encodes a uint32 to []byte.
func PutUint32(x uint32) []byte {
	return putUint(uint64(x), 4)
}

// PutUint16 encodes a uint16 to []byte.
func PutUint16(x uint16) []byte {
	return putUint(uint64(x), 2)
}

// PutUint8 encodes a uint8 to []byte.
func PutUint8(x uint8) []byte {
	return putUint(uint64(x), 1)
}

// Uint64 decodes a uint64 from []byte
func Uint64(buf []byte) (uint64, error) {
	return getUint(buf, 8)
}

// Uint32 decodes a uint32 from []byte
func Uint32(buf []byte) (uint32, error) {
	x, err := getUint(buf, 4)
	return uint32(x), err
}

// Uint16 decodes a uint16 from []byte
func Uint16(buf []byte) (uint16, error) {
	x, err := getUint(buf, 2)
	return uint16(x), err
}

// Uint8 decodes a uint8 from []byte
func Uint8(buf []byte) (uint8, error) {
	x, err := getUint(buf, 1)
	return uint8(x), err
}

// PutInt64 encodes a int64 to []byte.
func PutInt64(x int64) []byte {
	return putInt(x, 8)
}

// PutInt32 encodes a int32 to []byte.
func PutInt32(x int32) []byte {
	return putInt(int64(x), 4)
}

// PutInt16 encodes a int16 to []byte.
func PutInt16(x int16) []byte {
	return putInt(int64(x), 2)
}

// PutInt8 encodes a int8 to []byte.
func PutInt8(x int8) []byte {
	return putInt(int64(x), 1)
}

// Int64 decodes a int64 from []byte
func Int64(buf []byte) (int64, error) {
	return getInt(buf, 8)
}

// Int32 decodes a int32 from []byte
func Int32(buf []byte) (int32, error) {
	x, err := getInt(buf, 4)
	return int32(x), err
}

// Int16 decodes a int16 from []byte
func Int16(buf []byte) (int16, error) {
	x, err := getInt(buf, 2)
	return int16(x), err
}

// Int8 decodes a int8 from []byte
func Int8(buf []byte) (int8, error) {
	x, err := getInt(buf, 1)
	return int8(x), err
}
