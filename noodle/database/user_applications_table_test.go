package database_test

import (
	"net"
	"testing"

	"github.com/jackc/pgmock"
	"github.com/rs/zerolog/log"

	database_test "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/options"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserApplicationsTableTestSuite struct {
	suite.Suite

	script        *pgmock.Script
	listener      net.Listener
	appConfig     options.AllNoodleOptions
	testFunctions database_test.TestFunctions
}

func (suite *UserApplicationsTableTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
}

func (suite *UserApplicationsTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *UserApplicationsTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *UserApplicationsTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"CREATE TABLE IF NOT EXISTS user_applications (\n  id SERIAL PRIMARY KEY,\n  userid int REFERENCES users(id) ON DELETE CASCADE,\n  applicationid int REFERENCES applications(id) ON DELETE CASCADE\n)"}`,
		`B {"Type": "CommandComplete", "CommandTag": "CREATE TABLE"}`,
		`B {"Type": "ReadyForQuery", "TxStatus": "I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *UserApplicationsTableTestSuite) TestUpgrade() {
	table := database.NewUserApplicationsTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *UserApplicationsTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE user_applications"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *UserApplicationsTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_13","Query":"INSERT INTO user_applications (userid, applicationid) VALUES ($1, $2) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_13"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27369,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_13","ParameterFormatCodes":[1,1],"Parameters":[{"binary":"00000001"},{"binary":"00000001"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27369,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	ua1 := models.UserApplications{
		UserID:        1,
		ApplicationID: 1,
	}
	table := database.NewUserApplicationsTable(db)

	err = table.Insert(&ua1)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), ua1.ID, int64(0))

}

func (suite *UserApplicationsTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_15","Query":"DELETE FROM user_applications WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_15"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_15","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"CommandComplete","CommandTag":"DELETE 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)

	err = table.Delete(2)
	require.NoError(suite.T(), err)
}

func (suite *UserApplicationsTableTestSuite) TestGetUserApps() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_14","Query":"SELECT ua.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM user_applications ua, applications app WHERE ua.userid = $1 AND app.id = ua.applicationid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_14"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27369,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":27312,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":27312,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27312,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27312,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27312,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27312,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":27312,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27312,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_14","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27369,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":27312,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":27312,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27312,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27312,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27312,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27312,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":27312,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27312,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":"Apache License 2.0"},{"text":"Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB."},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)

	result, err := table.GetUserApps(1)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.UserApplications{
		{
			ID:            1,
			UserID:        1,
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
			UserID:        1,
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

func (suite *UserApplicationsTableTestSuite) TestGetUserAppsError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_14","Query":"SELECT ua.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon FROM user_applications ua, applications app WHERE ua.userid = $1 AND app.id = ua.applicationid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_14"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27369,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":27312,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":27312,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27312,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27312,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27312,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27312,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":27312,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27312,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_14","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27369,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":27312,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":27312,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":27312,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":27312,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":27312,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":27312,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":27312,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":27312,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":"Apache License 2.0"},{"text":"Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB."},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)
	db.Close()
	_, err = table.GetUserApps(1)
	require.Error(suite.T(), err)

}

func (suite *UserApplicationsTableTestSuite) TestGetUserAllowedApps() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT d.tabid, d.displayorder, a.id as application_id,a.name,a.website,a.license,a.description,a.enhanced,a.tilebackground,a.icon\nFROM applications a ,\n(\n  SELECT ua.applicationid, at.tabid, at.displayorder FROM user_applications ua, application_tabs at WHERE at.applicationid = ua.applicationid AND userid = $1\n  UNION\n  SELECT ga.applicationid, at.tabid, at.displayorder FROM user_groups ug, group_applications ga, application_tabs at WHERE at.applicationid = ga.applicationid AND ga.groupid = ug.groupid AND userid = $1\n  UNION\n  SELECT applicationid, tabid, displayorder FROM application_tabs at WHERE at.applicationid NOT IN (SELECT applicationid  FROM user_applications UNION select applicationid FROM group_applications)\n) as d\nWHERE a.id = d.applicationid\nORDER BY d.tabid, d.displayorder;\n","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000004"}],"ResultFormatCodes":[1,1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000000"},{"binary":"00000002"},{"text":"applicationtab1"},{"text":"string"},{"text":"string"},{"text":"application_tab_1"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000006"},{"binary":"00000001"},{"text":"usercustomapp"},{"text":"string"},{"text":"string"},{"text":"user custom app"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000000"},{"binary":"00000003"},{"text":"applicationtab2"},{"text":"string"},{"text":"string"},{"text":"application_tab_2"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"binary":"00000006"},{"text":"application2"},{"text":"string"},{"text":"string"},{"text":"application_2"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"00000004"},{"text":"applicationtab3"},{"text":"string"},{"text":"string"},{"text":"application_tab_3"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000001"},{"binary":"00000005"},{"text":"application1"},{"text":"string"},{"text":"string"},{"text":"application_1"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000002"},{"binary":"00000007"},{"text":"application3"},{"text":"string"},{"text":"string"},{"text":"application_3"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 7"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)

	result, err := table.GetUserAllowdApplications(4)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Len(suite.T(), result, 7)

	require.ElementsMatch(suite.T(), []*models.UsersApplicationItem{
		{
			Application: &models.Application{
				Description:    "application_tab_1",
				Enhanced:       false,
				Icon:           "string",
				ID:             2,
				License:        "string",
				Name:           "applicationtab1",
				TemplateAppid:  "",
				TileBackground: "string",
				Website:        "string",
			},
			DisplayOrder: 0,
			TabID:        1,
		},
		{
			Application: &models.Application{
				Description:    "user custom app",
				Enhanced:       false,
				Icon:           "string",
				ID:             1,
				License:        "string",
				Name:           "usercustomapp",
				TemplateAppid:  "",
				TileBackground: "string",
				Website:        "string",
			},
			DisplayOrder: 6,
			TabID:        1,
		},
		{
			Application: &models.Application{
				Description:    "application_tab_2",
				Enhanced:       false,
				Icon:           "string",
				ID:             3,
				License:        "string",
				Name:           "applicationtab2",
				TemplateAppid:  "",
				TileBackground: "string",
				Website:        "string",
			},
			DisplayOrder: 0,
			TabID:        2,
		},
		{
			Application: &models.Application{
				Description:    "application_2",
				Enhanced:       false,
				Icon:           "string",
				ID:             6,
				License:        "string",
				Name:           "application2",
				TemplateAppid:  "",
				TileBackground: "string",
				Website:        "string",
			},
			DisplayOrder: 2,
			TabID:        2,
		},
		{
			Application: &models.Application{
				Description:    "application_tab_3",
				Enhanced:       false,
				Icon:           "string",
				ID:             4,
				License:        "string",
				Name:           "applicationtab3",
				TemplateAppid:  "",
				TileBackground: "string",
				Website:        "string",
			},
			DisplayOrder: 0,
			TabID:        3,
		},
		{
			Application: &models.Application{
				Description:    "application_1",
				Enhanced:       false,
				Icon:           "string",
				ID:             5,
				License:        "string",
				Name:           "application1",
				TemplateAppid:  "",
				TileBackground: "string",
				Website:        "string",
			},
			DisplayOrder: 1,
			TabID:        3,
		},
		{
			Application: &models.Application{
				Description:    "application_3",
				Enhanced:       false,
				Icon:           "string",
				ID:             7,
				License:        "string",
				Name:           "application3",
				TemplateAppid:  "",
				TileBackground: "string",
				Website:        "string",
			},
			DisplayOrder: 2,
			TabID:        3,
		},
	}, result)

}

