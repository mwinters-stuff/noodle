package database_test

import (
	"net"
	"testing"

	"github.com/jackc/pgmock"

	database_test "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GroupApplicationsTableTestSuite struct {
	suite.Suite
	script        *pgmock.Script
	listener      net.Listener
	appConfig     yamltypes.AppConfig
	testFunctions database_test.TestFunctions
}

func (suite *GroupApplicationsTableTestSuite) SetupSuite() {
}

func (suite *GroupApplicationsTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *GroupApplicationsTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *GroupApplicationsTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.CreateGroupApplicationsTableSteps(suite.T(), suite.script)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewGroupApplicationsTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *GroupApplicationsTableTestSuite) TestUpgrade() {
	table := database.NewGroupApplicationsTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *GroupApplicationsTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE group_applications"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewGroupApplicationsTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *GroupApplicationsTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_12","Query":"INSERT INTO group_applications (groupid, applicationid) VALUES ($1, $2) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_12"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27128,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_12","ParameterFormatCodes":[1,1],"Parameters":[{"binary":"00000002"},{"binary":"00000001"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27128,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	ga1 := models.GroupApplications{
		GroupID:       2,
		ApplicationID: 1,
	}
	table := database.NewGroupApplicationsTable(db)

	err = table.Insert(&ga1)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), ga1.ID, int64(0))

}

func (suite *GroupApplicationsTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"DELETE FROM group_applications WHERE id = $1","ParameterOIDs":null}`,
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

	table := database.NewGroupApplicationsTable(db)

	err = table.Delete(1)
	require.NoError(suite.T(), err)
}

func (suite *GroupApplicationsTableTestSuite) TestGetGroupApps() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_13","Query":"SELECT ga.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM group_applications ga, applications app WHERE ga.groupid = $1 AND app.id = ga.applicationid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_13"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27128,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":27088,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":27088,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27088,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27088,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27088,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27088,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":27088,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27088,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_13","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27128,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":27088,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":27088,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27088,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27088,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27088,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27088,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":27088,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27088,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":"Apache License 2.0"},{"text":"Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB."},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewGroupApplicationsTable(db)

	result, err := table.GetGroupApps(2)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.GroupApplications{
		{
			ID:            1,
			GroupID:       2,
			ApplicationID: 1,
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
			GroupID:       2,
			ApplicationID: 2,
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

func (suite *GroupApplicationsTableTestSuite) TestGetGroupAppsError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_13","Query":"SELECT ga.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM group_applications ga, applications app WHERE ga.groupid = $1 AND app.id = ga.applicationid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_13"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27128,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":27088,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":27088,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27088,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27088,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27088,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27088,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":27088,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27088,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_13","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27128,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":27088,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":27088,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27088,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27088,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27088,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27088,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":27088,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27088,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":"Apache License 2.0"},{"text":"Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB."},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewGroupApplicationsTable(db)
	db.Close()
	_, err = table.GetGroupApps(2)
	require.Error(suite.T(), err)

}

func TestGroupApplicationsTableSuite(t *testing.T) {
	suite.Run(t, new(GroupApplicationsTableTestSuite))
}
