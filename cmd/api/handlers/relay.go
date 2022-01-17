package handlers

import (
	"arcade-tools/cmd/api/models"
	"arcade-tools/internal/relay"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetRelays(c echo.Context) error {
	r1 := relay.GetState(relay.Relay1)
	r2 := relay.GetState(relay.Relay2)

	var result = []models.Relay{
		models.NewRelay(relay.Relay1, r1),
		models.NewRelay(relay.Relay2, r2),
	}

	return c.JSON(http.StatusOK, result)
}

func GetRelay(c echo.Context) error {
	rly, err := getRelayFromId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewApiError(err))
	}

	state := relay.GetState(rly)

	return c.JSON(http.StatusOK, models.NewRelay(rly, state))
}

func SetRelayState(c echo.Context) error {
	rly, err := getRelayFromId(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewApiError(err))
	}

	var req struct {
		State uint8 `json:"state"`
	}

	if err = c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewApiError(err))
	}

	state, err := relay.ParseUInt8State(req.State)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewApiError(err))
	}

	relay.SetState(rly, models.ToActiveLowState(state))

	return c.NoContent(http.StatusNoContent)
}

func getRelayFromId(c echo.Context) (relay.Relay, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}

	rly, err := relay.ParseRelay(id)
	if err != nil {
		return 0, err
	}

	return rly, nil
}
