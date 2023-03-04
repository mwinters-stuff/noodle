package heimdall_test

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	"github.com/mwinters-stuff/noodle/noodle/options"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HeimdallAppsTestSuite struct {
	suite.Suite

	mockDatabase         *mocks.Database
	mockAppTemplateTable *mocks.AppTemplateTable

	testOptions options.NoodleOptions
}

func (suite *HeimdallAppsTestSuite) SetupSuite() {
	heimdall.Logger = log.Output(nil)
}

func (suite *HeimdallAppsTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockAppTemplateTable = mocks.NewAppTemplateTable(suite.T())

	database.NewAppTemplateTable = func(database database.Database) database.AppTemplateTable {
		return suite.mockAppTemplateTable
	}

	suite.testOptions.HeimdallIconBaseURL = "http://site/icons"
	suite.testOptions.HeimdallListJsonURL = "http://site/list.json"
	suite.testOptions.IconSavePath = "/tmp"

}

func (suite *HeimdallAppsTestSuite) TearDownTest() {
	os.Remove("/tmp/abc.def")
	database.NewAppTemplateTable = database.NewAppTemplateTableImpl
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

	app1 := models.ApplicationTemplate{
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

	app2 := models.ApplicationTemplate{
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

	gock.New("http://site").
		Get("/list.json").
		Reply(200).
		JSON(json)

	gock.New("http://site").
		Get("/icons/adguardhome.png").
		Reply(200).
		BodyString("FILECONTENTS")

	gock.New("http://site").
		Get("/icons/zulip.png").
		Reply(200).
		BodyString("FILECONTENTS")

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	suite.mockAppTemplateTable.EXPECT().Exists("140902edbcc424c09736af28ab2de604c3bde936").Return(false, nil)
	suite.mockAppTemplateTable.EXPECT().Exists("d17139efd0d8e0cba9bf8380c9818838911dfe0f").Return(true, nil)

	suite.mockAppTemplateTable.EXPECT().Insert(app1).Return(nil)
	suite.mockAppTemplateTable.EXPECT().Update(app2).Return(nil)

	err := h.UpdateFromServer()
	require.NoError(suite.T(), err)

}

func (suite *HeimdallAppsTestSuite) TestUpdateFromServerFailedGet() {

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").ReplyError(errors.New("failed"))

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	err := h.UpdateFromServer()
	require.EqualError(suite.T(), err, "Get \"http://site/list.json\": failed")

}

func (suite *HeimdallAppsTestSuite) TestUpdateFromServerFailed401() {

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").Response.Status(401).BodyString("not found")

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	err := h.UpdateFromServer()
	require.EqualError(suite.T(), err, "401 Unauthorized")

}

func (suite *HeimdallAppsTestSuite) TestUpdateFromServerFailedUnMarshal() {

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").Response.Status(200)

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	err := h.UpdateFromServer()
	require.EqualError(suite.T(), err, "unexpected end of JSON input")

}

func (suite *HeimdallAppsTestSuite) TestUpdateFromServerFailDatabaseExists() {
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
		
		}
	]
	}`

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").
		Reply(200).
		JSON(json)

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	suite.mockAppTemplateTable.EXPECT().Exists("140902edbcc424c09736af28ab2de604c3bde936").Return(false, errors.New("something went wrong"))

	err := h.UpdateFromServer()
	require.EqualError(suite.T(), err, "something went wrong")

}

func (suite *HeimdallAppsTestSuite) TestUpdateFromServerFailDatabaseUpdate() {
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
		
		}
	]
	}`

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").
		Reply(200).
		JSON(json)

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	app1 := models.ApplicationTemplate{
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

	suite.mockAppTemplateTable.EXPECT().Exists("140902edbcc424c09736af28ab2de604c3bde936").Return(false, nil)
	suite.mockAppTemplateTable.EXPECT().Insert(app1).Return(errors.New("something else went wrong"))

	err := h.UpdateFromServer()
	require.EqualError(suite.T(), err, "something else went wrong")
}

