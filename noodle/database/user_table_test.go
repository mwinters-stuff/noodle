package database_test

import (
	"net"
	"testing"

	"github.com/jackc/pgmock"
	"github.com/jackc/pgproto3/v2"
	dbf "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserTableTestSuite struct {
	suite.Suite
	script    *pgmock.Script
	listener  net.Listener
	appConfig yamltypes.AppConfig
}

func (suite *UserTableTestSuite) SetupSuite() {
}

func (suite *UserTableTestSuite) SetupTest() {
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = dbf.TestStepsRunner(suite.T(), suite.script)
}

func (suite *UserTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *UserTableTestSuite) TestCreateTable() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.CreateUserTableSteps(suite.T(), suite.script)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *UserTableTestSuite) TestInsert() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_?","Query":"INSERT INTO users (username, dn, displayname, givenname, surname, uidnumber) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_?"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1043,1043,1043,1043,1043,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":24907,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_?","ParameterFormatCodes":[0,0,0,0,0,1],"Parameters":[{"text":"bobe"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobextample"},{"text":"Bob"},{"text":"Extample"},{"binary":"000003e9"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":24907,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	user := database.User{
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UidNumber:   1001,
	}

	table := database.NewUserTable(db)

	err = table.Insert(&user)
	require.NoError(suite.T(), err)

}

func (suite *UserTableTestSuite) TestUpdate() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_?","Query":"UPDATE users SET username = $2,dn = $3,displayname = $4,givenname = $5,surname = $6,uidnumber = $7 WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_?"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,1043,1043,1043,1043,1043,23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_?","ParameterFormatCodes":[1,0,0,0,0,0,1],"Parameters":[{"binary":"00000001"},{"text":"bobe"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobextample"},{"text":"Bob"},{"text":"Extample"},{"binary":"000003e9"}],"ResultFormatCodes":[]}`,
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

	user := database.User{
		Id:          1,
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UidNumber:   1001,
	}

	table := database.NewUserTable(db)

	err = table.Update(user)
	require.NoError(suite.T(), err)
}

func (suite *UserTableTestSuite) TestDelete() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
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

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	user := database.User{
		Id:          1,
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UidNumber:   1001,
	}

	table := database.NewUserTable(db)

	err = table.Delete(user)
	require.NoError(suite.T(), err)
}

func (suite *UserTableTestSuite) TestGetAll() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)
	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
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
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.GetAll()
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []database.User{
		{
			Id:          1,
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UidNumber:   1001,
		},
		{
			Id:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UidNumber:   1002,
		},
	}, result)

}

func (suite *UserTableTestSuite) TestGetAllError() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)
	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
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
	db := database.NewDatabase(suite.appConfig)
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
	dbf.SetupConnectionSteps(suite.T(), suite.script)
	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
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
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.GetDN("CN=jack,DC=example,DC=nz")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Equal(suite.T(), database.User{
		Id:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UidNumber:   1002,
	}, result)

}

func (suite *UserTableTestSuite) TestGetDNError() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)
	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
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
	db := database.NewDatabase(suite.appConfig)
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
	dbf.SetupConnectionSteps(suite.T(), suite.script)
	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
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
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.GetID(2)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Equal(suite.T(), database.User{
		Id:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UidNumber:   1002,
	}, result)

}

