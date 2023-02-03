package database_test

import (
	"net"
	"testing"

	"github.com/jackc/pgmock"
	database_test "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/options"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserTableTestSuite struct {
	suite.Suite

	script        *pgmock.Script
	listener      net.Listener
	appConfig     options.AllNoodleOptions
	testFunctions database_test.TestFunctions
}

func (suite *UserTableTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
}

func (suite *UserTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}
	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *UserTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *UserTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type": "Query", "String": "CREATE TABLE IF NOT EXISTS users (\n  id SERIAL PRIMARY KEY,\n  username VARCHAR(50) UNIQUE,\n  dn VARCHAR(200) UNIQUE,\n  displayname VARCHAR(100),\n  givenname VARCHAR(100),\n  surname VARCHAR(100),\n  uidnumber INTEGER\n)"}`,
		`B {"Type": "CommandComplete", "CommandTag": "CREATE TABLE"}`,
		`B {"Type": "ReadyForQuery", "TxStatus": "I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *UserTableTestSuite) TestUpgrade() {
	table := database.NewUserTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *UserTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE users"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *UserTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_?","Query":"INSERT INTO users (username, dn, displayname, givenname, surname, uidnumber) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_?"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1043,1043,1043,1043,1043,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":24907,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_?","ParameterFormatCodes":[0,0,0,0,0,1],"Parameters":[{"text":"bobe"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"Bob Extample"},{"text":"Bob"},{"text":"Extample"},{"binary":"000003e9"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":24907,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	user := models.User{
		DN:        "CN=bob,DC=example,DC=nz",
		Username:  "bobe",
		Surname:   "Extample",
		GivenName: "Bob",
		UIDNumber: 1001,
	}

	table := database.NewUserTable(db)

	err = table.Insert(&user)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), user.ID, int64(0))

}

func (suite *UserTableTestSuite) TestUpdate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_?","Query":"UPDATE users SET username = $2,dn = $3,displayname = $4,givenname = $5,surname = $6,uidnumber = $7 WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_?"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,1043,1043,1043,1043,1043,23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_?","ParameterFormatCodes":[1,0,0,0,0,0,1],"Parameters":[{"binary":"00000001"},{"text":"bobe"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"Bob Extample"},{"text":"Bob"},{"text":"Extample"},{"binary":"000003e9"}],"ResultFormatCodes":[]}`,
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

	user := models.User{
		ID:        1,
		DN:        "CN=bob,DC=example,DC=nz",
		Username:  "bobe",
		Surname:   "Extample",
		GivenName: "Bob",
		UIDNumber: 1001,
	}

	table := database.NewUserTable(db)
	err = table.Update(user)
	require.NoError(suite.T(), err)
}

func (suite *UserTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_?","Query":"DELETE FROM users WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_?"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_?","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[]}`,
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

	user := models.User{
		ID:          1,
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}

	table := database.NewUserTable(db)

	err = table.Delete(user)
	require.NoError(suite.T(), err)
}

func (suite *UserTableTestSuite) TestGetAll() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_?","Query":"SELECT * FROM users ORDER BY username","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_?"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25060,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"username","TableOID":25060,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25060,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25060,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25060,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25060,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25060,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_?","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1,0,0,0,0,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25060,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"username","TableOID":25060,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25060,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25060,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25060,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25060,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25060,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"bobe"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobextample"},{"text":"Bob"},{"text":"Extample"},{"binary":"000003e9"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"text":"jack"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"Jack M"},{"text":"Jack"},{"text":"M"},{"binary":"000003ea"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.GetAll()
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.User{
		{
			ID:          1,
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
	}, result)

}

func (suite *UserTableTestSuite) TestGetAllError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_?","Query":"SELECT * FROM users ORDER BY username","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_?"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25060,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"username","TableOID":25060,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25060,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25060,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25060,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25060,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25060,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_?","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1,0,0,0,0,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25060,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"username","TableOID":25060,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25060,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25060,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25060,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25060,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25060,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"bobe"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobextample"},{"text":"Bob"},{"text":"Extample"},{"binary":"000003e9"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"text":"jack"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"Jack M"},{"text":"Jack"},{"text":"M"},{"binary":"000003ea"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)
	db.Close()

	_, err = table.GetAll()

	require.Error(suite.T(), err)

}

