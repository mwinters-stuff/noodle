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

type UserSessionTableTestSuite struct {
	suite.Suite

	script        *pgmock.Script
	listener      net.Listener
	appConfig     options.AllNoodleOptions
	testFunctions database_test.TestFunctions
}

func (suite *UserSessionTableTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
}

func (suite *UserSessionTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}
	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *UserSessionTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *UserSessionTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.CreateUserSessionTableSteps(suite.T(), suite.script)

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	require.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserSessionTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *UserSessionTableTestSuite) TestUpgrade() {
	table := database.NewUserSessionTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *UserSessionTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE user_sessions"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	require.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserSessionTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *UserSessionTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"INSERT INTO user_sessions (user_id, token) VALUES ($1, $2) RETURNING id, issued, expires","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,1043]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35853,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"issued","TableOID":35853,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0},{"Name":"expires","TableOID":35853,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[1,0],"Parameters":[{"binary":"00000006"},{"text":"tokentokentoken"}],"ResultFormatCodes":[1,1,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35853,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"issued","TableOID":35853,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1},{"Name":"expires","TableOID":35853,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00029634b8e39d28"},{"binary":"000298a455f83d28"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	userSession := models.UserSession{
		UserID: 6,
		Token:  "tokentokentoken",
	}

	table := database.NewUserSessionTable(db)

	err = table.Insert(&userSession)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), userSession.ID, int64(0))
	require.Equal(suite.T(), "2023-01-27T02:52:17.811Z", userSession.Issued.String())
	require.Equal(suite.T(), "2023-02-27T02:52:17.811Z", userSession.Expires.String())

}

func (suite *UserSessionTableTestSuite) TestDeleteID() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_6","Query":"DELETE FROM user_sessions WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_6"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_6","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000001"}],"ResultFormatCodes":[]}`,
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

	table := database.NewUserSessionTable(db)

	err = table.Delete(1)
	require.NoError(suite.T(), err)
}

func (suite *UserSessionTableTestSuite) TestDeleteExpired() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DELETE FROM user_sessions WHERE expires \u003c NOW()"}`,
		`B {"Type":"CommandComplete","CommandTag":"DELETE 0"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserSessionTable(db)

	err = table.DeleteExpired()
	require.NoError(suite.T(), err)
}

func (suite *UserSessionTableTestSuite) TestGetUser() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"SELECT * FROM user_sessions WHERE user_id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35867,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"user_id","TableOID":35867,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"token","TableOID":35867,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"issued","TableOID":35867,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0},{"Name":"expires","TableOID":35867,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1],"Parameters":[{"binary":"00000006"}],"ResultFormatCodes":[1,1,0,1,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35867,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"user_id","TableOID":35867,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"token","TableOID":35867,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"issued","TableOID":35867,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1},{"Name":"expires","TableOID":35867,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000006"},{"text":"tokentokentoken"},{"binary":"0002963514cebd24"},{"binary":"000298a4b1e35d24"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserSessionTable(db)

	result, err := table.GetUser(6)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Equal(suite.T(), int64(1), result[0].ID)
	require.Equal(suite.T(), int64(6), result[0].UserID)
	require.Equal(suite.T(), "tokentokentoken", result[0].Token)

	require.Equal(suite.T(), "2023-01-27T03:17:59.947Z", result[0].Issued.String())
	require.Equal(suite.T(), "2023-02-27T03:17:59.947Z", result[0].Expires.String())
}

func (suite *UserSessionTableTestSuite) TestGetToken() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT * FROM user_sessions WHERE token = $1 AND expires > NOW()","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35867,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"user_id","TableOID":35867,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"token","TableOID":35867,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"issued","TableOID":35867,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0},{"Name":"expires","TableOID":35867,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[0],"Parameters":[{"text":"tokentokentoken"}],"ResultFormatCodes":[1,1,0,1,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35867,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"user_id","TableOID":35867,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"token","TableOID":35867,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"issued","TableOID":35867,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1},{"Name":"expires","TableOID":35867,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000006"},{"text":"tokentokentoken"},{"binary":"0002963514cebd24"},{"binary":"000298a4b1e35d24"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewUserSessionTable(db)

	result, err := table.GetToken("tokentokentoken")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.Equal(suite.T(), int64(1), result.ID)
	require.Equal(suite.T(), int64(6), result.UserID)
	require.Equal(suite.T(), "tokentokentoken", result.Token)

	require.Equal(suite.T(), "2023-01-27T03:17:59.947Z", result.Issued.String())
	require.Equal(suite.T(), "2023-02-27T03:17:59.947Z", result.Expires.String())
}

func (suite *UserSessionTableTestSuite) TestGetTokenError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"SELECT * FROM user_sessions WHERE token = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35867,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"user_id","TableOID":35867,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"token","TableOID":35867,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"issued","TableOID":35867,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0},{"Name":"expires","TableOID":35867,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[0],"Parameters":[{"text":"tokentokentoken"}],"ResultFormatCodes":[1,1,0,1,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":35867,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"user_id","TableOID":35867,"TableAttributeNumber":2,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"token","TableOID":35867,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"issued","TableOID":35867,"TableAttributeNumber":4,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1},{"Name":"expires","TableOID":35867,"TableAttributeNumber":5,"DataTypeOID":1114,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"binary":"00000006"},{"text":"tokentokentoken"},{"binary":"0002963514cebd24"},{"binary":"000298a4b1e35d24"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)

	err := db.Connect()
	require.NoError(suite.T(), err)
	db.Close()

	table := database.NewUserSessionTable(db)

	_, err = table.GetToken("tokentokentoken")
	require.Error(suite.T(), err)

}

func TestUserSessionTableSuite(t *testing.T) {
	suite.Run(t, new(UserSessionTableTestSuite))
}
