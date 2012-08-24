package xattr

import (
	"syscall"
	"unsafe"
)

func getxattr(path string, name string, value *byte, size int) (result int, err error) {
	r0, _, e1 := syscall.Syscall6(syscall.SYS_GETXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), uintptr(unsafe.Pointer(value)), uintptr(size), 0, 0)
    //var err error = nil
    if e1 != 0 {
        err = e1
    }
	return int(r0), err

}

func listxattr(path string, namebuf *byte, size int) (result int, err error) {
	r0, _, e1 := syscall.Syscall(syscall.SYS_LISTXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(namebuf)), uintptr(size))
    //var err error = nil
    if e1 != 0 {
        err = e1
    }
	return int(r0), err
}

func setxattr(path string, name string, value *byte, size int) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), uintptr(unsafe.Pointer(value)), uintptr(size), 0, 0)
    //var err error = nil
    if e1 != 0 {
        err = e1
    }
	return e1
}

func removexattr(path string, name string) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_REMOVEXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), 0)
    //var err error = nil
    if e1 != 0 {
        err = e1
    }
	return e1
}
