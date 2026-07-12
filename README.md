# 🗣️ aquestalk-go

[![Go Reference](https://pkg.go.dev/badge/github.com/SladkyCitron/aquestalk-go.svg)](https://pkg.go.dev/github.com/SladkyCitron/aquestalk-go) [![CI (Go)](https://github.com/SladkyCitron/aquestalk-go/actions/workflows/ci.yml/badge.svg)](https://github.com/SladkyCitron/aquestalk-go/actions/workflows/ci.yml) [![GitHub license](https://img.shields.io/github/license/SladkyCitron/aquestalk-go)](LICENSE) ![Made in Slovakia](https://raw.githubusercontent.com/pedromxavier/flag-badges/refs/heads/main/badges/SK.svg)

**aquestalk-go** provides unofficial Go bindings for the [AquesTalk](https://www.a-quest.com/products/aquestalk.html) Japanese TTS engine using [purego](https://github.com/ebitengine/purego).

## ✨ Features

* Pure Go bindings, no cgo required
* Dynamic library loading
* Support for:
  * AquesTalk1
  * AquesTalk2
  * AquesTalk10

## 🚀 Getting Started

Install using this command:

```bash
go get github.com/SladkyCitron/aquestalk-go
```

> [!NOTE]
> You must independently obtain the official [AquesTalk dynamic libraries](https://www.a-quest.com/download.html) from AQUEST.
> **They are not included** in this repository.

## 📚 Example

```go
package main

import (
    "os"

    "github.com/SladkyCitron/aquestalk-go/aqtk10"
)

func main() {
    aq, err := aqtk10.New("AquesTalk.dll")
    if err != nil {
        panic(err)
    }

    wavBytes, err := aq.Synthe(aqtk10.VoiceF1, "こんにちわ。")
    if err != nil {
        panic(err)
    }

    if err := os.WriteFile("output.wav", wavBytes, 0644); err != nil {
        panic(err)
    }
}
```

Complete code examples can be found in the [example/](example/) directory.

## 📚 Documentation

All documentation is available at [pkg.go.dev/github.com/SladkyCitron/aquestalk-go](https://pkg.go.dev/github.com/SladkyCitron/aquestalk-go).

## ⚖️ License

Copyright © 2026 SladkyCitron

Licensed under the **MIT License** (see [LICENSE](LICENSE)) - free to use, fork, remix, and share!
