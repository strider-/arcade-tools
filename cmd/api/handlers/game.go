package handlers

import (
	"arcade-tools/cmd/api/models"
	"arcade-tools/internal/utils"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	ps "github.com/mitchellh/go-ps"
)

var (
	ErrFailedToListProcesses = errors.New("could not enumerate running processes")
	ErrFailedToKillRetroArch = errors.New("could not terminate the retroarch process")
	ErrRetroArchIsNotRunning = errors.New("retroarch is not running")
)

func NowPlaying(c echo.Context) error {
	npFile := utils.GetEnvOrDefault("AC_NOW_PLAYING_FILE", "/tmp/now-playing")

	contents, err := ioutil.ReadFile(npFile)

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

func KillRetroarch(c echo.Context) error {
	processList, err := ps.Processes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewApiError(ErrFailedToListProcesses))
	}

	for _, p := range processList {
		if p.Executable() == "retroarch" {
			osp := os.Process{Pid: p.Pid()}
			err = osp.Kill()

			if err != nil {
				return c.JSON(http.StatusInternalServerError, models.NewApiError(ErrFailedToKillRetroArch))
			}

			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusBadRequest, models.NewApiError(ErrRetroArchIsNotRunning))
}
