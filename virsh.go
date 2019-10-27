package dromaius

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"
	"golang.org/x/crypto/ssh/terminal"
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

func (v *VirshMachine) Status() error {
	return v.exec("domstate", false)
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
	case "domstate":
	default:
		return errCommandNotImplemented
	}

	var outerrs bytes.Buffer
	c := exec.Command(virshPath, "--connect=qemu:///system", cmd, v.name)
	if attachIO {
		fPTY, err := pty.Start(c)
		if err != nil {
			return err
		}
		defer fPTY.Close()

		// Set our stdin in raw mode
		oldStdinState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}
		defer terminal.Restore(int(os.Stdin.Fd()), oldStdinState)

		go io.Copy(fPTY, os.Stdin)
		io.Copy(os.Stdout, fPTY)

		return nil
	}

	c.Stdout = &outerrs
	c.Stderr = &outerrs
	err := c.Run()
	if err != nil {
		return fmt.Errorf(strings.TrimSpace(outerrs.String()))
	}

	fmt.Printf("%s", outerrs.String())

	return nil
}
