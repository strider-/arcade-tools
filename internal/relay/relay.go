package relay

import (
	"errors"

	"github.com/stianeikeland/go-rpio/v4"
)

type RelayState uint8
type Relay uint8

const (
	// The relay is active low, so 0 turns it on, 1 turns it off.
	On  RelayState = 0
	Off RelayState = 1

	Relay1 Relay = 18
	Relay2 Relay = 17
)

var (
	ErrInvalidRelay      = errors.New("invalid relay")
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

func GetState(r Relay) RelayState {
	pin, cleanUp := getPin(r)
	defer cleanUp()

	return RelayState(pin.Read())
}

func SetState(r Relay, s RelayState) {
	pin, cleanUp := getPin(r)
	defer cleanUp()

	pin.Write(rpio.State(s))
}

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

func ParseState(relayState bool) RelayState {
	switch relayState {
	case true:
		return On
	default:
		return Off
	}
}
