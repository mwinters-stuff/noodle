package heimdall

import (
	"io"
	"net/http"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
)

type Heimdall interface {
	UpdateFromServer() error
	FindApps(search string) ([]jsontypes.App, error)
}

type HeimdallImpl struct {
	database database.Database
}

// FindApp implements Heimdall
func (i *HeimdallImpl) FindApps(search string) ([]jsontypes.App, error) {
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
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		Logger.Error().Msgf("UpdateFromServer: %s", err.Error())
		return err
	}

	data, err := jsontypes.UnmarshalAppList(body)
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
			err = table.Update(app)
		} else {
			err = table.Insert(app)
		}
		if err != nil {
			Logger.Error().Msgf("UpdateFromServer: %s", err.Error())
			return err
		}

	}

	return nil

}

func NewHeimdall(database database.Database) Heimdall {
	return &HeimdallImpl{
		database: database,
	}
}