func (suite *UserTableTestSuite) TestGetDN() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT * FROM users WHERE dn = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25089,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"username","TableOID":25089,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25089,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25089,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25089,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25089,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25089,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"CN=jack,DC=example,DC=nz"}],"ResultFormatCodes":[1,0,0,0,0,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25089,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"username","TableOID":25089,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25089,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25089,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25089,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25089,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25089,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"text":"jack"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"Jack M"},{"text":"Jack"},{"text":"M"},{"binary":"000003ea"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.GetDN("CN=jack,DC=example,DC=nz")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Equal(suite.T(), models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}, result)

}

func (suite *UserTableTestSuite) TestGetDNError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT * FROM users WHERE dn = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25089,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"username","TableOID":25089,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25089,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25089,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25089,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25089,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25089,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"CN=jack,DC=example,DC=nz"}],"ResultFormatCodes":[1,0,0,0,0,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25089,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"username","TableOID":25089,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25089,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25089,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25089,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25089,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25089,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"text":"jack"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"Jack M"},{"text":"Jack"},{"text":"M"},{"binary":"000003ea"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)
	db.Close()

	_, err = table.GetDN("CN=jack,DC=example,DC=nz")
	require.Error(suite.T(), err)

}

func (suite *UserTableTestSuite) TestGetID() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT * FROM users WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25117,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"username","TableOID":25117,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25117,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25117,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25117,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25117,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25117,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[1,0,0,0,0,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25117,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"username","TableOID":25117,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":25117,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":25117,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":25117,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":25117,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":25117,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"text":"jack"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"Jack M"},{"text":"Jack"},{"text":"M"},{"binary":"000003ea"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.GetID(2)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Equal(suite.T(), models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}, result)
}

func (suite *UserTableTestSuite) TestGetIDNoResultError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT * FROM users WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27670,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"username","TableOID":27670,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":27670,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":27670,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":27670,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":27670,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":27670,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[1],"Parameters":[{"binary":"ffffffff"}],"ResultFormatCodes":[1,0,0,0,0,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":27670,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"username","TableOID":27670,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0},{"Name":"dn","TableOID":27670,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayname","TableOID":27670,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"givenname","TableOID":27670,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"surname","TableOID":27670,"TableAttributeNumber":6,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"uidnumber","TableOID":27670,"TableAttributeNumber":7,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 0"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	_, err = table.GetID(-1)
	require.EqualError(suite.T(), err, "no user with id -1")

}

func (suite *UserTableTestSuite) TestGetIDError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT * FROM users WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ErrorResponse","Severity":"ERROR","SeverityUnlocalized":"ERROR","Code":"42P01","Message":"relation \"users\" does not exist","Detail":"","Hint":"","Position":21,"InternalPosition":0,"InternalQuery":"","Where":"","SchemaName":"","TableName":"","ColumnName":"","DataTypeName":"","ConstraintName":"","File":"parse_relation.c","Line":1392,"Routine":"parserOpenTable","UnknownFields":null}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	_, err = table.GetID(12)
	require.EqualError(suite.T(), err, "ERROR: relation \"users\" does not exist (SQLSTATE 42P01)")

}

func (suite *UserTableTestSuite) TestExistsDN() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT COUNT(*) FROM users WHERE dn = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"CN=jack,DC=example,DC=nz"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.ExistsDN("CN=jack,DC=example,DC=nz")
	require.NoError(suite.T(), err)
	require.True(suite.T(), result)

}

func (suite *UserTableTestSuite) TestExistsUsername() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT COUNT(*) FROM users WHERE username = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"bobe"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.ExistsUsername("bobe")
	require.NoError(suite.T(), err)
	require.True(suite.T(), result)

}

func (suite *UserTableTestSuite) TestExistsNotDN() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT COUNT(*) FROM users WHERE dn = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"CN=bob,DC=example,DC=nz"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000000"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.ExistsDN("CN=bob,DC=example,DC=nz")
	require.NoError(suite.T(), err)
	require.False(suite.T(), result)
}

func (suite *UserTableTestSuite) TestNotExistsUsername() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT COUNT(*) FROM users WHERE username = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"bob"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000000"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.ExistsUsername("bob")
	require.NoError(suite.T(), err)
	require.False(suite.T(), result)
}

func TestUserTableSuite(t *testing.T) {
	suite.Run(t, new(UserTableTestSuite))
}
