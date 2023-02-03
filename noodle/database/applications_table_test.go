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

type ApplicationsTableTestSuite struct {
	suite.Suite
	script        *pgmock.Script
	listener      net.Listener
	appConfig     options.AllNoodleOptions
	testFunctions database_test.TestFunctions
}

func (suite *ApplicationsTableTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
}

func (suite *ApplicationsTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *ApplicationsTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *ApplicationsTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"CREATE TABLE IF NOT EXISTS applications (\n  id SERIAL PRIMARY KEY,\n  template_appid CHAR(40) REFERENCES application_template(appid) ON DELETE SET NULL,\n  name VARCHAR(50),\n  website VARCHAR(256),\n  license VARCHAR(100),\n  description VARCHAR(1000),\n  enhanced BOOL,\n  tilebackground VARCHAR(256),\n  icon VARCHAR(256)\n)"}`,
		`B {"Type": "CommandComplete", "CommandTag": "CREATE TABLE"}`,
		`B {"Type": "ReadyForQuery", "TxStatus": "I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationsTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *ApplicationsTableTestSuite) TestUpgrade() {
	table := database.NewApplicationsTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *ApplicationsTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE applications"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationsTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *ApplicationsTableTestSuite) TestInsertWithTemplate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"INSERT INTO applications (template_appid,name,website,license,description,enhanced,tilebackground,icon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042,1043,1043,1043,1043,16,1043,1043]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26221,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0,0,0,0,0,1,0,0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26221,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)
	application := models.Application{
		TemplateAppid:  "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}

	table := database.NewApplicationsTable(db)

	err = table.Insert(&application)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), application.ID, int64(0))

}

func (suite *ApplicationsTableTestSuite) TestInsertWithOutTemplate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"INSERT INTO applications (name,website,license,description,enhanced,tilebackground,icon) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1043,1043,1043,1043,16,1043,1043]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27006,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[0,0,0,0,1,0,0],"Parameters":[{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":"Apache License 2.0"},{"text":"Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB."},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27006,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)
	application := models.Application{
		Name:           "Adminer",
		Website:        "https://www.adminer.org",
		License:        "Apache License 2.0",
		Description:    "Adminer (formerly phpMinAdmin) is a full-featured database management tool written in PHP. Conversely to phpMyAdmin, it consists of a single file ready to deploy to the target server. Adminer is available for MySQL, MariaDB, PostgreSQL, SQLite, MS SQL, Oracle, Firebird, SimpleDB, Elasticsearch and MongoDB.",
		Enhanced:       false,
		TileBackground: "light",
		Icon:           "adminer.svg",
	}

	table := database.NewApplicationsTable(db)

	err = table.Insert(&application)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), application.ID, int64(0))

}

func (suite *ApplicationsTableTestSuite) TestUpdate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"UPDATE applications SET template_appid = $2, name = $3, website = $4, license = $5, description = $6, enhanced = $7,tilebackground = $8,icon = $9 WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,1042,1043,1043,1043,1043,16,1043,1043]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[1,0,0,0,0,0,1,0,0],"Parameters":[{"binary":"00000001"},{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"CommandComplete","CommandTag":"UPDATE 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)
	application := models.Application{
		ID:             1,
		TemplateAppid:  "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}

	table := database.NewApplicationsTable(db)

	err = table.Update(application)
	require.NoError(suite.T(), err)
}

func (suite *ApplicationsTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"DELETE FROM applications WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[]}`,
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

	table := database.NewApplicationsTable(db)

	err = table.Delete(int64(1))
	require.NoError(suite.T(), err)
}

func (suite *ApplicationsTableTestSuite) TestGetID() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT * FROM applications WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26334,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"template_appid","TableOID":26334,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26334,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26334,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26334,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26334,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26334,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":26334,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26334,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[1,0,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26334,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"template_appid","TableOID":26334,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26334,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26334,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26334,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26334,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26334,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":26334,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26334,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationsTable(db)

	result, err := table.GetID(1)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Equal(suite.T(), models.Application{
		ID:             1,
		TemplateAppid:  "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}, result)

}

func (suite *ApplicationsTableTestSuite) TestGetIDError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT * FROM applications WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26334,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"template_appid","TableOID":26334,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26334,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26334,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26334,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26334,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26334,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":26334,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26334,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[1,0,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26334,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"template_appid","TableOID":26334,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26334,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26334,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26334,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26334,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26334,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":26334,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26334,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationsTable(db)
	db.Close()
	_, err = table.GetID(1)
	require.Error(suite.T(), err)

}

func (suite *ApplicationsTableTestSuite) TestGetAppTemplateID() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT * FROM applications WHERE template_appid = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26390,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"template_appid","TableOID":26390,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26390,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26390,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26390,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26390,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26390,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":26390,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26390,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"}],"ResultFormatCodes":[1,0,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26390,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"template_appid","TableOID":26390,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26390,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26390,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26390,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26390,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26390,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":26390,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26390,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationsTable(db)

	result, err := table.GetTemplateID("140902edbcc424c09736af28ab2de604c3bde936")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.Application{
		{
			ID:             1,
			TemplateAppid:  "140902edbcc424c09736af28ab2de604c3bde936",
			Name:           "AdGuard Home",
			Website:        "https://github.com/AdguardTeam/AdGuardHome",
			License:        "GNU General Public License v3.0 only",
			Description:    "AdGuard Home is a network-wide software for blocking ads.",
			Enhanced:       true,
			TileBackground: "light",
			Icon:           "adguardhome.png",
		},
	}, result)

}

func (suite *ApplicationsTableTestSuite) TestGetAppTemplateIDError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT * FROM applications WHERE template_appid = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26390,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"template_appid","TableOID":26390,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26390,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26390,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26390,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26390,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26390,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":26390,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26390,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"}],"ResultFormatCodes":[1,0,0,0,0,0,1,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26390,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"template_appid","TableOID":26390,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":26390,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":26390,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":26390,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":26390,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":26390,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":26390,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":26390,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewApplicationsTable(db)
	db.Close()
	_, err = table.GetTemplateID("140902edbcc424c09736af28ab2de604c3bde936")
	require.Error(suite.T(), err)

}

func TestApplicationsTableSuite(t *testing.T) {
	suite.Run(t, new(ApplicationsTableTestSuite))
}