func (suite *UserApplicationsTableTestSuite) TestGetUserAllowedAppsError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT d.tabid, d.displayorder, a.id as application_id,a.name,a.website,a.license,a.description,a.enhanced,a.tilebackground,a.icon\nFROM applications a ,\n(\n  SELECT ua.applicationid, at.tabid, at.displayorder FROM user_applications ua, application_tabs at WHERE at.applicationid = ua.applicationid AND userid = $1\n  UNION\n  SELECT ga.applicationid, at.tabid, at.displayorder FROM user_groups ug, group_applications ga, application_tabs at WHERE at.applicationid = ga.applicationid AND ga.groupid = ug.groupid AND userid = $1\n  UNION\n  SELECT applicationid, tabid, displayorder FROM application_tabs at WHERE at.applicationid NOT IN (SELECT applicationid  FROM user_applications UNION select applicationid FROM group_applications)\n) as d\nWHERE a.id = d.applicationid\nORDER BY d.tabid, d.displayorder;\n","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000004"}],"ResultFormatCodes":[1,1,1,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000000"},{"binary":"00000002"},{"text":"applicationtab1"},{"text":"string"},{"text":"string"},{"text":"application_tab_1"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000006"},{"binary":"00000001"},{"text":"usercustomapp"},{"text":"string"},{"text":"string"},{"text":"user custom app"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000000"},{"binary":"00000003"},{"text":"applicationtab2"},{"text":"string"},{"text":"string"},{"text":"application_tab_2"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000002"},{"binary":"00000006"},{"text":"application2"},{"text":"string"},{"text":"string"},{"text":"application_2"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"00000004"},{"text":"applicationtab3"},{"text":"string"},{"text":"string"},{"text":"application_tab_3"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000001"},{"binary":"00000005"},{"text":"application1"},{"text":"string"},{"text":"string"},{"text":"application_1"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000002"},{"binary":"00000007"},{"text":"application3"},{"text":"string"},{"text":"string"},{"text":"application_3"},{"binary":"00"},{"text":"string"},{"text":"string"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 7"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)
	db.Close()
	_, err = table.GetUserAllowdApplications(1)
	require.Error(suite.T(), err)

}

func TestUserApplicationsTableSuite(t *testing.T) {
	suite.Run(t, new(UserApplicationsTableTestSuite))
}
