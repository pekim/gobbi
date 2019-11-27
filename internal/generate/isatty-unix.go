// +build linux aix

package generate

import (
	"golang.org/x/sys/unix"
	"os"
)

func isTerminal() bool {
	fd := os.Stdout.Fd()
	_, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)

	return err == nil
}
