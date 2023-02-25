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
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT ua.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon,app.template_appid FROM user_applications ua, applications app WHERE ua.userid = $1 AND app.id = ua.applicationid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":36710,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000004"}],"ResultFormatCodes":[1,1,0,0,0,0,1,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":36710,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000005"},{"binary":"0000000a"},{"text":"Cloudflare"},{"text":"https://dash.cloudflare.com/"},{"text":""},{"text":"Cloudflare"},{"binary":"00"},{"text":"dark"},{"text":"cloudflare.svg"},{"text":"f036f579066ad71bd653f5a6418dbede5b500370"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000006"},{"binary":"0000000b"},{"text":"Bazarr"},{"text":"https://github.com/morpheus65535/bazarr"},{"text":""},{"text":"Bazarr"},{"binary":"00"},{"text":"dark"},{"text":"bazarr.png"},{"text":"085f0b437f9bf9c98bb68b745c8dcf323a7e0499"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserApplicationsTable(db)

	result, err := table.GetUserApps(4)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.UserApplications{
		{
			Application: &models.Application{
				Description:    "Cloudflare",
				Icon:           "cloudflare.svg",
				ID:             10,
				Name:           "Cloudflare",
				TemplateAppid:  "f036f579066ad71bd653f5a6418dbede5b500370",
				TileBackground: "dark",
				Website:        "https://dash.cloudflare.com/",
			},
			ApplicationID: 10,
			ID:            5,
			UserID:        4,
		},
		{
			Application: &models.Application{
				Description:    "Bazarr",
				Icon:           "bazarr.png",
				ID:             11,
				Name:           "Bazarr",
				TemplateAppid:  "085f0b437f9bf9c98bb68b745c8dcf323a7e0499",
				TileBackground: "dark",
				Website:        "https://github.com/morpheus65535/bazarr",
			},
			ApplicationID: 11,
			ID:            6,
			UserID:        4,
		},
	}, result)

}

func (suite *UserApplicationsTableTestSuite) TestGetUserAppsError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT ua.id, app.id, app.name,app.website,app.license,app.description,app.enhanced,app.tilebackground,app.icon,app.template_appid FROM user_applications ua, applications app WHERE ua.userid = $1 AND app.id = ua.applicationid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":36710,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000004"}],"ResultFormatCodes":[1,1,0,0,0,0,1,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":36710,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000005"},{"binary":"0000000a"},{"text":"Cloudflare"},{"text":"https://dash.cloudflare.com/"},{"text":""},{"text":"Cloudflare"},{"binary":"00"},{"text":"dark"},{"text":"cloudflare.svg"},{"text":"f036f579066ad71bd653f5a6418dbede5b500370"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000006"},{"binary":"0000000b"},{"text":"Bazarr"},{"text":"https://github.com/morpheus65535/bazarr"},{"text":""},{"text":"Bazarr"},{"binary":"00"},{"text":"dark"},{"text":"bazarr.png"},{"text":"085f0b437f9bf9c98bb68b745c8dcf323a7e0499"}]}`,
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
	_, err = table.GetUserApps(4)
	require.Error(suite.T(), err)

}

