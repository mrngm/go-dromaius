package dromaius

var (
	// AllowedCommands define which commands a user may execute.
	AllowedCommands = map[string]bool{
		"start":    true,
		"reboot":   true,
		"shutdown": true,
		"destroy":  true,
		"reset":    true,
		"console":  true,
	}
)

func RunCommand(hostname, cmd string) error {
	if allowed, ok := AllowedCommands[cmd]; !ok || (ok && !allowed) {
		// Don't allow commands we do not recognize, or that we've explicitly disabled
		return errCommandNotAllowed
	}
	if !KnownHost(hostname) {
		return errHostNotRecognized
	}

	machine := NewMachine(hostname)
	switch cmd {
	case "start":
		return machine.Start()
	case "reboot":
		return machine.Reboot()
	case "shutdown":
		return machine.Shutdown()
	case "destroy":
		return machine.Destroy()
	case "reset":
		return machine.Reset()
	case "console":
		return machine.Console()
	default:
		return errCommandNotAllowed
	}
}
