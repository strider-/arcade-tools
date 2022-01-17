package models

import "arcade-tools/internal/relay"

type Relay struct {
	Id    relay.Relay      `json:"id"`
	Name  string           `json:"name"`
	State relay.RelayState `json:"state"`
}

func getName(r relay.Relay) string {
	switch r {
	case relay.Relay1:
		return "lighting"
	case relay.Relay2:
		return "monitor"
	default:
		return "unknown"
	}
}

func NewRelay(r relay.Relay, s relay.RelayState) Relay {
	return Relay{
		Id:    r,
		Name:  getName(r),
		State: s,
	}
}
