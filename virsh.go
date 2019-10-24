package dromaius

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var (
	virshPath = "/usr/bin/virsh"
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
	return v.exec("start", false)
}

func (v *VirshMachine) Shutdown() error {
	return v.exec("shutdown", false)
}

func (v *VirshMachine) Reset() error {
	return v.exec("reset", false)
}

func (v *VirshMachine) Reboot() error {
	return v.exec("reboot", false)
}

func (v *VirshMachine) Destroy() error {
	return v.exec("destroy", false)
}

func (v *VirshMachine) Console() error {
	return v.exec("console", true)
}

func (v *VirshMachine) exec(cmd string, attachIO bool) error {
	switch cmd {
	case "start":
	case "shutdown":
	case "reset":
	case "reboot":
	case "destroy":
	case "console":
	default:
		return errCommandNotImplemented
	}

	var outerrs bytes.Buffer
	c := exec.Command(virshPath, cmd, v.name)
	if attachIO {
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
	} else {
		c.Stdout = &outerrs
		c.Stderr = &outerrs
	}
	err := c.Run()
	if err != nil {
		fmt.Printf("stdout/err: %q\n", outerrs.String())
		return err
	}

	return nil
}
