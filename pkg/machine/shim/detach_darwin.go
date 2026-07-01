//go:build darwin

package shim

import (
	"os/exec"
	"syscall"
)

func detachHelperProcess(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}
