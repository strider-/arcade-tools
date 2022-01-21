package models

import (
	"arcade-tools/internal/relay"
	"arcade-tools/internal/utils"
)

type Relay struct {
	Id    uint8            `json:"id"`
	Name  string           `json:"name"`
	State relay.RelayState `json:"state"`
}

func toDTO(r relay.Relay) (uint8, string) {
	switch r {
	case relay.Relay1:
		return 1, utils.GetEnvOrDefault("AC_RELAY1_NAME", "lighting")
	case relay.Relay2:
		return 2, utils.GetEnvOrDefault("AC_RELAY2_NAME", "monitor")
	default:
		return 0, "unknown"
	}
}

func ToActiveLowState(s relay.RelayState) relay.RelayState {
	// since the relay is active low, using a
	// binary NOT to toggle the 1 bit, to ensure clients
	// get and send 1 for on and 0 for off.

	return s ^ 1
}

func NewRelay(r relay.Relay, s relay.RelayState) Relay {
	id, name := toDTO(r)
	return Relay{
		Id:    id,
		Name:  name,
		State: ToActiveLowState(s),
	}
}
