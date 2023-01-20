package database_test

import (
	"net"
	"testing"

	"github.com/jackc/pgmock"
	"github.com/rs/zerolog/log"

	database_test "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ApplicationTabTableTestSuite struct {
	suite.Suite
	script        *pgmock.Script
	listener      net.Listener
	appConfig     yamltypes.AppConfig
	testFunctions database_test.TestFunctions
}

func (suite *ApplicationTabTableTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
}

func (suite *ApplicationTabTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *ApplicationTabTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *ApplicationTabTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.CreateApplicationTabsTableSteps(suite.T(), suite.script)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationTabTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *ApplicationTabTableTestSuite) TestUpgrade() {
	table := database.NewApplicationTabTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *ApplicationTabTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE application_tabs"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationTabTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *ApplicationTabTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_7","Query":"INSERT INTO application_tabs (tabid, applicationid, displayorder) VALUES ($1, $2, $3) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_7"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,23,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27029,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_7","ParameterFormatCodes":[1,1,1],"Parameters":[{"binary":"00000001"},{"binary":"00000001"},{"binary":"00000001"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27029,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	at1 := models.ApplicationTab{
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  1,
	}

	table := database.NewApplicationTabTable(db)

	err = table.Insert(&at1)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), at1.ID, int64(0))

}

func (suite *ApplicationTabTableTestSuite) TestUpdate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_8","Query":"UPDATE application_tabs SET displayorder = $2 WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_8"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_8","ParameterFormatCodes":[1,1],"Parameters":[{"binary":"00000003"},{"binary":"00000001"}],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"CommandComplete","CommandTag":"UPDATE 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)
	at3 := models.ApplicationTab{
		ID:            3,
		TabID:         2,
		ApplicationID: 2,
		DisplayOrder:  1,
	}
	table := database.NewApplicationTabTable(db)

	err = table.Update(at3)
	require.NoError(suite.T(), err)
}

func (suite *ApplicationTabTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"DELETE FROM application_tabs WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"CommandComplete","CommandTag":"DELETE 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationTabTable(db)

	err = table.Delete(int64(1))
	require.NoError(suite.T(), err)
}

func (suite *ApplicationTabTableTestSuite) TestGetTabApps() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_9","Query":"SELECT at.id, at.displayorder, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM application_tabs at, applications app WHERE at.tabid = $1 AND app.id = at.applicationid ORDER BY at.displayorder","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_9"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27029,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"displayorder","TableOID":27029,"TableAttributeNumber":4,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":27006,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":27006,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27006,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27006,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27006,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27006,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":27006,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27006,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_9","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[1,1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27029,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"displayorder","TableOID":27029,"TableAttributeNumber":4,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":27006,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":27006,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27006,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27006,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27006,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27006,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":27006,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27006,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"binary":"00000001"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"binary":"00000002"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":"Apache License 2.0"},{"text":"Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB."},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationTabTable(db)

	result, err := table.GetTabApps(1)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.ApplicationTab{
		{
			ID:            1,
			TabID:         1,
			ApplicationID: 1,
			DisplayOrder:  1,
			Application: &models.Application{
				ID:             1,
				TemplateAppid:  "",
				Name:           "AdGuard Home",
				Website:        "https://github.com/AdguardTeam/AdGuardHome",
				License:        "GNU General Public License v3.0 only",
				Description:    "AdGuard Home is a network-wide software for blocking ads.",
				Enhanced:       true,
				TileBackground: "light",
				Icon:           "adguardhome.png",
			},
		},
		{
			ID:            2,
			TabID:         1,
			ApplicationID: 2,
			DisplayOrder:  2,
			Application: &models.Application{
				ID:             2,
				Name:           "Adminer",
				TemplateAppid:  "",
				Website:        "https://www.adminer.org",
				License:        "Apache License 2.0",
				Description:    "Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB.",
				Enhanced:       false,
				TileBackground: "light",
				Icon:           "adminer.svg",
			},
		},
	}, result)

}

func (suite *ApplicationTabTableTestSuite) TestGetTabAppsError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_9","Query":"SELECT at.id, at.displayorder, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM application_tabs at, applications app WHERE at.tabid = $1 AND app.id = at.applicationid ORDER BY at.displayorder","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_9"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27029,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"displayorder","TableOID":27029,"TableAttributeNumber":4,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":27006,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":27006,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27006,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27006,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27006,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27006,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":27006,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27006,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_9","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[1,1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27029,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"displayorder","TableOID":27029,"TableAttributeNumber":4,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":27006,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":27006,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27006,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27006,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27006,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27006,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":27006,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27006,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"binary":"00000001"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"binary":"00000002"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":"Apache License 2.0"},{"text":"Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB."},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationTabTable(db)
	db.Close()
	_, err = table.GetTabApps(1)
	require.Error(suite.T(), err)

}

func TestApplicationTabTableSuite(t *testing.T) {
	suite.Run(t, new(ApplicationTabTableTestSuite))
}