func (suite *HeimdallAppsTestSuite) TestDownloadIconError() {
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
			}
	]
	}`

	app1 := models.ApplicationTemplate{
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

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").
		Reply(200).
		JSON(json)

	gock.New("http://site").
		Get("/icons/adguardhome.png").
		Reply(401)

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	suite.mockAppTemplateTable.EXPECT().Exists("140902edbcc424c09736af28ab2de604c3bde936").Return(false, nil)

	suite.mockAppTemplateTable.EXPECT().Insert(app1).Return(nil)

	err := h.UpdateFromServer()
	require.ErrorContains(suite.T(), err, "download icon failed: 401 Unauthorized")

}

func (suite *HeimdallAppsTestSuite) TestDownloadIconHttpError() {
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
			}
	]
	}`

	app1 := models.ApplicationTemplate{
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

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").
		Reply(200).
		JSON(json)

	gock.New("http://site").
		Get("/icons/adguardhome.png").
		ReplyError(errors.New("failed"))

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	suite.mockAppTemplateTable.EXPECT().Exists("140902edbcc424c09736af28ab2de604c3bde936").Return(false, nil)

	suite.mockAppTemplateTable.EXPECT().Insert(app1).Return(nil)

	err := h.UpdateFromServer()
	require.ErrorContains(suite.T(), err, "failed")

}

func (suite *HeimdallAppsTestSuite) TestDownloadIconWriteFileError() {
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
			}
	]
	}`

	app1 := models.ApplicationTemplate{
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

	suite.testOptions.IconSavePath = "/somepath/that/doesnt/exist"

	defer gock.Off()

	gock.New("http://site").
		Get("/list.json").
		Reply(200).
		JSON(json)

	gock.New("http://site").
		Get("/icons/adguardhome.png").
		Reply(200).
		BodyString("FILECONTENTS")

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	suite.mockAppTemplateTable.EXPECT().Exists("140902edbcc424c09736af28ab2de604c3bde936").Return(false, nil)

	suite.mockAppTemplateTable.EXPECT().Insert(app1).Return(nil)

	err := h.UpdateFromServer()
	require.ErrorContains(suite.T(), err, "open /somepath/that/doesnt/exist/adguardhome.png: no such file or directory")

}

func (suite *HeimdallAppsTestSuite) TestListIcons() {
	suite.testOptions.IconSavePath = "/tmp"

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	files, err := h.ListIcons()

	require.NoError(suite.T(), err)
	require.NotEmpty(suite.T(), files)
}

func (suite *HeimdallAppsTestSuite) TestListIconsErr() {
	suite.testOptions.IconSavePath = "/tmptemp"

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	files, err := h.ListIcons()

	require.Error(suite.T(), err)
	require.Empty(suite.T(), files)
}

func (suite *HeimdallAppsTestSuite) TestUploadIcon() {
	suite.testOptions.IconSavePath = "/tmp"

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	reader := io.NopCloser(strings.NewReader("hello world"))

	err := h.UploadIcon("abc.def", reader)

	require.NoError(suite.T(), err)

	require.FileExists(suite.T(), "/tmp/abc.def")

	content, err := os.ReadFile("/tmp/abc.def")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "hello world", string(content))

}

func (suite *HeimdallAppsTestSuite) TestUploadIconError() {
	suite.testOptions.IconSavePath = "/tmptemp"

	h := heimdall.NewHeimdall(suite.mockDatabase, suite.testOptions)
	require.NotNil(suite.T(), h)

	reader := io.NopCloser(strings.NewReader("hello world"))

	err := h.UploadIcon("abc.def", reader)

	require.Error(suite.T(), err)

	require.NoFileExists(suite.T(), "/tmp/abc.def")

}

func TestHeimdallAppsSuite(t *testing.T) {
	suite.Run(t, new(HeimdallAppsTestSuite))
}