func (suite *UserTableTestSuite) TestGetIDError() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)
	dbf.LoadDatabaseSteps(suite.T(), suite.script, []string{
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
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	db.Close()

	_, err = table.GetID(-1)
	require.Error(suite.T(), err)

}

// func (suite *UserTableTestSuite) TestSearchQueryFails() {
// 	dbf.SetupConnectionSteps(suite.T(), suite.script)

// 	dbf.SelectMock(suite.script, `SELECT * FROM application_template WHERE name LIKE $1`,
// 		pgproto3.ParameterDescription{ParameterOIDs: []uint32{25}},
// 		[]pgproto3.FieldDescription{
// 			{Name: []byte("appid"), TableOID: 16666, TableAttributeNumber: 1, DataTypeOID: 1042, DataTypeSize: -1, TypeModifier: 44, Format: 0},
// 			{Name: []byte("name"), TableOID: 16666, TableAttributeNumber: 2, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 24, Format: 0},
// 			{Name: []byte("website"), TableOID: 16666, TableAttributeNumber: 3, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 104, Format: 0},
// 			{Name: []byte("license"), TableOID: 16666, TableAttributeNumber: 4, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 104, Format: 0},
// 			{Name: []byte("description"), TableOID: 16666, TableAttributeNumber: 5, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 1004, Format: 0},
// 			{Name: []byte("enhanced"), TableOID: 16666, TableAttributeNumber: 6, DataTypeOID: 16, DataTypeSize: 1, TypeModifier: -1, Format: 0},
// 			{Name: []byte("tilebackground"), TableOID: 16666, TableAttributeNumber: 7, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 260, Format: 0},
// 			{Name: []byte("icon"), TableOID: 16666, TableAttributeNumber: 8, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 260, Format: 0},
// 			{Name: []byte("sha"), TableOID: 16666, TableAttributeNumber: 9, DataTypeOID: 1042, DataTypeSize: -1, TypeModifier: 44, Format: 0},
// 		},

// 		pgproto3.Bind{
// 			DestinationPortal:    "",
// 			PreparedStatement:    "stmtcache_?",
// 			ParameterFormatCodes: []int16{0},
// 			Parameters: [][]byte{
// 				[]byte("%AdGuard%"),
// 			},
// 			ResultFormatCodes: []int16{0, 0, 0, 0, 0, 1, 0, 0, 0},
// 		},
// 		[][]byte{
// 			[]byte("140902edbcc424c09736af28ab2de604c3bde936"),
// 			[]byte("AdGuard Home"),
// 			[]byte("https://github.com/AdguardTeam/AdGuardHome"),
// 			[]byte("GNU General Public License v3.0 only"),
// 			[]byte("AdGuard Home is a network-wide software for blocking ads."),
// 			{1},
// 			[]byte("light"),
// 			[]byte("adguardhome.png"),
// 			[]byte("ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"),
// 		},
// 	)

// 	db := database.NewDatabase(suite.appConfig)
// 	assert.NotNil(suite.T(), db)

// 	err := db.Connect()
// 	require.NoError(suite.T(), err)

// 	table := app_template_table.NewAppTemplateTable(db)

// 	db.Close()

// 	result, err := table.Search(`A`)
// 	require.Error(suite.T(), err)
// 	require.Nil(suite.T(), result)

// }

func (suite *UserTableTestSuite) TestExistsDN() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.SelectMock(suite.script, `SELECT COUNT(*) FROM users WHERE dn = $1`,
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042}},
		[]pgproto3.FieldDescription{
			{Name: []byte("count"), TableOID: 0, TableAttributeNumber: 0, DataTypeOID: 20, DataTypeSize: 8, TypeModifier: -1, Format: 0},
		},

		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("CN=bob,DC=example,DC=nz"),
			},
			ResultFormatCodes: []int16{1},
		},
		[][]byte{
			[]byte("0000000000000001"),
		},
	)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.ExistsDN("CN=bob,DC=example,DC=nz")
	require.NoError(suite.T(), err)
	require.True(suite.T(), result)
}

func (suite *UserTableTestSuite) TestExistsUsername() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.SelectMock(suite.script, `SELECT COUNT(*) FROM users WHERE username = $1`,
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042}},
		[]pgproto3.FieldDescription{
			{Name: []byte("count"), TableOID: 0, TableAttributeNumber: 0, DataTypeOID: 20, DataTypeSize: 8, TypeModifier: -1, Format: 0},
		},

		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("bob"),
			},
			ResultFormatCodes: []int16{1},
		},
		[][]byte{
			[]byte("0000000000000001"),
		},
	)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserTable(db)

	result, err := table.ExistsUsername("bob")
	require.NoError(suite.T(), err)
	require.True(suite.T(), result)
}

func (suite *UserTableTestSuite) TestExistsNotDN() {
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.SelectMock(suite.script, `SELECT COUNT(*) FROM users WHERE dn = $1`,
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042}},
		[]pgproto3.FieldDescription{
			{Name: []byte("count"), TableOID: 0, TableAttributeNumber: 0, DataTypeOID: 20, DataTypeSize: 8, TypeModifier: -1, Format: 0},
		},

		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("CN=bob,DC=example,DC=nz"),
			},
			ResultFormatCodes: []int16{1},
		},
		[][]byte{
			[]byte("0000000000000000"),
		},
	)

	db := database.NewDatabase(suite.appConfig)
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
	dbf.SetupConnectionSteps(suite.T(), suite.script)

	dbf.SelectMock(suite.script, `SELECT COUNT(*) FROM users WHERE username = $1`,
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042}},
		[]pgproto3.FieldDescription{
			{Name: []byte("count"), TableOID: 0, TableAttributeNumber: 0, DataTypeOID: 20, DataTypeSize: 8, TypeModifier: -1, Format: 0},
		},

		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("bob"),
			},
			ResultFormatCodes: []int16{1},
		},
		[][]byte{
			[]byte("0000000000000000"),
		},
	)

	db := database.NewDatabase(suite.appConfig)
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
