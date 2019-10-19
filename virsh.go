package dromaius

import (
	"errors"
	"os/exec"
)

var (
	virshPath                = "/usr/bin/virsh"
	errCommandNotImplemented = errors.New("command not implemented")
)

type VirshMachine struct {
	name string
}

func NewMachine(hostname string) *VirshMachine {
	return &VirshMachine{
		name: hostname,
	}
}

func (v *VirshMachine) Start() error {
	return errCommandNotImplemented
}

func (v *VirshMachine) Shutdown() error {
	return errCommandNotImplemented
}

func (v *VirshMachine) Reset() error {
	return errCommandNotImplemented
}

func (v *VirshMachine) Reboot() error {
	return errCommandNotImplemented
}

func (v *VirshMachine) Destroy() error {
	return errCommandNotImplemented
}

func (v *VirshMachine) Console() error {
	return errCommandNotImplemented
}
