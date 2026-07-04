//go:build windows

package lib

import "golang.org/x/sys/windows"

func Dlopen(path string) (uintptr, error) {
	h, err := windows.LoadDLL(path)
	return uintptr(h.Handle), err
}