func (suite *UserApplicationsTableTestSuite) TestGetUserAllowedApps() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT d.tabid, d.displayorder, a.id as application_id,a.name,a.website,a.license,a.description,a.enhanced,a.tilebackground,a.icon,a.template_appid\nFROM applications a ,\n(\n  SELECT ua.applicationid, at.tabid, at.displayorder FROM user_applications ua, application_tabs at WHERE at.applicationid = ua.applicationid AND userid = $1\n  UNION\n  SELECT ga.applicationid, at.tabid, at.displayorder FROM user_groups ug, group_applications ga, application_tabs at WHERE at.applicationid = ga.applicationid AND ga.groupid = ug.groupid AND userid = $1\n  UNION\n  SELECT applicationid, tabid, displayorder FROM application_tabs at WHERE at.applicationid NOT IN (SELECT applicationid  FROM user_applications UNION select applicationid FROM group_applications)\n) as d\nWHERE a.id = d.applicationid\nORDER BY d.tabid, d.displayorder;\n","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000004"}],"ResultFormatCodes":[1,1,1,0,0,0,0,1,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000000"},{"binary":"00000010"},{"text":"ArchiveTeam Warrior"},{"text":"https://www.archiveteam.org/index.php?title=ArchiveTeam_Warrior"},{"text":""},{"text":"ArchiveTeam Warrior"},{"binary":"00"},{"text":"light"},{"text":"archiveteamwarrior.png"},{"text":"5eef559f19eadb9593bafbd3ca6155dc6721a0d7"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000000"},{"binary":"00000013"},{"text":"Argo CD"},{"text":"https://argoproj.github.io/cd/"},{"text":""},{"text":"Argo CD"},{"binary":"00"},{"text":"dark"},{"text":"argocd.svg"},{"text":"88dc19bddba6e23ec39f978777b5adc5784ca27a"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000000"},{"binary":"0000000f"},{"text":"AMP"},{"text":"https://cubecoders.com/AMP"},{"text":""},{"text":"AMP"},{"binary":"00"},{"text":"light"},{"text":"amp.png"},{"text":"65f59ec6b1ecd6170d5044474043cca9560a8071"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"0000000e"},{"text":"CloudflareX"},{"text":"https://dash.cloudflare.com/"},{"text":""},{"text":"CloudflareX"},{"binary":"00"},{"text":"dark"},{"text":"cloudflare.svg"},{"text":"f036f579066ad71bd653f5a6418dbede5b500370"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"00000011"},{"text":"openHAB"},{"text":"https://openhab.org"},{"text":""},{"text":"openHAB"},{"binary":"00"},{"text":"light"},{"text":"openhab.png"},{"text":"c6a4fe0b25a74497e966f279f5186c99e5ce30e3"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"0000000d"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":""},{"text":"Adminer"},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"},{"text":"653caf8bdf55d6a99d77ceacd79f622353cd821a"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 6"}`,
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

	require.Len(suite.T(), result, 6)

	require.ElementsMatch(suite.T(), []*models.UsersApplicationItem{
		{
			Application: &models.Application{
				Description:    "ArchiveTeam Warrior",
				Icon:           "archiveteamwarrior.png",
				ID:             16,
				Name:           "ArchiveTeam Warrior",
				TemplateAppid:  "5eef559f19eadb9593bafbd3ca6155dc6721a0d7",
				TileBackground: "light",
				Website:        "https://www.archiveteam.org/index.php?title=ArchiveTeam_Warrior",
			},
			TabID: 1,
		},
		{
			Application: &models.Application{
				Description:    "Argo CD",
				Icon:           "argocd.svg",
				ID:             19,
				Name:           "Argo CD",
				TemplateAppid:  "88dc19bddba6e23ec39f978777b5adc5784ca27a",
				TileBackground: "dark",
				Website:        "https://argoproj.github.io/cd/",
			},
			TabID: 2,
		},
		{
			Application: &models.Application{
				Description:    "AMP",
				Icon:           "amp.png",
				ID:             15,
				Name:           "AMP",
				TemplateAppid:  "65f59ec6b1ecd6170d5044474043cca9560a8071",
				TileBackground: "light",
				Website:        "https://cubecoders.com/AMP",
			},
			TabID: 2,
		},
		{
			Application: &models.Application{
				Description:    "CloudflareX",
				Icon:           "cloudflare.svg",
				ID:             14,
				Name:           "CloudflareX",
				TemplateAppid:  "f036f579066ad71bd653f5a6418dbede5b500370",
				TileBackground: "dark",
				Website:        "https://dash.cloudflare.com/",
			},
			TabID: 3,
		},
		{
			Application: &models.Application{
				Description:    "openHAB",
				Icon:           "openhab.png",
				ID:             17,
				Name:           "openHAB",
				TemplateAppid:  "c6a4fe0b25a74497e966f279f5186c99e5ce30e3",
				TileBackground: "light",
				Website:        "https://openhab.org",
			},
			TabID: 3,
		},
		{
			Application: &models.Application{
				Description:    "Adminer",
				Icon:           "adminer.svg",
				ID:             13,
				Name:           "Adminer",
				TemplateAppid:  "653caf8bdf55d6a99d77ceacd79f622353cd821a",
				TileBackground: "light",
				Website:        "https://www.adminer.org",
			},
			TabID: 3,
		},
	}, result)

}

