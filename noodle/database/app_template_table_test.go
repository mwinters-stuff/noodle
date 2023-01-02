package database_test

import (
	"net"
	"testing"

	"github.com/jackc/pgmock"
	"github.com/jackc/pgproto3/v2"
	dbf "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AppTemplateTableTestSuite struct {
	suite.Suite
	script   *pgmock.Script
	listener net.Listener
	connStr  string
}

func (suite *AppTemplateTableTestSuite) SetupSuite() {
}

func (suite *AppTemplateTableTestSuite) SetupTest() {
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.connStr = dbf.TestStepsRunner(suite.T(), suite.script)
}

func (suite *AppTemplateTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *AppTemplateTableTestSuite) TestCreateTable() {
	dbf.SetupConnectionSteps(suite.script)

	suite.script.Steps = append(suite.script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	suite.script.Steps = append(suite.script.Steps, dbf.ExpectMessageX(&pgproto3.Query{
		String: "CREATE TABLE IF NOT EXISTS application_template (\n  appid CHAR(40) PRIMARY KEY,\n  name VARCHAR(20) UNIQUE,\n  website VARCHAR(100) UNIQUE,\n  license VARCHAR(100),\n  description VARCHAR(1000),\n  enhanced BOOL,\n  tilebackground VARCHAR(256),\n  icon VARCHAR(256), \n  sha CHAR(40)\n)"}))
	suite.script.Steps = append(suite.script.Steps, pgmock.SendMessage(&pgproto3.CommandComplete{CommandTag: []byte("CREATE TABLE")}))

	suite.script.Steps = append(suite.script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))
	suite.script.Steps = append(suite.script.Steps, dbf.ExpectMessageX(&pgproto3.Query{String: "CREATE INDEX IF NOT EXISTS application_template_idx1 ON application_template(name)"}))
	suite.script.Steps = append(suite.script.Steps, pgmock.SendMessage(&pgproto3.CommandComplete{CommandTag: []byte("CREATE INDEX")}))

	suite.script.Steps = append(suite.script.Steps, pgmock.SendMessage(&pgproto3.ReadyForQuery{TxStatus: 'I'}))

	db := database.NewDatabase(suite.connStr)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestInsert() {
	dbf.SetupConnectionSteps(suite.script)

	dbf.InsertUpdateDeleteMock(suite.script, "INSERT INTO application_template (appid,name,website,license,description,enhanced,tilebackground,icon,sha) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042, 1043, 1043, 1043, 1043, 16, 1043, 1043, 1042}},
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0, 0, 0, 0, 0, 1, 0, 0, 0},
			Parameters: [][]byte{
				[]byte("140902edbcc424c09736af28ab2de604c3bde936"),
				[]byte("AdGuard Home"),
				[]byte("https://github.com/AdguardTeam/AdGuardHome"),
				[]byte("GNU General Public License v3.0 only"),
				[]byte("AdGuard Home is a network-wide software for blocking ads."),
				{1},
				[]byte("light"),
				[]byte("adguardhome.png"),
				[]byte("ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"),
			},
			ResultFormatCodes: []int16{},
		},
	)

	db := database.NewDatabase(suite.connStr)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := jsontypes.App{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	err = table.Insert(app)
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestUpdate() {
	dbf.SetupConnectionSteps(suite.script)

	dbf.InsertUpdateDeleteMock(suite.script, "UPDATE application_template SET name = $2,website = $3,license = $4,description = $5,enhanced = $6,tilebackground = $7,icon = $8,sha = $9 WHERE appid = $1",
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042, 1043, 1043, 1043, 1043, 16, 1043, 1043, 1042}},
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0, 0, 0, 0, 0, 1, 0, 0, 0},
			Parameters: [][]byte{
				[]byte("140902edbcc424c09736af28ab2de604c3bde936"),
				[]byte("AdGuard Home"),
				[]byte("https://github.com/AdguardTeam/AdGuardHome"),
				[]byte("GNU General Public License v3.0 only"),
				[]byte("AdGuard Home is a network-wide software for blocking ads."),
				{1},
				[]byte("light"),
				[]byte("adguardhome.png"),
				[]byte("ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"),
			},
			ResultFormatCodes: []int16{},
		},
	)

	db := database.NewDatabase(suite.connStr)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := jsontypes.App{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	err = table.Update(app)
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestDelete() {
	dbf.SetupConnectionSteps(suite.script)

	dbf.InsertUpdateDeleteMock(suite.script, "DELETE FROM application_template WHERE appid = $1",
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042}},
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("140902edbcc424c09736af28ab2de604c3bde936"),
			},
			ResultFormatCodes: []int16{},
		},
	)

	db := database.NewDatabase(suite.connStr)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := jsontypes.App{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	err = table.Delete(app)
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestSearch() {
	dbf.SetupConnectionSteps(suite.script)

	dbf.SelectMock(suite.script, `SELECT * FROM application_template WHERE name LIKE $1`,
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{25}},
		[]pgproto3.FieldDescription{
			{Name: []byte("appid"), TableOID: 16666, TableAttributeNumber: 1, DataTypeOID: 1042, DataTypeSize: -1, TypeModifier: 44, Format: 0},
			{Name: []byte("name"), TableOID: 16666, TableAttributeNumber: 2, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 24, Format: 0},
			{Name: []byte("website"), TableOID: 16666, TableAttributeNumber: 3, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 104, Format: 0},
			{Name: []byte("license"), TableOID: 16666, TableAttributeNumber: 4, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 104, Format: 0},
			{Name: []byte("description"), TableOID: 16666, TableAttributeNumber: 5, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 1004, Format: 0},
			{Name: []byte("enhanced"), TableOID: 16666, TableAttributeNumber: 6, DataTypeOID: 16, DataTypeSize: 1, TypeModifier: -1, Format: 0},
			{Name: []byte("tilebackground"), TableOID: 16666, TableAttributeNumber: 7, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 260, Format: 0},
			{Name: []byte("icon"), TableOID: 16666, TableAttributeNumber: 8, DataTypeOID: 1043, DataTypeSize: -1, TypeModifier: 260, Format: 0},
			{Name: []byte("sha"), TableOID: 16666, TableAttributeNumber: 9, DataTypeOID: 1042, DataTypeSize: -1, TypeModifier: 44, Format: 0},
		},

		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("%AdGuard%"),
			},
			ResultFormatCodes: []int16{0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		[][]byte{
			[]byte("140902edbcc424c09736af28ab2de604c3bde936"),
			[]byte("AdGuard Home"),
			[]byte("https://github.com/AdguardTeam/AdGuardHome"),
			[]byte("GNU General Public License v3.0 only"),
			[]byte("AdGuard Home is a network-wide software for blocking ads."),
			{1},
			[]byte("light"),
			[]byte("adguardhome.png"),
			[]byte("ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"),
		},
	)

	db := database.NewDatabase(suite.connStr)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := jsontypes.App{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	result, err := table.Search("AdGuard")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []jsontypes.App{app}, result)

}

