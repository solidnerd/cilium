// mksyscall.pl chmod_darwin.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package sysx

import (
	"syscall"
	"unsafe"
)

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := syscall.Syscall6(SYS_FCHMODAT, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	use(unsafe.Pointer(_p0))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