func (suite *UserApplicationsTableTestSuite) TestGetUserAllowedAppsError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT d.tabid, d.displayorder, a.id as application_id,a.name,a.website,a.license,a.description,a.enhanced,a.tilebackground,a.icon,a.template_appid\nFROM applications a ,\n(\n  SELECT ua.applicationid, at.tabid, at.displayorder FROM user_applications ua, application_tabs at WHERE at.applicationid = ua.applicationid AND userid = $1\n  UNION\n  SELECT ga.applicationid, at.tabid, at.displayorder FROM user_groups ug, group_applications ga, application_tabs at WHERE at.applicationid = ga.applicationid AND ga.groupid = ug.groupid AND userid = $1\n  UNION\n  SELECT applicationid, tabid, displayorder FROM application_tabs at WHERE at.applicationid NOT IN (SELECT applicationid  FROM user_applications UNION select applicationid FROM group_applications)\n) as d\nWHERE a.id = d.applicationid\nORDER BY d.tabid, d.displayorder;\n","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000004"}],"ResultFormatCodes":[1,1,1,0,0,0,0,1,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"tabid","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"displayorder","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"application_id","TableOID":36653,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"name","TableOID":36653,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"website","TableOID":36653,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"license","TableOID":36653,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":36653,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":36653,"TableAttributeNumber":7,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":36653,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":36653,"TableAttributeNumber":9,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"template_appid","TableOID":36653,"TableAttributeNumber":2,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000000"},{"binary":"00000010"},{"text":"ArchiveTeam Warrior"},{"text":"https://www.archiveteam.org/index.php?title=ArchiveTeam_Warrior"},{"text":""},{"text":"ArchiveTeam Warrior"},{"binary":"00"},{"text":"light"},{"text":"archiveteamwarrior.png"},{"text":"5eef559f19eadb9593bafbd3ca6155dc6721a0d7"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000000"},{"binary":"00000013"},{"text":"Argo CD"},{"text":"https://argoproj.github.io/cd/"},{"text":""},{"text":"Argo CD"},{"binary":"00"},{"text":"dark"},{"text":"argocd.svg"},{"text":"88dc19bddba6e23ec39f978777b5adc5784ca27a"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000000"},{"binary":"0000000f"},{"text":"AMP"},{"text":"https://cubecoders.com/AMP"},{"text":""},{"text":"AMP"},{"binary":"00"},{"text":"light"},{"text":"amp.png"},{"text":"65f59ec6b1ecd6170d5044474043cca9560a8071"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"0000000e"},{"text":"CloudflareX"},{"text":"https://dash.cloudflare.com/"},{"text":""},{"text":"CloudflareX"},{"binary":"00"},{"text":"dark"},{"text":"cloudflare.svg"},{"text":"f036f579066ad71bd653f5a6418dbede5b500370"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"00000011"},{"text":"openHAB"},{"text":"https://openhab.org"},{"text":""},{"text":"openHAB"},{"binary":"00"},{"text":"light"},{"text":"openhab.png"},{"text":"c6a4fe0b25a74497e966f279f5186c99e5ce30e3"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000000"},{"binary":"0000000d"},{"text":"Adminer"},{"text":"https://www.adminer.org"},{"text":""},{"text":"Adminer"},{"binary":"00"},{"text":"light"},{"text":"adminer.svg"},{"text":"653caf8bdf55d6a99d77ceacd79f622353cd821a"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 6"}`,
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
