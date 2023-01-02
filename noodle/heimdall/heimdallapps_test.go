package heimdall_test

import (
	"testing"

	"github.com/h2non/gock"
	mocks "github.com/mwinters-stuff/noodle/internal/mocks/app"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HeimdallAppsTestSuite struct {
	suite.Suite

	mockDatabase         *mocks.Database
	mockAppTemplateTable *mocks.AppTemplateTable
}

func (suite *HeimdallAppsTestSuite) SetupSuite() {
}

func (suite *HeimdallAppsTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockAppTemplateTable = mocks.NewAppTemplateTable(suite.T())

	database.NewAppTemplateTable = func(database database.Database) database.AppTemplateTable {
		return suite.mockAppTemplateTable
	}
}

func (suite *HeimdallAppsTestSuite) TearDownTest() {

}

func (suite *HeimdallAppsTestSuite) TestUpdateFromServer() {
	json := `
	{
		"appcount": 2,
		"apps": [
			{
				"appid": "140902edbcc424c09736af28ab2de604c3bde936",
				"name": "AdGuard Home",
				"website": "https://github.com/AdguardTeam/AdGuardHome",
				"license": "GNU General Public License v3.0 only",
				"description": "AdGuard Home is a network-wide software for blocking ads & tracking.",
				"enhanced": true,
				"tile_background": "light",
				"icon": "adguardhome.png",
				"sha": "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"
			},
			{
			"appid": "d17139efd0d8e0cba9bf8380c9818838911dfe0f",
			"name": "Zulip",
			"website": "https://zulipchat.com",
			"license": "Apache License 2.0",
			"description": "Powerful open source team chat.",
			"enhanced": false,
			"tile_background": "light",
			"icon": "zulip.png",
			"sha": "3a0df46433fcc2077745b553566c7064958c5092"
		}
	]
	}`

	app1 := jsontypes.App{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads & tracking.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	app2 := jsontypes.App{
		Appid:          "d17139efd0d8e0cba9bf8380c9818838911dfe0f",
		Name:           "Zulip",
		Website:        "https://zulipchat.com",
		License:        "Apache License 2.0",
		Description:    "Powerful open source team chat.",
		Enhanced:       false,
		TileBackground: "light",
		Icon:           "zulip.png",
		SHA:            "3a0df46433fcc2077745b553566c7064958c5092",
	}

	defer gock.Off()

	gock.New("https://appslist.heimdall.site").
		Get("/list.json").
		Reply(200).
		JSON(json)

	h := heimdall.NewHeimdall(suite.mockDatabase)
	require.NotNil(suite.T(), h)

	suite.mockAppTemplateTable.EXPECT().Exists("140902edbcc424c09736af28ab2de604c3bde936").Return(false, nil)
	suite.mockAppTemplateTable.EXPECT().Exists("d17139efd0d8e0cba9bf8380c9818838911dfe0f").Return(true, nil)

	suite.mockAppTemplateTable.EXPECT().Insert(app1).Return(nil)
	suite.mockAppTemplateTable.EXPECT().Update(app2).Return(nil)

	err := h.UpdateFromServer()
	require.NoError(suite.T(), err)

}

func (suite *HeimdallAppsTestSuite) TestFindApps() {

	result := []jsontypes.App{
		{
			Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
			Name:           "AdGuard Home",
			Website:        "https://github.com/AdguardTeam/AdGuardHome",
			License:        "GNU General Public License v3.0 only",
			Description:    "AdGuard Home is a network-wide software for blocking ads & tracking.",
			Enhanced:       true,
			TileBackground: "light",
			Icon:           "adguardhome.png",
			SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
		},

		{
			Appid:          "d17139efd0d8e0cba9bf8380c9818838911dfe0f",
			Name:           "Zulip Home",
			Website:        "https://zulipchat.com",
			License:        "Apache License 2.0",
			Description:    "Powerful open source team chat.",
			Enhanced:       false,
			TileBackground: "light",
			Icon:           "zulip.png",
			SHA:            "3a0df46433fcc2077745b553566c7064958c5092",
		},
	}

	h := heimdall.NewHeimdall(suite.mockDatabase)
	require.NotNil(suite.T(), h)

	suite.mockAppTemplateTable.EXPECT().Search("Home").Return(result, nil)

	apps, err := h.FindApps("Home")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), apps)
	require.ElementsMatch(suite.T(), apps, result)

}

func TestHeimdallAppsSuite(t *testing.T) {
	suite.Run(t, new(HeimdallAppsTestSuite))
}
