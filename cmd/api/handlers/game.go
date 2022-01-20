package handlers

import (
	"arcade-tools/cmd/api/models"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const TempFile = "/tmp/now-playing"

func NowPlaying(c echo.Context) error {
	contents, err := ioutil.ReadFile(TempFile)

	if errors.Is(err, os.ErrNotExist) || len(contents) == 0 {
		return c.NoContent(http.StatusNoContent)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewApiError(err))
	}

	items := strings.Split(string(contents), "\n")

	return c.JSON(http.StatusOK, struct {
		System   string `json:"system"`
		Emulator string `json:"emulator"`
		RomPath  string `json:"romPath"`
		RomName  string `json:"romName"`
		Command  string `json:"command"`
	}{
		System:   items[0],
		Emulator: items[1],
		RomPath:  items[2],
		RomName:  items[3],
		Command:  items[4],
	})
}
