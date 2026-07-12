package aqtk10

import "fmt"

// Errno represents an error code.
type Errno int32

func (e Errno) Error() string {
	return fmt.Sprintf("%s (errno %d)", e.message(), int32(e))
}

func (e Errno) message() string {
	switch e {
	case 100:
		return "other error"
	case 101:
		return "out of memory"
	case 103:
		return "phonetic symbol string specification error"
	case 104:
		return "no valid reading found in the phonetic symbol string"
	case 105:
		return "undefined phonetic symbol specified in the input string"
	case 106:
		return "invalid tag specification in the input string"
	case 107:
		return "tag length exceeds the limit (or closing '>' bracket is missing)"
	case 108:
		return "invalid value specified inside the tag"
	case 120:
		return "phonetic symbol string is too long"
	case 121:
		return "too many phonetic symbols in a single phrase"
	case 122:
		return "phonetic symbol string is too long (internal buffer overflow 1)"
	default:
		return "unknown error"
	}
}
