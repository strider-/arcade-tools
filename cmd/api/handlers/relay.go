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

	return c.JSON(200, result)
}

func GetRelay(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewApiError(err))
	}

	rly, err := relay.ParseRelay(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewApiError(err))
	}

	state := relay.GetState(rly)

	return c.JSON(200, models.NewRelay(rly, state))
}

func SetRelayState(c echo.Context) error {
	return c.JSON(200, struct {
		Ok bool `json:"ok"`
	}{Ok: true})
}
