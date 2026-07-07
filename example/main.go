package main

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/SladkyCitron/aquestalk-go/aqtk1"
	"github.com/SladkyCitron/resona/codec/wav"
	"github.com/SladkyCitron/resona/playback"
	_ "github.com/SladkyCitron/resona/playback/driver/oto"
)

const voice = "f1"

func getLibrary() string {
	switch runtime.GOOS {
	case "windows":
		return "AquesTalk.dll"
	case "darwin":
		return "libAquesTalk1-" + voice + ".dylib"
	case "linux":
		return "libAquesTalk.so"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	aq, err := aqtk1.New(getLibrary())
	if err != nil {
		panic(err)
	}

	wavBytes, err := aq.Synthe("こんにちわ。", 100)
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
