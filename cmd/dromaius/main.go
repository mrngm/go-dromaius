package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	dromaius "github.com/mrngm/go-dromaius"
)

var (
	hostName = flag.String("host", "", "Access this hostname")
	cmd      = flag.String("cmd", "", "Run this command (start, shutdown, reset, reboot, destroy)")
	config   = flag.String("config", "dromaius.json", "use this configuration file (list of known hosts)")
)

type cfg struct {
	AllowedHosts []string
}

func main() {
	flag.Parse()
	var runningConfig cfg

	if *hostName == "" || *cmd == "" {
		fmt.Printf("Usage: dromaius -host HOST -cmd CMD")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *config != "" {
		c, err := ioutil.ReadFile(*config)
		if err != nil {
			fmt.Printf("Error reading configuration: %v\n", err)
			os.Exit(3)
		}
		if err := json.Unmarshal(c, &runningConfig); err != nil {
			fmt.Printf("Error parsing JSON config: %v\n", err)
			os.Exit(4)
		}
	}

	os.Clearenv()

	for _, host := range runningConfig.AllowedHosts {
		dromaius.AddHost(host)
	}

	if *cmd == "interactive" {
		runInteractive(*hostName)
		return
	}

	err := dromaius.RunCommand(*hostName, *cmd)
	if err != nil {
		fmt.Printf("Error when running command %q on host %q: %v\n", *cmd, *hostName, err)
		os.Exit(2)
	}
}
