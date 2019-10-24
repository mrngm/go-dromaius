package dromaius

var knownHosts = map[string]bool{}

func AddHost(hostname string) {
	knownHosts[hostname] = true
}

func KnownHost(hostname string) bool {
	_, ok := knownHosts[hostname]
	return ok
}
