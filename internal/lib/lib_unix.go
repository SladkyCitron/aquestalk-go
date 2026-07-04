//go:build !windows

package lib

import "github.com/ebitengine/purego"

func Dlopen(path string) (uintptr, error) {
	return purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
