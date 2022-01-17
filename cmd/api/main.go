package main

import (
	"github.com/labstack/echo"
)

func placeHolder(c echo.Context) error {
	c.JSON(200, struct {
		Ok bool `json:"ok"`
	}{Ok: true})
	return nil
}

func main() {
	e := echo.New()

	v1 := e.Group("/api/v1")

	v1.GET("/relays", placeHolder)
	v1.GET("/relay/:id", placeHolder)

	v1.POST("/relay/:id", placeHolder)

	e.Logger.Fatal(e.Start(":4380"))
}
