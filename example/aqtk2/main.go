package main

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/SladkyCitron/aquestalk-go/aqtk2"
	"github.com/SladkyCitron/resona/codec/wav"
	"github.com/SladkyCitron/resona/playback"
	_ "github.com/SladkyCitron/resona/playback/driver/oto"
)

func getLibrary() string {
	switch runtime.GOOS {
	case "windows":
		return "AquesTalk2.dll"
	case "darwin":
		return "libAquesTalk2.dylib"
	case "linux":
		return "libAquesTalk2.so"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	aq, err := aqtk2.New(getLibrary())
	if err != nil {
		panic(err)
	}

	wavBytes, err := aq.Synthe("こんにちわ。", 100, nil)
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
