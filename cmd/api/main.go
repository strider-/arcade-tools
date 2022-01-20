package main

import (
	"arcade-tools/cmd/api/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	v1 := e.Group("/api/v1")

	v1.GET("/now-playing", handlers.NowPlaying)

	v1.GET("/relays", handlers.GetRelays)
	v1.GET("/relay/:id", handlers.GetRelay)

	v1.POST("/relay/:id", handlers.SetRelayState)

	e.Logger.Fatal(e.Start(":4380"))
}
