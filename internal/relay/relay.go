package relay

import (
	"errors"

	"github.com/stianeikeland/go-rpio/v4"
)

type RelayState uint8
type Relay uint8

const (
	// The relay is active low, so 0 turns it on
	On RelayState = 0
	// The relay is active low, so 1 turns it off
	Off RelayState = 1

	// Relay1 represents GPIO 18 (BCM not physical)
	Relay1 Relay = 18
	// Relay2 represents GPIO 17 (BCM not physical)
	Relay2 Relay = 17
)

var (
	// Returned when calling relay.ParseRelay with a value that's not 1 or 2
	ErrInvalidRelay = errors.New("invalid relay")
	// Returned when calling relay.ParseUInt8State with a value that's not 1 or 2
	ErrInvalidRelayState = errors.New("invalid relay state")
)

func getPin(r Relay) (rpio.Pin, func() error) {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	pin := rpio.Pin(r)
	pin.Output()

	return pin, rpio.Close
}

// Gets the current RelayState of the given Relay
func GetState(r Relay) RelayState {
	pin, cleanUp := getPin(r)
	defer cleanUp()

	return RelayState(pin.Read())
}

// Sets the RelayState of the given Relay
func SetState(r Relay, s RelayState) {
	pin, cleanUp := getPin(r)
	defer cleanUp()

	pin.Write(rpio.State(s))
}

// Parses an integer (1 or 2) into either Relay1 or Relay2. Returns ErrInvalidRelay with any other value
func ParseRelay(relayNumber int) (Relay, error) {
	switch relayNumber {
	case 1:
		return Relay1, nil
	case 2:
		return Relay2, nil
	default:
		return 0, ErrInvalidRelay
	}
}

// Parses a boolean into a RelayState. true is relay.On, false is relay.Off
func ParseBoolState(relayState bool) RelayState {
	switch relayState {
	case true:
		return On
	default:
		return Off
	}
}

// Parses a uint8 into a RelayState. 0 is relay.On, 1 is relay.Off. Returns ErrInvalidRelayState with any other value
func ParseUInt8State(relayState uint8) (RelayState, error) {
	switch relayState {
	case 0:
		return On, nil
	case 1:
		return Off, nil
	default:
		return 0, ErrInvalidRelayState
	}
}
