//go:build !darwin

package shim

import "os/exec"

func detachHelperProcess(cmd *exec.Cmd) {
}
