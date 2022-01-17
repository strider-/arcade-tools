package handlers

import "github.com/labstack/echo"

func GetRelays(c echo.Context) error {
	return c.JSON(200, struct {
		Ok bool `json:"ok"`
	}{Ok: true})
}

func GetRelay(c echo.Context) error {
	return c.JSON(200, struct {
		Ok bool `json:"ok"`
	}{Ok: true})
}

func SetRelayState(c echo.Context) error {
	return c.JSON(200, struct {
		Ok bool `json:"ok"`
	}{Ok: true})
}
