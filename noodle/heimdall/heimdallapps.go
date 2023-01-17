package heimdall

import (
	"errors"
	"io"
	"net/http"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
)

var (
	NewHeimdall = NewHeimdallImpl
)

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name Heimdall

type Heimdall interface {
	UpdateFromServer() error
	FindApps(search string) ([]models.ApplicationTemplate, error)
}

type HeimdallImpl struct {
	database database.Database
}

// FindApp implements Heimdall
func (i *HeimdallImpl) FindApps(search string) ([]models.ApplicationTemplate, error) {
	table := database.NewAppTemplateTable(i.database)
	return table.Search(search)
}

// UpdateFromServer implements Heimdall
func (i *HeimdallImpl) UpdateFromServer() error {
	response, err := http.Get("https://appslist.heimdall.site/list.json")

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

	}

	return nil

}

func NewHeimdallImpl(database database.Database) Heimdall {
	return &HeimdallImpl{
		database: database,
	}
}
