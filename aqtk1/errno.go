package aqtk1

import "fmt"

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
	case 105:
		return "undefined phonetic symbol specified in the input string"
	case 106:
		return "incorrect tag specification in the input string"
	case 107:
		return "tag length exceeds the limit (or closing bracket [>] is missing)"
	case 108:
		return "invalid value specified inside the tag"
	case 200, 202, 204:
		return "phonetic symbol string is too long"
	case 201:
		return "too many phonetic symbols within a single phrase"
	case 203:
		return "insufficient heap memory"
	default:
		return "unknown error"
	}
}
