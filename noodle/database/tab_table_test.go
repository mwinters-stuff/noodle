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

type TabTableTestSuite struct {
	suite.Suite
	script        *pgmock.Script
	listener      net.Listener
	appConfig     yamltypes.AppConfig
	testFunctions database_test.TestFunctions
}

func (suite *TabTableTestSuite) SetupSuite() {
}

func (suite *TabTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *TabTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *TabTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.CreateTabsTableSteps(suite.T(), suite.script)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewTabTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *TabTableTestSuite) TestUpgrade() {
	table := database.NewTabTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *TabTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE tabs"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewTabTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *TabTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_4","Query":"INSERT INTO tabs (label,displayorder) VALUES ($1, $2) RETURNING id","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_4"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1043,23]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26461,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_4","ParameterFormatCodes":[0,1],"Parameters":[{"text":"Servers"},{"binary":"00000001"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26461,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	tab1 := models.Tab{
		Label:        "Servers",
		DisplayOrder: 1,
	}

	table := database.NewTabTable(db)

	err = table.Insert(&tab1)
	require.NoError(suite.T(), err)
	require.Greater(suite.T(), tab1.ID, int64(0))

}

func (suite *TabTableTestSuite) TestUpdate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_5","Query":"UPDATE tabs SET label = $2,displayorder = $3 WHERE id = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_5"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[23,1043,23]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_5","ParameterFormatCodes":[1,0,1],"Parameters":[{"binary":"00000001"},{"text":"Servers"},{"binary":"00000001"}],"ResultFormatCodes":[]}`,
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
	tab1 := models.Tab{
		ID:           1,
		Label:        "Servers",
		DisplayOrder: 1,
	}
	table := database.NewTabTable(db)

	err = table.Update(tab1)
	require.NoError(suite.T(), err)
}

func (suite *TabTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"DELETE FROM tabs WHERE id = $1","ParameterOIDs":null}`,
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

	table := database.NewTabTable(db)

	err = table.Delete(1)
	require.NoError(suite.T(), err)
}

func (suite *TabTableTestSuite) TestGetAll() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_6","Query":"SELECT * FROM tabs ORDER BY displayorder","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_6"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26592,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"label","TableOID":26592,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayorder","TableOID":26592,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_6","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26592,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"label","TableOID":26592,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayorder","TableOID":26592,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"Servers"},{"binary":"00000001"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"text":"Apps"},{"binary":"00000002"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewTabTable(db)

	result, err := table.GetAll()
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []*models.Tab{
		{
			ID:           1,
			Label:        "Servers",
			DisplayOrder: 1,
		},
		{
			ID:           2,
			Label:        "Apps",
			DisplayOrder: 2,
		},
	}, result)

}

func (suite *TabTableTestSuite) TestGetAllError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_6","Query":"SELECT * FROM tabs ORDER BY displayorder","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_6"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26592,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0},{"Name":"label","TableOID":26592,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayorder","TableOID":26592,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_6","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1,0,1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"id","TableOID":26592,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1},{"Name":"label","TableOID":26592,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":204,"Format":0},{"Name":"displayorder","TableOID":26592,"TableAttributeNumber":3,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"},{"text":"Servers"},{"binary":"00000001"}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000002"},{"text":"Apps"},{"binary":"00000002"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 2"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewTabTable(db)
	db.Close()
	_, err = table.GetAll()
	require.Error(suite.T(), err)

}

func TestTabTableSuite(t *testing.T) {
	suite.Run(t, new(TabTableTestSuite))
}
