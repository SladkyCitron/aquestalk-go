package main

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/SladkyCitron/aquestalk-go/aqtk10"
	"github.com/SladkyCitron/resona/codec/wav"
	"github.com/SladkyCitron/resona/playback"
	_ "github.com/SladkyCitron/resona/playback/driver/oto"
)

func getLibrary() string {
	switch runtime.GOOS {
	case "windows":
		return "AquesTalk.dll"
	case "darwin":
		return "libAquesTalk10.dylib"
	case "linux":
		return "libAquesTalk10.so"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	aq, err := aqtk10.New(getLibrary())
	if err != nil {
		panic(err)
	}

	wavBytes, err := aq.Synthe(aqtk10.VoiceF1, "こんにちわ。")
	if err != nil {
		panic(err)
	}

	deco, err := wav.NewDecoder(bytes.NewReader(wavBytes))
	if err != nil {
		panic(err)
	}

	playbackCtx, err := playback.NewContext(deco.Format())
	if err != nil {
		panic(err)
	}

	player := playbackCtx.NewPlayer(deco)
	player.PlayAndWait()
}
