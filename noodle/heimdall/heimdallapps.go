package heimdall

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/options"
	"github.com/mwinters-stuff/noodle/server/models"
)

var (
	NewHeimdall = NewHeimdallImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --case underscore --name Heimdall
type Heimdall interface {
	UpdateFromServer() error
}

type HeimdallImpl struct {
	database database.Database
	options  options.NoodleOptions
}

// UpdateFromServer implements Heimdall
func (i *HeimdallImpl) UpdateFromServer() error {
	response, err := http.Get(i.options.HeimdallListJsonURL)

	if err != nil {
		Logger.Error().Msgf("UpdateFromServer: %s", err.Error())
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	body, _ := io.ReadAll(response.Body)
	data := models.AppList{}
	err = data.UnmarshalBinary(body)
	if err != nil {
		Logger.Error().Msgf("UpdateFromServer: %s", err.Error())
		return err
	}

	table := database.NewAppTemplateTable(i.database)

	for _, app := range data.Apps {
		found, err := table.Exists(app.Appid)
		if err != nil {
			Logger.Error().Msgf("UpdateFromServer: %s", err.Error())
			return err
		}

		if found {
			Logger.Info().Msgf("Update: %#v", *app)
			err = table.Update(*app)
		} else {
			Logger.Info().Msgf("Insert: %#v", *app)
			err = table.Insert(*app)
		}
		if err != nil {
			Logger.Error().Msgf("UpdateFromServer: %s", err.Error())
			return err
		}

		// download icon and save.
		url, _ := url.JoinPath(i.options.HeimdallIconBaseURL, app.Icon)
		response, err = http.Get(url)
		if err != nil {
			Logger.Error().Msgf("Download Icon: %s", err.Error())
			return err
		} else if response.StatusCode != 200 {
			Logger.Error().Msgf("Download Icon: %s", response.Status)
			return fmt.Errorf("download icon failed: %s", response.Status)
		} else {
			data, _ := io.ReadAll(response.Body)
			err = os.WriteFile(path.Join(i.options.IconSavePath, app.Icon), data, 0644)
			if err != nil {
				Logger.Error().Msgf("Download Icon Write: %s", err.Error())
				return err
			}
		}

	}

	return nil

}

func NewHeimdallImpl(database database.Database, options options.NoodleOptions) Heimdall {
	return &HeimdallImpl{
		database: database,
		options:  options,
	}
}
