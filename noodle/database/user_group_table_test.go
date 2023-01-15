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

type UserGroupTableTestSuite struct {
	suite.Suite
	script        *pgmock.Script
	listener      net.Listener
	appConfig     yamltypes.AppConfig
	testFunctions database_test.TestFunctions
}

func (suite *UserGroupTableTestSuite) SetupSuite() {
}

func (suite *UserGroupTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *UserGroupTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *UserGroupTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.CreateUserGroupsTableSteps(suite.T(), suite.script)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *UserGroupTableTestSuite) TestUpgrade() {
	table := database.NewUserGroupsTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *UserGroupTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE user_groups"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *UserGroupTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"INSERT INTO user_groups (groupid, userid) VALUES ($1, $2) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25487,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[1,1],"Parameters":[{"binary":"00000001"},{"binary":"00000001"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25487,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	usergroup := models.UserGroup{
		GroupID: 1,
		UserID:  1,
	}

	table := database.NewUserGroupsTable(db)

	err = table.Insert(&usergroup)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), usergroup.ID, int64(0))

}

func (suite *UserGroupTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"DELETE FROM user_groups WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[]}`,
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

	usergroup2 := models.UserGroup{
		ID:      2,
		GroupID: 1,
		UserID:  1,
	}

	table := database.NewUserGroupsTable(db)

	err = table.Delete(usergroup2)
	require.NoError(suite.T(), err)
}

func (suite *UserGroupTableTestSuite) TestExists() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT COUNT(*) FROM user_groups WHERE groupid = $1 AND userid = $2","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1,1],"Parameters":[{"binary":"00000001"},{"binary":"00000001"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)

	result, err := table.Exists(1, 1)
	require.NoError(suite.T(), err)
	require.True(suite.T(), result)
}

func (suite *UserGroupTableTestSuite) TestNotExists() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT COUNT(*) FROM user_groups WHERE groupid = $1 AND userid = $2","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1,1],"Parameters":[{"binary":"00000001"},{"binary":"00000001"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000000"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)

	result, err := table.Exists(1, 1)
	require.NoError(suite.T(), err)
	require.False(suite.T(), result)
}

func (suite *UserGroupTableTestSuite) TestGetAll() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE g.id = ug.groupid AND u.id = ug.userid ORDER BY g.name","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25651,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"groupid","TableOID":25651,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"userid","TableOID":25651,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"group_dn","TableOID":25642,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25642,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25629,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25629,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1,1,1,0,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25651,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"groupid","TableOID":25651,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"userid","TableOID":25651,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"group_dn","TableOID":25642,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25642,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25629,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25629,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"binary":"00000001"},{"text":"cn=admins,ou=groups,dc=example,dc=nz"},{"text":"Admins"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobe"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000001"},{"binary":"00000002"},{"text":"cn=admins,ou=groups,dc=example,dc=nz"},{"text":"Admins"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"jack"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000002"},{"binary":"00000001"},{"text":"cn=users,ou=groups,dc=example,dc=nz"},{"text":"Users"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobe"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 3"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)

	result, err := table.GetAll()
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.UserGroup{
		{
			ID:        1,
			GroupID:   1,
			UserID:    1,
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=bob,DC=example,DC=nz",
			UserName:  "bobe",
		},
		{
			ID:        2,
			GroupID:   1,
			UserID:    2,
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
		{
			ID:        3,
			GroupID:   2,
			UserID:    1,
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=bob,DC=example,DC=nz",
			UserName:  "bobe",
		},
	}, result)

}

func (suite *UserGroupTableTestSuite) TestGetAllError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE g.id = ug.groupid AND u.id = ug.userid ORDER BY g.name","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25651,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"groupid","TableOID":25651,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"userid","TableOID":25651,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"group_dn","TableOID":25642,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25642,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25629,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25629,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1,1,1,0,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25651,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"groupid","TableOID":25651,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"userid","TableOID":25651,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"group_dn","TableOID":25642,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25642,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25629,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25629,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000001"},{"binary":"00000001"},{"text":"cn=admins,ou=groups,dc=example,dc=nz"},{"text":"Admins"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobe"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000001"},{"binary":"00000002"},{"text":"cn=admins,ou=groups,dc=example,dc=nz"},{"text":"Admins"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"jack"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000002"},{"binary":"00000001"},{"text":"cn=users,ou=groups,dc=example,dc=nz"},{"text":"Users"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobe"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 3"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)
	db.Close()
	_, err = table.GetAll()
	require.Error(suite.T(), err)

}

