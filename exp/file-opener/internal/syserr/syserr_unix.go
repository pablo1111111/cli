package syserr

import "syscall"

const (
	EMFILE = syscall.Errno(0x18)
)
