package database_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type databaseLogHook struct {
	LastEvent *zerolog.Event
	LastLevel zerolog.Level
	LastMsg   string
}

func (h *databaseLogHook) Run(e *zerolog.Event, l zerolog.Level, m string) {
	h.LastEvent = e
	h.LastLevel = l
	h.LastMsg = m
}

type DatabaseTestInitialSuite struct {
	suite.Suite
	loghook databaseLogHook

	connectionString      string
	createConnectionSring string
}

func (suite *DatabaseTestInitialSuite) SetupSuite() {
	suite.connectionString = fmt.Sprintf("postgres://testu:testu@%s/testd?sslmode=disable", os.Getenv("POSTGRES_HOSTNAME"))

	pguser := os.Getenv("POSTGRES_USER")
	pgpassword := os.Getenv("POSTGRES_PASSWORD")
	pgdb := os.Getenv("POSTGRES_DB")
	pghostname := os.Getenv("POSTGRES_HOSTNAME")
	suite.createConnectionSring = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pguser, pgpassword, pghostname, pgdb)
}

func (suite *DatabaseTestInitialSuite) SetupTest() {
	suite.loghook = databaseLogHook{}
	database.Logger = log.Hook(&suite.loghook)

	conn, err := pgx.Connect(context.Background(), suite.createConnectionSring)
	if err != nil {
		assert.Error(suite.T(), err)
	}
	defer conn.Close(context.Background())
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `create user testu with password 'testu'`).Close()
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `create database testd with owner testu`).Close()
	assert.Nil(suite.T(), err)
}

func (suite *DatabaseTestInitialSuite) TearDownTest() {
	conn, err := pgx.Connect(context.Background(), suite.createConnectionSring)
	if err != nil {
		assert.Error(suite.T(), err)
	}
	defer conn.Close(context.Background())
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `drop database testd`).Close()
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `drop user testu`).Close()
	assert.Nil(suite.T(), err)

}

func (suite *DatabaseTestInitialSuite) TearDownSuite() {

}

func (suite *DatabaseTestInitialSuite) TestBadConnect() {
	db := database.NewDatabase("connectionstring")
	assert.NotNil(suite.T(), db)

	err := db.Connect()
	assert.Error(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "database connection failed cannot parse `connectionstring`: failed to parse as DSN (invalid dsn)"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestConnect() {

	db := database.NewDatabase(suite.connectionString)
	err := db.Connect()
	assert.Nil(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "database connected"
	}, time.Second, time.Millisecond*100)

	db.Close()
}

func (suite *DatabaseTestInitialSuite) TestCreate() {

	db := database.NewDatabase(suite.connectionString)
	err := db.Connect()
	assert.Nil(suite.T(), err)

	err = db.Create()
	assert.Nil(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "creating database"
	}, time.Second, time.Millisecond*100)

	db.Close()
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestInitialSuite))
}

type DatabaseTestCreatedSuite struct {
	suite.Suite
	loghook databaseLogHook
	db      database.Database

	connectionString      string
	createConnectionSring string
}

func (suite *DatabaseTestCreatedSuite) SetupSuite() {
	suite.connectionString = fmt.Sprintf("postgres://testu:testu@%s/testd?sslmode=disable", os.Getenv("POSTGRES_HOSTNAME"))

	pguser := os.Getenv("POSTGRES_USER")
	pgpassword := os.Getenv("POSTGRES_PASSWORD")
	pgdb := os.Getenv("POSTGRES_DB")
	pghostname := os.Getenv("POSTGRES_HOSTNAME")
	suite.createConnectionSring = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pguser, pgpassword, pghostname, pgdb)
}

func (suite *DatabaseTestCreatedSuite) SetupTest() {
	suite.loghook = databaseLogHook{}
	database.Logger = log.Hook(&suite.loghook)

	conn, err := pgx.Connect(context.Background(), suite.createConnectionSring)
	if err != nil {
		assert.Error(suite.T(), err)
	}
	defer conn.Close(context.Background())
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `create user testu with password 'testu'`).Close()
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `create database testd with owner testu`).Close()
	assert.Nil(suite.T(), err)

	suite.db = database.NewDatabase(suite.connectionString)
	err = suite.db.Connect()
	assert.Nil(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "database connected"
	}, time.Second, time.Millisecond*100)

	err = suite.db.Create()
	assert.Nil(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "creating database"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestCreatedSuite) TearDownTest() {
	suite.db.Close()

	conn, err := pgx.Connect(context.Background(), suite.createConnectionSring)
	if err != nil {
		assert.Error(suite.T(), err)
	}
	defer conn.Close(context.Background())
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `drop database testd`).Close()
	assert.Nil(suite.T(), err)
	err = conn.PgConn().Exec(context.Background(), `drop user testu`).Close()
	assert.Nil(suite.T(), err)

}

func (suite *DatabaseTestCreatedSuite) TearDownSuite() {
}

func (suite *DatabaseTestCreatedSuite) TestGetVersion() {
	version, err := suite.db.GetVersion()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, version)
	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "current database version 1"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestCreatedSuite) TestUpgradeSameVersion() {
	err := suite.db.CheckUpgrade()
	assert.Nil(suite.T(), err)
	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "no database upgrade required"
	}, time.Second, time.Millisecond*100)
}

func (suite *DatabaseTestCreatedSuite) TestUpgradeDifferentVersion() {
	suite.db.Pool().Exec(context.Background(), "UPDATE version SET version = 0")

	err := suite.db.CheckUpgrade()
	assert.Nil(suite.T(), err)
	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "upgrade database from 0 to 1"
	}, time.Second, time.Millisecond*100)
}

func (suite *DatabaseTestCreatedSuite) TestDownUpgrade() {
	suite.db.Pool().Exec(context.Background(), "UPDATE version SET version = 100")

	err := suite.db.CheckUpgrade()
	assert.Error(suite.T(), err)
	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "cannot downgrade database"
	}, time.Second, time.Millisecond*100)
}

// func (suite *DatabaseTestCreatedSuite) TestCreate() {

// }

func TestDatabaseCreatedSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestCreatedSuite))
}
