package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	dromaius "github.com/mrngm/go-dromaius"
)

func runInteractive(hostname string) {
	buf := bufio.NewReader(os.Stdin)
	hn := strings.Split(hostname, ".")
	var lastInput string
	var err error

	for {
		fmt.Printf("dromaius (%v) > ", hn[0])
		if n := buf.Buffered(); n > 0 {
			fmt.Printf("(discarding %d bytes)\n", n)
			buf.Discard(n)
		}

		lastInput, err = buf.ReadString('\n')
		if err != nil && err == io.EOF {
			fmt.Printf("\nGoodbye!\n")
			return
		}

		if err != nil {
			fmt.Printf("Cannot read input correctly, exiting.\n")
			return
		}

		lastInput = strings.TrimSpace(lastInput)

		switch lastInput {
		case "start":
		case "shutdown":
		case "reboot":
		case "reset":
		case "destroy":
		case "console":
		case "status":
		case "exit":
			fmt.Printf("\nGoodbye!\n")
			return
		case "":
			continue
		case "help":
			printHelp()
			continue
		default:
			fmt.Printf("Expecting one command (start, shutdown, reboot, reset, destroy, console, status, exit, help)\n")
			continue
		}

		err = dromaius.RunCommand(hostname, lastInput)
		if err != nil && lastInput != "console" {
			fmt.Printf("Error when running command %q: %v\n", lastInput, err)
		}
	}
}

func printHelp() {
	fmt.Println("Available commands: start, shutdown, reboot, reset, destroy, console, status, exit, help")
	fmt.Println()
	fmt.Println("start           : start the domain")
	fmt.Println("shutdown, reboot: send ACPI signal to the domain to shutdown or reboot")
	fmt.Println("destroy, reset  : immediately turn off or reboot the domain")
	fmt.Println("console         : connect to the serial console of the domain")
	fmt.Println("status          : show the state of the domain")
	fmt.Println("exit            : exit the dromaius shell")
	fmt.Println("help            : show this help message")
}
