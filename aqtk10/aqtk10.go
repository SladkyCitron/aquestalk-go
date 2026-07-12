// Package aqtk10 provides Go bindings for AquesTalk10.
package aqtk10

import (
	"unsafe"

	"github.com/SladkyCitron/aquestalk-go/internal/lib"
	"github.com/ebitengine/purego"
)

// VoiceBase represents the base voice unit.
type VoiceBase int32

const (
	VoiceBaseF1E VoiceBase = 0
	VoiceBaseF2E VoiceBase = 1
	VoiceBaseM1E VoiceBase = 2
)

// Voice represents voice quality parameters.
type Voice struct {
	// Bas is the base voice unit.
	Bas VoiceBase

	// Spd is the speaking speed. (range: 50-300, default: 100)
	Spd int32

	// Vol is the volume. (range: 0-300, default: 100)
	Vol int32

	// Pit is the pitch. (range: 20-200, default depends on base voice unit)
	Pit int32

	// Acc is the accent. (range: 0-200, default depends on base voice unit)
	Acc int32

	// Lmd is the 1st pitch parameter. (range: 0-200, default: 100)
	Lmd int32

	// Fsc is the 2nd pitch parameter. (range: 50-200, default: 100)
	Fsc int32
}

var (
	VoiceF1 = Voice{Bas: VoiceBaseF1E, Spd: 100, Vol: 100, Pit: 100, Acc: 100, Lmd: 100, Fsc: 100}
	VoiceF2 = Voice{Bas: VoiceBaseF2E, Spd: 100, Vol: 100, Pit: 77, Acc: 150, Lmd: 100, Fsc: 100}
	VoiceF3 = Voice{Bas: VoiceBaseF1E, Spd: 80, Vol: 100, Pit: 100, Acc: 100, Lmd: 61, Fsc: 148}
	VoiceM1 = Voice{Bas: VoiceBaseM1E, Spd: 100, Vol: 100, Pit: 30, Acc: 100, Lmd: 100, Fsc: 100}
	VoiceM2 = Voice{Bas: VoiceBaseM1E, Spd: 105, Vol: 100, Pit: 45, Acc: 130, Lmd: 120, Fsc: 100}
	VoiceR1 = Voice{Bas: VoiceBaseM1E, Spd: 100, Vol: 100, Pit: 30, Acc: 20, Lmd: 190, Fsc: 100}
	VoiceR2 = Voice{Bas: VoiceBaseF2E, Spd: 70, Vol: 100, Pit: 50, Acc: 50, Lmd: 50, Fsc: 180}
)

// AquesTalk represents the AquesTalk10 library and provides methods to interact with it.
type AquesTalk struct {
	_AquesTalk_Synthe_Utf8 func(pParam unsafe.Pointer, koe string, size *int32) *byte
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
func (aq *AquesTalk) Synthe(voice Voice, koe string) ([]byte, error) {
	var size int32
	wavPtr := aq._AquesTalk_Synthe_Utf8(unsafe.Pointer(&voice), koe, &size)
	if wavPtr == nil {
		return nil, Errno(size)
	}
	defer aq._AquesTalk_FreeWave(wavPtr)

	wavBytes := unsafe.Slice(wavPtr, size)
	newWavBytes := make([]byte, size)
	copy(newWavBytes, wavBytes)

	return newWavBytes, nil
}

// SetDevKey sets the development license key.
func (aq *AquesTalk) SetDevKey(key string) (ok bool) {
	ret := aq._AquesTalk_SetDevKey(key)
	return ret == 0
}

// SetUsrKey sets the user license key.
func (aq *AquesTalk) SetUsrKey(key string) (ok bool) {
	ret := aq._AquesTalk_SetUsrKey(key)
	return ret == 0
}
