// +build darwin freebsd openbsd netbsd dragonfly

package generate

import (
	"os"
)

func isTerminal() bool {
	fd := os.Stdout.Fd()
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)

	return err == 0
}
