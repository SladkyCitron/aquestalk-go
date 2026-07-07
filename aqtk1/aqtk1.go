// Package aqtk1 provides Go bindings for AquesTalk1.
package aqtk1

import (
	"unsafe"

	"github.com/SladkyCitron/aquestalk-go/internal/lib"
	"github.com/ebitengine/purego"
)

// AquesTalk represents the AquesTalk library and provides methods to interact with it.
type AquesTalk struct {
	_AquesTalk_Synthe_Utf8 func(koe string, speed int32, size *int32) *byte
	_AquesTalk_FreeWave    func(wav *byte)
	_AquesTalk_SetDevKey   func(key string) int32
	_AquesTalk_SetUsrKey   func(key string) int32
}

// New initializes a new instance of [AquesTalk] and loads the dynamic library specified.
func New(path string) (*AquesTalk, error) {
	h, err := lib.Dlopen(path)
	if err != nil {
		return nil, err
	}

	aq := &AquesTalk{}

	purego.RegisterLibFunc(&aq._AquesTalk_Synthe_Utf8, h, "AquesTalk_Synthe_Utf8")
	purego.RegisterLibFunc(&aq._AquesTalk_FreeWave, h, "AquesTalk_FreeWave")
	purego.RegisterLibFunc(&aq._AquesTalk_SetDevKey, h, "AquesTalk_SetDevKey")
	purego.RegisterLibFunc(&aq._AquesTalk_SetUsrKey, h, "AquesTalk_SetUsrKey")

	return aq, nil
}

// Synthe synthesizes speech from the given phonetic string (koe) at the specified speed.
// It returns the synthesized speech as a byte slice in WAV format or an error if the synthesis fails.
func (aq *AquesTalk) Synthe(koe string, speed int) ([]byte, error) {
	var size int32
	wavPtr := aq._AquesTalk_Synthe_Utf8(koe, int32(speed), &size)
	if wavPtr == nil {
		return nil, Errno(size)
	}
	defer aq._AquesTalk_FreeWave(wavPtr)

	wavBytes := unsafe.Slice(wavPtr, size)
	newWavBytes := make([]byte, size)
	copy(newWavBytes, wavBytes)

	return newWavBytes, nil
}

// SetDevKey sets the developer key.
func (aq *AquesTalk) SetDevKey(key string) (ok bool) {
	ret := aq._AquesTalk_SetDevKey(key)
	return ret == 0
}

// SetUsrKey sets the user key.
func (aq *AquesTalk) SetUsrKey(key string) (ok bool) {
	ret := aq._AquesTalk_SetUsrKey(key)
	return ret == 0
}
