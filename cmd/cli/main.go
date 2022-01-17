package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	// BCM #, not physical pin #
	Relay1Pin = 18
	Relay2Pin = 17
)

func main() {
	gpioPin, state, err := parseArgs()
	if err != nil {
		fmt.Println("invalid arguments, use --help for usage.")
		os.Exit(0)
	}

	err = rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()

	pin := rpio.Pin(gpioPin)
	pin.Output()

	pin.Write(state)
}

func parseArgs() (rpio.Pin, rpio.State, error) {
	var relay int
	var on bool
	var off bool

	flag.IntVar(&relay, "relay", 0, "Relay to change (1 or 2)")
	flag.BoolVar(&on, "on", false, "Turn the relay on")
	flag.BoolVar(&off, "off", true, "Turn the relay off")
	flag.Parse()

	if relay != 1 && relay != 2 {
		return 0, 0, errors.New("invalid relay and/or state")
	}

	var pin uint8
	if relay == 1 {
		pin = Relay1Pin
	} else {
		pin = Relay2Pin
	}

	var state uint8
	if on {
		state = 0
	} else {
		state = 1
	}

	return rpio.Pin(pin), rpio.State(state), nil
}
