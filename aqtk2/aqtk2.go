// Package aqtk2 provides Go bindings for AquesTalk2.
package aqtk2

import (
	"unsafe"

	"github.com/SladkyCitron/aquestalk-go/internal/lib"
	"github.com/ebitengine/purego"
)

// AquesTalk represents the AquesTalk2 library and provides methods to interact with it.
type AquesTalk struct {
	_AquesTalk_Synthe_Utf8 func(koe string, iSpeed int32, pSize *int32, phontDat unsafe.Pointer) *byte
	_AquesTalk_FreeWave    func(wav *byte)
}

// New initializes a new instance of [AquesTalk] and loads the dynamic library specified.
func New(path string) (*AquesTalk, error) {
	h, err := lib.Dlopen(path)
	if err != nil {
		return nil, err
	}

	aq := &AquesTalk{}

	purego.RegisterLibFunc(&aq._AquesTalk_Synthe_Utf8, h, "AquesTalk2_Synthe_Utf8")
	purego.RegisterLibFunc(&aq._AquesTalk_FreeWave, h, "AquesTalk2_FreeWave")

	return aq, nil
}

// Synthe synthesizes speech from the given phonetic string (koe) and speed, using optional phont data.
// It returns the synthesized speech as a byte slice in WAV format or an error if the synthesis fails.
func (aq *AquesTalk) Synthe(koe string, speed int, phontDat []byte) ([]byte, error) {
	var phontDatPtr unsafe.Pointer
	if len(phontDat) > 0 {
		phontDatPtr = unsafe.Pointer(&phontDat[0])
	}

	var size int32
	wavPtr := aq._AquesTalk_Synthe_Utf8(koe, int32(speed), &size, phontDatPtr)
	if wavPtr == nil {
		return nil, Errno(size)
	}
	defer aq._AquesTalk_FreeWave(wavPtr)

	wavBytes := unsafe.Slice(wavPtr, size)
	newWavBytes := make([]byte, size)
	copy(newWavBytes, wavBytes)

	return newWavBytes, nil
}
