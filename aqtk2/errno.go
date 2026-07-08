package aqtk2

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
	case 102, 105:
		return "undefined phonetic symbol specified in the input string"
	case 103:
		return "prosodic data duration is negative"
	case 104:
		return "internal error (undefined delimiter code detected)"
	case 106:
		return "incorrect tag specification in the input string"
	case 107:
		return "tag length exceeds the limit (or closing '>' bracket is missing)"
	case 108:
		return "invalid value specified inside the tag"
	case 111:
		return "no data to speak"
	case 200:
		return "phonetic symbol string is too long"
	case 201:
		return "too many phonetic symbols in a single phrase"
	case 202, 204:
		return "phonetic symbol string is too long (internal buffer overflow 1)"
	case 203:
		return "insufficient heap memory"
	default:
		if e >= 1000 && e <= 1008 {
			return "invalid phont data"
		}
		return "unknown error"
	}
}
