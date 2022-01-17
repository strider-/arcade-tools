package main

import (
	"arcade-tools/internal/relay"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(flag.Args()) == 0 {
		fmt.Println("missing arguments, use --help for usage")
		os.Exit(0)
	}

	rly, state, err := parseArgs()

	if err != nil {
		fmt.Println("invalid arguments, use --help for usage.")
		os.Exit(1)
	}

	relay.SetState(rly, state)
}

func parseArgs() (relay.Relay, relay.RelayState, error) {
	var r int
	var on bool
	var off bool

	flag.IntVar(&r, "relay", 0, "Relay to change (1 or 2)")
	flag.BoolVar(&on, "on", false, "Turn the relay on")
	flag.BoolVar(&off, "off", true, "Turn the relay off")
	flag.Parse()

	pin, err := relay.ParseRelay(r)
	if err != nil {
		return 0, 0, err
	}

	state := relay.ParseState(on)
	return pin, state, nil
}