func (suite *UserGroupTableTestSuite) TestGetGroup() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE ug.groupid = $1 AND g.id = ug.groupid AND u.id = ug.userid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25705,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"groupid","TableOID":25705,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"userid","TableOID":25705,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"group_dn","TableOID":25696,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25696,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25683,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25683,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[1,1,1,0,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25705,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"groupid","TableOID":25705,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"userid","TableOID":25705,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"group_dn","TableOID":25696,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25696,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25683,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25683,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000002"},{"binary":"00000001"},{"text":"cn=users,ou=groups,dc=example,dc=nz"},{"text":"Users"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobe"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)

	result, err := table.GetGroup(2)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.UserGroup{
		{
			ID:        3,
			GroupID:   2,
			UserID:    1,
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=bob,DC=example,DC=nz",
			UserName:  "bobe",
		},
	}, result)

}

func (suite *UserGroupTableTestSuite) TestGetGroupError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE ug.groupid = $1 AND g.id = ug.groupid AND u.id = ug.userid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25705,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"groupid","TableOID":25705,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"userid","TableOID":25705,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"group_dn","TableOID":25696,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25696,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25683,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25683,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[1,1,1,0,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25705,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"groupid","TableOID":25705,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"userid","TableOID":25705,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"group_dn","TableOID":25696,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25696,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25683,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25683,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000003"},{"binary":"00000002"},{"binary":"00000001"},{"text":"cn=users,ou=groups,dc=example,dc=nz"},{"text":"Users"},{"text":"CN=bob,DC=example,DC=nz"},{"text":"bobe"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)
	db.Close()
	_, err = table.GetGroup(2)
	require.Error(suite.T(), err)

}

func (suite *UserGroupTableTestSuite) TestGetUser() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE ug.userid = $1 AND g.id = ug.groupid AND u.id = ug.userid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25760,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"groupid","TableOID":25760,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"userid","TableOID":25760,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"group_dn","TableOID":25751,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25751,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25738,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25738,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[1,1,1,0,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25760,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"groupid","TableOID":25760,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"userid","TableOID":25760,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"group_dn","TableOID":25751,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25751,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25738,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25738,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000001"},{"binary":"00000002"},{"text":"cn=admins,ou=groups,dc=example,dc=nz"},{"text":"Admins"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"jack"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)

	result, err := table.GetUser(2)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.UserGroup{
		{
			ID:        2,
			GroupID:   1,
			UserID:    2,
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	}, result)

}

func (suite *UserGroupTableTestSuite) TestGetUserError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT ug.id, ug.groupid, ug.userid, g.dn group_dn, g.name group_name, u.dn user_dn, u.username user_username FROM user_groups ug, groups g, users u WHERE ug.userid = $1 AND g.id = ug.groupid AND u.id = ug.userid","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25760,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"groupid","TableOID":25760,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"userid","TableOID":25760,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"group_dn","TableOID":25751,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25751,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25738,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25738,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000002"}],"ResultFormatCodes":[1,1,1,0,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":25760,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"groupid","TableOID":25760,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"userid","TableOID":25760,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"group_dn","TableOID":25751,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"group_name","TableOID":25751,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"user_dn","TableOID":25738,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"user_username","TableOID":25738,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":54,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"binary":"00000001"},{"binary":"00000002"},{"text":"cn=admins,ou=groups,dc=example,dc=nz"},{"text":"Admins"},{"text":"CN=jack,DC=example,DC=nz"},{"text":"jack"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserGroupsTable(db)
	db.Close()
	_, err = table.GetUser(2)
	require.Error(suite.T(), err)

}

func TestUserGroupTableSuite(t *testing.T) {
	suite.Run(t, new(UserGroupTableTestSuite))
}