func (suite *AppTemplateTableTestSuite) TestExists() {
	dbf.SetupConnectionSteps(suite.script)

	dbf.SelectMock(suite.script, `SELECT COUNT(*) FROM application_template WHERE appid = $1`,
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042}},
		[]pgproto3.FieldDescription{
			{Name: []byte("count"), TableOID: 0, TableAttributeNumber: 0, DataTypeOID: 20, DataTypeSize: 8, TypeModifier: -1, Format: 0},
		},

		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("140902edbcc424c09736af28ab2de604c3bde936"),
			},
			ResultFormatCodes: []int16{1},
		},
		[][]byte{
			[]byte("0000000000000001"),
		},
	)

	db := database.NewDatabase(suite.connStr)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	result, err := table.Exists("140902edbcc424c09736af28ab2de604c3bde936")
	require.NoError(suite.T(), err)
	require.True(suite.T(), result)

}

func (suite *AppTemplateTableTestSuite) TestNotExists() {
	dbf.SetupConnectionSteps(suite.script)

	dbf.SelectMock(suite.script, `SELECT COUNT(*) FROM application_template WHERE appid = $1`,
		pgproto3.ParameterDescription{ParameterOIDs: []uint32{1042}},
		[]pgproto3.FieldDescription{
			{Name: []byte("count"), TableOID: 0, TableAttributeNumber: 0, DataTypeOID: 20, DataTypeSize: 8, TypeModifier: -1, Format: 0},
		},

		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: []int16{0},
			Parameters: [][]byte{
				[]byte("140902edbcc424c09736af28ab2de604c3bde936"),
			},
			ResultFormatCodes: []int16{1},
		},
		[][]byte{
			[]byte("0000000000000000"),
		},
	)

	db := database.NewDatabase(suite.connStr)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	result, err := table.Exists("140902edbcc424c09736af28ab2de604c3bde936")
	require.NoError(suite.T(), err)
	require.False(suite.T(), result)

}

func TestAppTemplateTableSuite(t *testing.T) {
	suite.Run(t, new(AppTemplateTableTestSuite))
}
